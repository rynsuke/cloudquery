package client

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/schema"
)

const maxFileSize = 1024 * 1024 * 20

func (c *Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvReverseTransformer.ReverseTransformValues(table, values)
	case FormatTypeJSON:
		return c.jsonReverseTransformer.ReverseTransformValues(table, values)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := strings.ReplaceAll(c.pluginSpec.Path, PathVarTable, table.Name)
	writerAtBuffer := manager.NewWriteAtBuffer(make([]byte, 0, maxFileSize))
	_, err := c.downloader.Download(ctx,
		writerAtBuffer,
		&s3.GetObjectInput{
			Bucket: aws.String(c.pluginSpec.Bucket),
			Key:    aws.String(name),
		})
	if err != nil {
		return err
	}
	r := bytes.NewReader(writerAtBuffer.Bytes())

	switch c.pluginSpec.FileSpec.Format {
	case FormatTypeCSV:
		opts := []csv.Options{
			csv.WithDelimiter(c.pluginSpec.FileSpec.Delimiter),
		}
		if c.pluginSpec.FileSpec.IncludeHeaders {
			opts = append(opts, csv.WithHeader())
		}

		client, err := csv.NewClient(opts...)
		if err != nil {
			return err
		}
		if err = client.Read(r, table, sourceName, res); err != nil {
			return err
		}
	case FormatTypeJSON:
		client, err := json.NewClient()
		if err != nil {
			return err
		}
		if err := client.Read(r, table, sourceName, res); err != nil {
			return err
		}
	default:
		panic("unknown format " + c.pluginSpec.FileSpec.Format)
	}
	return nil
}

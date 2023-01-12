package client

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	conn clickhouse.Conn

	logger zerolog.Logger

	spec specs.Destination

	destination.UnimplementedUnmanagedWriter
	*destination.DefaultReverseTransformer
}

func (c *Client) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	return &Client{}, nil
}

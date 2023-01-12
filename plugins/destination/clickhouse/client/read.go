package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	//TODO implement me
	panic("implement me")
}

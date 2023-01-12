package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	//TODO implement me
	panic("implement me")
}

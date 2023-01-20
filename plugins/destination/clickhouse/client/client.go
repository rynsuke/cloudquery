package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	db *sql.DB

	logger zerolog.Logger

	spec specs.Destination

	destination.UnimplementedUnmanagedWriter
	*destination.DefaultReverseTransformer
}

var _ destination.Client = (*Client)(nil)

func (c *Client) Close(context.Context) error {
	return c.db.Close()
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var pluginSpec Spec
	if err := spec.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	connector, err := pluginSpec.Connector()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connection %w", err)
	}

	return &Client{
		db:     sql.OpenDB(connector),
		logger: logger.With().Str("module", "dest-clickhouse").Logger(),
		spec:   spec,
	}, nil
}

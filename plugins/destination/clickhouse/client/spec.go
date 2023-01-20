package client

import (
	"crypto/x509"
	"database/sql/driver"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	CACert           string `json:"ca_cert,omitempty"`
}

func (s *Spec) Connector() (driver.Connector, error) {
	options, err := clickhouse.ParseDSN(s.ConnectionString)
	if err != nil {
		return nil, err
	}

	if tlsConfig := options.TLS; tlsConfig != nil && len(s.CACert) > 0 {
		// need to fill it the CACerts if required
		caCert, err := os.ReadFile(s.CACert)
		if err != nil {
			return nil, err
		}

		tlsConfig.RootCAs = x509.NewCertPool()
		tlsConfig.RootCAs.AppendCertsFromPEM(caCert)
	}

	return clickhouse.Connector(options), nil
}

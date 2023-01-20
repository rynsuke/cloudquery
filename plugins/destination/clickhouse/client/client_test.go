package client

import (
	"net/url"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_CH_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return (&url.URL{
		User: url.UserPassword("cq", "test"),
		Host: "localhost:9000",
	}).String()
}

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("clickhouse", plugin.Version, New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{ConnectionString: getTestConnection()},
		destination.PluginTestSuiteTests{},
	)
}

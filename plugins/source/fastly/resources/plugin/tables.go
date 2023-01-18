// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/account"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/auth"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/services"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/resources/services/stats"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		account.AccountUsers(),
		account.AccountEvents(),
		auth.AuthTokens(),
		services.Services(),
		stats.StatsRegions(),
		stats.StatsServices(),
	}
}

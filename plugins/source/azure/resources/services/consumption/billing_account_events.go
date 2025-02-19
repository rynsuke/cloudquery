package consumption

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func BillingAccountEvents() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_billing_account_events",
		Resolver:    fetchBillingAccountEvents,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/events/list-by-billing-account?tabs=HTTP#eventsummary",
		Multiplex:   client.BillingAccountMultiplex,
		Transform:   transformers.TransformWithStruct(&armconsumption.EventSummary{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewEventsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListByBillingAccountPager(*cl.BillingAccount.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			var respError *azcore.ResponseError
			// If there's no data a 404 error is returned so we ignore it
			if errors.As(err, &respError) && respError.StatusCode == 404 {
				cl.Logger().Debug().Msg("No data for billing account events")
				return nil
			}
			return err
		}
		res <- p.Value
	}
	return nil
}

package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ApplicationSecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_application_security_groups",
		Resolver:    fetchApplicationSecurityGroups,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/application-security-groups/list?tabs=HTTP#applicationsecuritygroup",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_application_security_groups", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.ApplicationSecurityGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchApplicationSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewApplicationSecurityGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

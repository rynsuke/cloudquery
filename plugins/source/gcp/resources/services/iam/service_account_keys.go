// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ServiceAccountKeys() *schema.Table {
	return &schema.Table{
		Name:      "gcp_iam_service_account_keys",
		Resolver:  fetchServiceAccountKeys,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "service_account_unique_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("unique_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyAlgorithm"),
			},
			{
				Name:     "key_origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyOrigin"),
			},
			{
				Name:     "key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyType"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "private_key_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateKeyData"),
			},
			{
				Name:     "private_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateKeyType"),
			},
			{
				Name:     "public_key_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicKeyData"),
			},
			{
				Name:     "valid_after_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValidAfterTime"),
			},
			{
				Name:     "valid_before_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValidBeforeTime"),
			},
		},
	}
}

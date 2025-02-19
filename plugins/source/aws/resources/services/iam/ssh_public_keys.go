package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func sshPublicKeys() *schema.Table {
	tableName := "aws_iam_ssh_public_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SSHPublicKeyMetadata.html`,
		Resolver:    fetchIamSshPublicKeys,
		Transform:   transformers.TransformWithStruct(&types.SSHPublicKeyMetadata{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
			{
				Name:     "ssh_public_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SSHPublicKeyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
func fetchIamSshPublicKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	paginator := iam.NewListSSHPublicKeysPaginator(svc, &iam.ListSSHPublicKeysInput{
		UserName: parent.Item.(*types.User).UserName,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SSHPublicKeys
	}
	return nil
}

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func DbSecurityGroups() *schema.Table {
	tableName := "aws_rds_db_security_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html`,
		Resolver:    fetchRdsDbSecurityGroups,
		Transform: transformers.TransformWithStruct(
			&types.DBSecurityGroup{},
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"e_c2": "ec2"})),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSecurityGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsDbSecurityGroupTags,
			},
		},
	}
}

func fetchRdsDbSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	var input rds.DescribeDBSecurityGroupsInput
	paginator := rds.NewDescribeDBSecurityGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBSecurityGroups
	}
	return nil
}

func resolveRdsDbSecurityGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBSecurityGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBSecurityGroupArn}, func(options *rds.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}

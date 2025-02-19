package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_elasticache_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html`,
		Resolver:    fetchElasticacheClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.CacheCluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveClusterTags,
			},
		},
	}
}

func fetchElasticacheClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeCacheClustersInput
	input.ShowCacheNodeInfo = aws.Bool(true)
	cl := meta.(*client.Client)
	paginator := elasticache.NewDescribeCacheClustersPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.CacheClusters
	}
	return nil
}

func resolveClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.CacheCluster)

	cl := meta.(*client.Client)
	svc := cl.Services().Elasticache
	response, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: cluster.ARN,
	}, func(options *elasticache.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.TagList))
}

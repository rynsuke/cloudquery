package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/digitalocean/godo"
)

func fetchProjectsResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(godo.Project)
	svc := meta.(*client.Client)

	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}

	done := false
	listFunc := func() error {
		data, resp, err := svc.Services.Projects.ListResources(ctx, p.ID, opt)
		if err != nil {
			return err
		}
		// pass the current page's data to our result channel
		res <- data
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			done = true
			return nil
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
		return nil
	}

	for !done {
		err := client.ThrottleWrapper(ctx, svc, listFunc)
		if err != nil {
			return err
		}
	}
	return nil
}

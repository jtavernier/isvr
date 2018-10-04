package client

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/jtavernier/isvr/cli/types"
)

//ResourceSave create a resource or update it if it already exists
func (cli *Client) ResourceSave(ctx context.Context, resource types.Resource, out io.Writer) error {
	query := url.Values{}
	resourcePath := "/resources/" + resource.ID

	resp, err := cli.head(ctx, resourcePath, query, nil)
	if resp.statusCode != 404 && err != nil {
		return err
	}

	if resp.statusCode == 200 {
		fmt.Fprintf(out, "Updating '%v'...", resource.ID)

		_, err := cli.put(ctx, resourcePath, nil, resource, nil)
		if err != nil {
			fmt.Fprintf(out, "ERROR\n")
			return err
		}

		fmt.Fprintf(out, "SUCCESS\n")

	} else {
		fmt.Fprintf(out, "Creating '%v'...", resource.ID)

		_, err := cli.post(ctx, "/resources", nil, resource, nil)
		if err != nil {
			fmt.Fprintf(out, "ERROR\n")
			return err
		}

		fmt.Fprintf(out, "SUCCESS\n")

	}

	return nil
}

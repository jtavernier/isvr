package client

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/jtavernier/isvr/cli/types"
)

//ClientSave create a client or update it if it already exists
func (cli *Client) ClientSave(ctx context.Context, client types.Client, out io.Writer) error {
	query := url.Values{}
	resourcePath := "/clients/" + client.ID

	resp, err := cli.head(ctx, resourcePath, query, nil)
	if resp.statusCode != 404 && err != nil {
		return err
	}

	if resp.statusCode == 200 {
		fmt.Fprintf(out, "Updating '%v'...", client.ID)

		_, err := cli.put(ctx, resourcePath, nil, client, nil)
		if err != nil {
			fmt.Fprintf(out, "ERROR\n")
			return err
		}

		fmt.Fprintf(out, "SUCCESS\n")

	} else {
		fmt.Fprintf(out, "Creating '%v'...", client.ID)

		_, err := cli.post(ctx, "/clients", nil, client, nil)
		if err != nil {
			fmt.Fprintf(out, "ERROR\n")
			return err
		}

		fmt.Fprintf(out, "SUCCESS\n")

	}

	return nil
}

package client

import (
	"context"
	"net/url"
)

//Health returns checks health of the ConfigurationAPI
func (cli *Client) Health(ctx context.Context) error {
	query := url.Values{}

	_, err := cli.get(ctx, "/checkHealth", query, nil)
	if err != nil {
		return err
	}

	return nil
}

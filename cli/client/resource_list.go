package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jtavernier/isvr/cli/types"
)

//ResourceList returns the list of resource.
func (cli *Client) ResourceList(ctx context.Context) ([]types.Resource, error) {
	query := url.Values{}

	resp, err := cli.get(ctx, "/resources", query, nil)
	if err != nil {
		return nil, err
	}

	var resources []types.Resource
	err = json.NewDecoder(resp.body).Decode(&resources)

	ensureReaderClosed(resp)
	return resources, err
}

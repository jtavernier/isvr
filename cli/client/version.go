package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jtavernier/isvr/cli/types"
)

//Version returns version of the ConfigurationAPI
func (cli *Client) Version(ctx context.Context) (types.Version, error) {
	query := url.Values{}
	var serverVersion types.Version

	resp, err := cli.get(ctx, "/version", query, nil)
	if err != nil {
		return serverVersion, err
	}

	err = json.NewDecoder(resp.body).Decode(&serverVersion)

	ensureReaderClosed(resp)
	return serverVersion, err
}

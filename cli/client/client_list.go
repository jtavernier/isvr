package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/jtavernier/isvr/cli/types"
)

//ClientList returns the list of clients.
func (cli *Client) ClientList(ctx context.Context) ([]types.Client, error) {
	query := url.Values{}

	resp, err := cli.get(ctx, "/clients", query, nil)
	if err != nil {
		return nil, err
	}

	var clients []types.Client
	err = json.NewDecoder(resp.body).Decode(&clients)

	ensureReaderClosed(resp)
	return clients, err
}

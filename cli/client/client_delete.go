package client

import "context"

//ClientDelete delete an existing client
func (cli *Client) ClientDelete(ctx context.Context, clientID string) (statuscode int, err error) {
	resourcePath := "/clients/" + clientID
	var resp serverResponse

	resp, err = cli.delete(ctx, resourcePath, nil, nil)
	statuscode = resp.statusCode

	return
}

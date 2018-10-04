package client

import "context"

//ResourceDelete delete an existing resource
func (cli *Client) ResourceDelete(ctx context.Context, resourceID string) (statuscode int, err error) {
	resourcePath := "/resources/" + resourceID
	var resp serverResponse

	resp, err = cli.delete(ctx, resourcePath, nil, nil)
	statuscode = resp.statusCode

	return
}

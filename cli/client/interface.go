package client

import (
	"context"
	"io"

	"github.com/jtavernier/isvr/cli/types"
)

//APIClient defines Configuration API client methods
type APIClient interface {
	ClientList(ctx context.Context) ([]types.Client, error)
	ResourceList(ctx context.Context) ([]types.Resource, error)
	ResourceSave(ctx context.Context, resource types.Resource, out io.Writer) error
	ResourceDelete(ctx context.Context, resourceID string) (statuscode int, err error)
	ClientSave(ctx context.Context, client types.Client, out io.Writer) error
	ClientDelete(ctx context.Context, clientID string) (statuscode int, err error)
	Version(ctx context.Context) (types.Version, error)
	Health(ctx context.Context) error
	GetHost() string
}

// Ensure that Client always implements APIClient.
var _ APIClient = &Client{}

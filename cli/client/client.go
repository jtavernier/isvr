package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

const (
	//DefaultConfigAPI is The default configuration API URL
	DefaultConfigAPI = "http://localhost:5020"
)

// ErrRedirect is the error returned by checkRedirect when the request is non-GET.
var ErrRedirect = errors.New("unexpected redirect in response")

//Client is the API client that performs all operations
// against Configuration API
type Client struct {
	// scheme sets the scheme for the client
	scheme string
	// host holds the server address to connect to
	host string
	// client used to send and receive http requests
	client *http.Client
	// addr holds the client address.
	addr string
	// basePath holds the path to prepend to the requests.
	basePath string
}

// GetHost returne the server address use by the client
func (c *Client) GetHost() string {
	return c.host
}

// NewClientWithOpts initializes a new API Client with default values.
func NewClientWithOpts(ops ...func(*Client) error) (*Client, error) {
	var url *url.URL
	var err error

	if host := os.Getenv("ISVR_HOST"); host == "" {
		url, err = ParseHostURL(DefaultConfigAPI)
		if err != nil {
			return nil, err
		}
	} else {
		url, err = ParseHostURL(host)
		if err != nil {
			return nil, err
		}
	}

	client := http.DefaultClient

	c := &Client{
		host:   url.Host,
		scheme: url.Scheme,
		client: client,
	}

	return c, nil
}

// ParseHostURL parses a url string, validates the string is a host url, and
// returns the parsed URL
func ParseHostURL(host string) (*url.URL, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return nil, fmt.Errorf("unable to parse host `%s` should be 'https://host:port'", host)
	}

	var basePath string
	proto, addr := protoAddrParts[0], protoAddrParts[1]

	return &url.URL{
		Scheme: proto,
		Host:   addr,
		Path:   basePath,
	}, nil
}

// CheckRedirect specifies the policy for dealing with redirect responses:
// If the request is non-GET return `ErrRedirect`. Otherwise use the last response.
func CheckRedirect(req *http.Request, via []*http.Request) error {
	if via[0].Method == http.MethodGet {
		return http.ErrUseLastResponse
	}
	return ErrRedirect
}

// getAPIPath returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func (cli *Client) getAPIPath(p string, query url.Values) string {
	var apiPath string
	apiPath = path.Join(cli.basePath, "/api", p)

	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}
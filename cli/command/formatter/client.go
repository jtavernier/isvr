package formatter

import (
	"strconv"
	"strings"

	"github.com/jtavernier/isvr/cli/types"
)

const (
	defaultClientTableFormat = "table {{.ID}}\t{{.Name}}\t{{.NbSecrets}}\t{{.GrantTypes}}\t{{.Scopes}}"

	clientIDHeader         = "ID"
	clientNameHeader       = "Name"
	clientScopesHeader     = "Scope"
	clientGrantTypesHeader = "Grant Types"
	clientSecretsHeader    = "Nb Secrets"
)

// NewClientFormat returns a Format for rendering using a Context
func NewClientFormat(source string, quiet bool) Format {
	switch source {
	case TableFormatKey:
		if quiet {
			return defaultQuietFormat
		}
		format := defaultClientTableFormat
		return Format(format)
	}

	return Format(source)
}

// ClientWrite renders the context for a list of clients
func ClientWrite(ctx Context, clients []types.Client) error {
	render := func(format func(subContext subContext) error) error {
		for _, client := range clients {
			err := format(&clientContext{c: client})
			if err != nil {
				return err
			}
		}
		return nil
	}
	return ctx.Write(newClientContext(), render)
}

type clientHeaderContext map[string]string

type clientContext struct {
	HeaderContext
	clientHeaderContext
	c types.Client
}

func newClientContext() *clientContext {
	clientCtx := clientContext{}
	clientCtx.header = clientHeaderContext{
		"ID":         clientIDHeader,
		"Name":       clientNameHeader,
		"NbSecrets":  clientSecretsHeader,
		"GrantTypes": clientGrantTypesHeader,
		"Scopes":     clientScopesHeader,
	}

	return &clientCtx
}

func (c *clientContext) ID() string {
	return c.c.ID
}

func (c *clientContext) Name() string {
	return c.c.Name
}

func (c *clientContext) GrantTypes() string {
	return strings.Join(c.c.AllowedGrantTypes, ",")
}

func (c *clientContext) Scopes() string {
	return strings.Join(c.c.AllowedScopes, ",")
}

func (c *clientContext) NbSecrets() string {
	return strconv.Itoa(len(c.c.Secrets))
}

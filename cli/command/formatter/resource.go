package formatter

import (
	"strconv"
	"strings"

	"github.com/jtavernier/isvr/cli/types"
)

const (
	defaultResourceTableFormat = "table {{.ID}}\t{{.Description}}\t{{.NbSecrets}}\t{{.Scopes}}"

	resourceIDHeader          = "ID"
	resourceDescriptionHeader = "Description"
	resourceScopesHeader      = "Scopes"
	resourceSecretsHeader     = "Nb Secrets"
)

// NewResourceFormat returns a Format for rendering using a Context
func NewResourceFormat(source string, quiet bool) Format {
	switch source {
	case TableFormatKey:
		if quiet {
			return defaultQuietFormat
		}
		format := defaultResourceTableFormat
		return Format(format)
	}

	return Format(source)
}

//ResourceWrite renders the context for a list of resources
func ResourceWrite(ctx Context, resources []types.Resource) error {
	render := func(format func(subContext subContext) error) error {
		for _, resource := range resources {
			err := format(&resourceContext{r: resource})
			if err != nil {
				return err
			}
		}
		return nil
	}
	return ctx.Write(newResourceContext(), render)
}

type resourceHeaderContext map[string]string

type resourceContext struct {
	HeaderContext
	resourceHeaderContext
	r types.Resource
}

func newResourceContext() *resourceContext {
	resourceCtx := resourceContext{}
	resourceCtx.header = resourceHeaderContext{
		"ID":          resourceIDHeader,
		"Description": resourceDescriptionHeader,
		"NbSecrets":   resourceSecretsHeader,
		"Scopes":      resourceScopesHeader,
	}

	return &resourceCtx
}

func (r *resourceContext) ID() string {
	return r.r.ID
}

func (r *resourceContext) Description() string {
	return r.r.Description
}

func (r *resourceContext) Scopes() string {
	var results []string

	for _, scope := range r.r.Scopes {
		results = append(results, scope.Name)
	}

	return strings.Join(results, ",")
}

func (r *resourceContext) NbSecrets() string {
	return strconv.Itoa(len(r.r.Secrets))
}

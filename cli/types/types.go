package types

// Config is a full auth configuration file
type Config struct {
	Resources []Resource `yaml:",omitempty"`
	Clients   []Client   `yaml:",omitempty"`
}

// Client the configuration for an identity server resource
type Client struct {
	ID                     string   `yaml:",omitempty" json:"clientId,omitempty"`
	Name                   string   `yaml:",omitempty" json:"clientName,omitempty"`
	AllowedGrantTypes      []string `yaml:"allowed_grant_types,omitempty" json:"allowedGrantTypes,omitempty"`
	Secrets                []string `yaml:",omitempty" json:"clientSecrets,omitempty"`
	AllowedScopes          []string `yaml:"allowed_scopes,omitempty" json:"allowedScopes,omitempty"`
	RedirectUris           []string `yaml:"redirect_uris,omitempty" json:"redirectUris,omitempty"`
	PostLogoutRedirectUris []string `yaml:"post_logout_redirect_uris,omitempty" json:"postLogoutRedirectUris,omitempty"`
}

// Resource the configuration for an identity server resource
type Resource struct {
	ID          string   `yaml:"name,omitempty" json:"name,omitempty"`
	Description string   `yaml:",omitempty" json:"description,omitempty"`
	Secrets     []string `yaml:",omitempty" json:"apiSecrets,omitempty"`
	Scopes      []Scope  `yaml:",omitempty" json:"scopes,omitempty"`
}

// Scope the configuration for an identity server scope
type Scope struct {
	Name        string `yaml:",omitempty" json:"name,omitempty"`
	Description string `yaml:",omitempty" json:"description,omitempty"`
}

//Version the Version of the Configuration API
type Version struct {
	Version string `json:"version,omitempty"`
}

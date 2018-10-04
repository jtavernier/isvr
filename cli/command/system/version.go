package system

import (
	"context"
	"text/tabwriter"
	"text/template"

	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/jtavernier/isvr/cli/templates"
	"github.com/jtavernier/isvr/cli/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var versionTemplate = `
CLI:
 Version:	{{.Client.Version}}
{{if .ServerOK}}
Configuration API:
 Version:	{{.Server.Version}}
{{- end}}`

// NewVersionCommand returns a cobra command for `info` subcommands
func NewVersionCommand(isvrCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display versions of the CLI and Server",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(isvrCli)
		},
	}

	return cmd
}

type versionInfo struct {
	Client *types.Version
	Server *types.Version
}

func (v versionInfo) ServerOK() bool {
	return v.Server != nil
}

func runVersion(isvrCli command.Cli) error {

	ctx := context.Background()
	version := versionInfo{}

	serverVersion, err := isvrCli.Client().Version(ctx)
	if err != nil {
		version.Server = nil
	} else {
		version.Server = &serverVersion
	}

	version.Client = &types.Version{Version: cli.Version}

	tmpl, err := newVersionTemplate()
	if err != nil {
		return err
	}

	t := tabwriter.NewWriter(isvrCli.Out(), 20, 1, 1, ' ', 0)
	err = tmpl.Execute(t, version)
	t.Write([]byte("\n"))
	t.Flush()

	return err
}

func newVersionTemplate() (*template.Template, error) {
	templateFormat := versionTemplate

	tmpl := templates.New("version")
	tmpl, err := tmpl.Parse(templateFormat)

	return tmpl, errors.Wrap(err, "Template parsing error")
}

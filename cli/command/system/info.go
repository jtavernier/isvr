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

var infoTemplate = `
CLI:
 Version:	{{.Version.Client.Version}}

Configuration API:
 Host:	{{.ServerHost}}
 Health:	{{.ServerHealth}}{{if .Version.ServerOK}}
 Version:	{{.Version.Server.Version}}{{- end}}
 `

// NewInfoCommand returns a cobra command for `info` subcommands
func NewInfoCommand(isvrCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Display system-wide information",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInfo(isvrCli)
		},
	}

	return cmd
}

type systemInfo struct {
	Version      *versionInfo
	ServerHost   string
	ServerHealth string
}

func runInfo(isvrCli command.Cli) error {

	ctx := context.Background()
	info := systemInfo{}
	version := versionInfo{}
	info.Version = &version

	serverVersion, err := isvrCli.Client().Version(ctx)
	if err != nil {
		info.Version.Server = nil
	} else {
		info.Version.Server = &serverVersion
	}

	info.Version.Client = &types.Version{Version: cli.Version}

	if info.Version.Server == nil {
		info.ServerHealth = "NO RESPONSE"
	} else {
		err = isvrCli.Client().Health(ctx)
		if err != nil {
			info.ServerHealth = "NOT OK"
		} else {
			info.ServerHealth = "OK"
		}
	}

	info.ServerHost = isvrCli.Client().GetHost()

	tmpl, err := newInfoTemplate()
	if err != nil {
		return err
	}

	t := tabwriter.NewWriter(isvrCli.Out(), 20, 1, 1, ' ', 0)
	err = tmpl.Execute(t, info)
	t.Write([]byte("\n"))
	t.Flush()

	return err
}

func newInfoTemplate() (*template.Template, error) {
	templateFormat := infoTemplate

	tmpl := templates.New("info")
	tmpl, err := tmpl.Parse(templateFormat)

	return tmpl, errors.Wrap(err, "Template parsing error")
}

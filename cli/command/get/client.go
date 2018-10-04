package get

import (
	"context"

	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/jtavernier/isvr/cli/command/formatter"
	"github.com/spf13/cobra"
)

type clientOptions struct {
	clients []string
	quiet   bool
	format  string
}

//NewClientCommand create a new cobra.Command for `isvr get client`
func NewClientCommand(isvrCli command.Cli) *cobra.Command {
	options := clientOptions{}

	cmd := &cobra.Command{
		Use:     "client",
		Aliases: []string{"clients"},
		Short:   "List clients",
		Args:    cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			options.clients = args
			return getClients(isvrCli, &options)
		},
	}

	flags := cmd.Flags()

	flags.BoolVarP(&options.quiet, "quiet", "q", false, "Only display IDs")
	flags.StringVarP(&options.format, "format", "", "", "Pretty-print containers using a Go template")

	return cmd
}

func getClients(isvrCli command.Cli, options *clientOptions) error {
	ctx := context.Background()

	clients, err := isvrCli.Client().ClientList(ctx)
	if err != nil {
		return err
	}

	format := options.format
	if len(format) == 0 {
		format = formatter.TableFormatKey
	}

	clientCtx := formatter.Context{
		Output: isvrCli.Out(),
		Format: formatter.NewClientFormat(format, options.quiet),
	}

	return formatter.ClientWrite(clientCtx, clients)
}

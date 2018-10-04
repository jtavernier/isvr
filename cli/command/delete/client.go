package delete

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/spf13/cobra"
)

type clientOptions struct {
	clients []string
}

// NewClientCommand creates a new cobra.command for `isvr delete client`
func NewClientCommand(isvrCli command.Cli) *cobra.Command {
	options := clientOptions{}

	cmd := &cobra.Command{
		Use:     "client CLIENT [CLIENT...]",
		Aliases: []string{"clients"},
		Short:   "Delete one or more clients",
		Args:    cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.clients = args
			return deleteClients(isvrCli, &options)
		},
	}

	return cmd
}

func deleteClients(isvrCli command.Cli, options *clientOptions) error {
	ctx := context.Background()

	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)

	for _, client := range options.clients {
		fmt.Fprintf(isvrCli.Out(), "Deleting '%v'...", client)
		statusCode, err := isvrCli.Client().ClientDelete(ctx, client)

		if statusCode == 404 {
			yellow.Fprintf(isvrCli.Out(), "NOT FOUND\n")
		} else if err != nil {
			red.Fprintf(isvrCli.Out(), "ERROR\n")
			return err
		} else {
			green.Fprintf(isvrCli.Out(), "SUCCESS\n")
		}
	}

	return nil
}

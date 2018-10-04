package delete

import (
	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/spf13/cobra"
)

// NewDeleteCommand returns a cobra command for `delete` subcommands
func NewDeleteCommand(isvrCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete on or more configuration elements",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(isvrCli.Out()),
	}

	cmd.AddCommand(
		NewClientCommand(isvrCli),
		NewResourceCommand(isvrCli),
	)

	return cmd
}

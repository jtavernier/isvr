package get

import (
	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/spf13/cobra"
)

// NewGetCommand returns a cobra command for `get` subcommands
func NewGetCommand(IsvrCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "List configuration elements",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(IsvrCli.Out()),
	}

	cmd.AddCommand(
		NewClientCommand(IsvrCli),
		NewResourceCommand(IsvrCli),
	)
	return cmd
}

package commands

import (
	"github.com/jtavernier/isvr/cli/command"
	"github.com/jtavernier/isvr/cli/command/apply"
	"github.com/jtavernier/isvr/cli/command/delete"
	"github.com/jtavernier/isvr/cli/command/get"
	"github.com/jtavernier/isvr/cli/command/system"
	"github.com/spf13/cobra"
)

// AddCommands adds alll the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, isvrCli command.Cli) {
	cmd.AddCommand(
		//get
		get.NewGetCommand(isvrCli),

		//apply
		apply.NewApplyCommand(isvrCli),

		//delete
		delete.NewDeleteCommand(isvrCli),

		//version
		system.NewVersionCommand(isvrCli),

		//info
		system.NewInfoCommand(isvrCli),
	)
}

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/jtavernier/isvr/cli/command/commands"
	"github.com/spf13/cobra"
)

func newIsvrCommand(isvrCli *command.IsvrCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isvr [OPTIONS] COMMAND [ARG]",
		Short: "Configuration manager for Identity Server 4",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := isvrCli.Initialize(); err != nil {
				return err
			}
			return nil
		},
	}

	cli.SetupRootCommand(cmd)
	commands.AddCommands(cmd, isvrCli)

	return cmd
}

func main() {
	var buf bytes.Buffer

	IsvrCli := command.NewIsvrCli(&buf)
	cmd := newIsvrCommand(IsvrCli)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

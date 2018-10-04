package cli

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SetupRootCommand sets default usage, help and error handling for the root command
func SetupRootCommand(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.SetHelpCommand(helpCommand)
}

var helpCommand = &cobra.Command{
	Use:               "help [command]",
	Short:             "Help about the command",
	PersistentPreRun:  func(cmd *cobra.Command, args []string) {},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {},
	RunE: func(c *cobra.Command, args []string) error {
		cmd, args, e := c.Root().Find(args)
		if cmd == nil || e != nil || len(args) > 0 {
			return errors.Errorf("unknown help topic: %v", strings.Join(args, " "))
		}

		helpFunc := cmd.HelpFunc()
		helpFunc(cmd, args)
		return nil
	},
}

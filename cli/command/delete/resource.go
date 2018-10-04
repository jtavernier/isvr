package delete

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command"
	"github.com/spf13/cobra"
)

type resourceOptions struct {
	resources []string
}

// NewResourceCommand creates a new cobra.command for `isvr delete resource`
func NewResourceCommand(isvrCli command.Cli) *cobra.Command {
	options := resourceOptions{}

	cmd := &cobra.Command{
		Use:     "resource RESOURCE [RESOURCE...]",
		Aliases: []string{"resources"},
		Short:   "Delete one or more resources",
		Args:    cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.resources = args
			return deleteResources(isvrCli, &options)
		},
	}

	return cmd
}

func deleteResources(isvrCli command.Cli, options *resourceOptions) error {
	ctx := context.Background()

	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)

	for _, resource := range options.resources {
		fmt.Fprintf(isvrCli.Out(), "Deleting '%v'...", resource)
		statusCode, err := isvrCli.Client().ResourceDelete(ctx, resource)

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

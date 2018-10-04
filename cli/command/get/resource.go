package get

import (
	"context"

	"github.com/jtavernier/isvr/cli"
	"github.com/jtavernier/isvr/cli/command/formatter"

	"github.com/jtavernier/isvr/cli/command"
	"github.com/spf13/cobra"
)

type resourceOptions struct {
	resources []string
	quiet     bool
	format    string
}

//NewResourceCommand create a new cobra.Command for 'isvr get resource'
func NewResourceCommand(isvrCli command.Cli) *cobra.Command {
	options := resourceOptions{}

	cmd := &cobra.Command{
		Use:     "resource",
		Aliases: []string{"resources"},
		Short:   "List API Resources",
		Args:    cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			options.resources = args
			return getResources(isvrCli, &options)
		},
	}

	flags := cmd.Flags()

	flags.BoolVarP(&options.quiet, "quiet", "q", false, "Only display IDs")
	flags.StringVarP(&options.format, "format", "", "", "Pretty-print containers using a Go template")

	return cmd
}

func getResources(isvrCli command.Cli, options *resourceOptions) error {
	ctx := context.Background()

	resources, err := isvrCli.Client().ResourceList(ctx)
	if err != nil {
		return err
	}

	format := options.format
	if len(format) == 0 {
		format = formatter.TableFormatKey
	}

	resourceCtx := formatter.Context{
		Output: isvrCli.Out(),
		Format: formatter.NewResourceFormat(format, options.quiet),
	}

	return formatter.ResourceWrite(resourceCtx, resources)
}

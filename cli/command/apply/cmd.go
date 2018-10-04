package apply

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/jtavernier/isvr/cli/command"
	"github.com/jtavernier/isvr/cli/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultFilePath = "../../samples/AuthConfiguration.yml"
)

type applyOptions struct {
	filePath string
}

// NewApplyCommand returns a cobra comand for `get` subcommands
func NewApplyCommand(isvrCli command.Cli) *cobra.Command {
	options := applyOptions{}

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a configuration to a resource by a filename",
		RunE: func(cmd *cobra.Command, args []string) error {
			return applyConfiguration(isvrCli, &options)
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&options.filePath, "file", "f", defaultFilePath, "Specify the path of the configuration file")

	return cmd
}

func applyConfiguration(isvrCli command.Cli, options *applyOptions) error {
	ctx := context.Background()
	configs, err := unmarshalConfigurationFile(options.filePath)
	if err != nil {
		return err
	}

	for _, resource := range configs.Resources {
		isvrCli.Client().ResourceSave(ctx, resource, isvrCli.Out())
	}

	for _, client := range configs.Clients {
		isvrCli.Client().ClientSave(ctx, client, isvrCli.Out())
	}

	return nil
}

func unmarshalConfigurationFile(configFilePath string) (types.Config, error) {
	configs := types.Config{}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Printf("ERROR while trying to read the file")
		return configs, err
	}

	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		return configs, err
	}

	return configs, nil
}

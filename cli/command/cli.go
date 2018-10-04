package command

import (
	"io"
	"os"

	"github.com/jtavernier/isvr/cli/client"
	"github.com/spf13/cobra"
)

// Cli represents the idsvr command line client
type Cli interface {
	Client() client.APIClient
	Err() io.Writer
	Out() io.Writer
}

//IsvrCli is an instance of the isvr command line client
type IsvrCli struct {
	client     client.APIClient
	serverInfo interface{}
	clientInfo interface{}
	out        io.Writer
	err        io.Writer
}

//Err returns the writer used for stderr
func (cli *IsvrCli) Err() io.Writer {
	return cli.err
}

//Out returns the writter used for stdout
func (cli *IsvrCli) Out() io.Writer {
	return cli.out
}

// NewIsvrCli returns a IsvrCli instance
func NewIsvrCli(err io.Writer) *IsvrCli {
	return &IsvrCli{err: err}
}

// Initialize the isvrCli
func (cli *IsvrCli) Initialize() error {
	var err error
	cli.client, err = client.NewClientWithOpts()
	if err != nil {
		return err
	}
	cli.out = os.Stdout

	return nil
}

// Client returns the APIClient
func (cli *IsvrCli) Client() client.APIClient {
	return cli.client
}

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOutput(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

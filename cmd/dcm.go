package main

import (
	"akc-dcm-cli/commands"
	"github.com/spf13/cobra"
	"os"
)

func NewDCMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dcm",
		Short: "The command line interface for digital certificate management (dcm)",
	}

	cmd.AddCommand(commands.All()...)

	return cmd
}

func main() {
	cmd := NewDCMCommand()
	flags := cmd.PersistentFlags()
	_ = flags.Parse(os.Args[1:])

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

package version

import (
	"akc-dcm-cli/glossary"
	"akc-dcm-cli/glossary/metadata"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

//NewVersionCommand creates a new "dcm version" command
func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print dcm command version",
		Long:  "Print current version of digital certificate management (dcm) command line tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("trailing args detected")
			}
			cmd.SilenceUsage = true
			fmt.Print(GetMetaInfo())
			return nil
		},
	}

	return cmd
}

//GetMetaInfo returns version information for the fabric cmd.
func GetMetaInfo() string {
	return fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go Version: %s\n"+
		" OS/Arch: %s\n",
		glossary.ProgramName, metadata.Version, metadata.CommitSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}

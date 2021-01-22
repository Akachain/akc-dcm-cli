package commands

import (
	"aka-dcm-cli/commands/certificate"
	"aka-dcm-cli/commands/version"
	"github.com/spf13/cobra"
)

func All() []*cobra.Command {
	return []*cobra.Command{
		certificate.NewCertificateCommand(),
		version.NewVersionCommand(),
	}
}

package commands

import (
	"akc-dcm-cli/commands/certificate"
	"akc-dcm-cli/commands/version"
	"github.com/spf13/cobra"
)

func All() []*cobra.Command {
	return []*cobra.Command{
		certificate.NewCertificateCommand(),
		version.NewVersionCommand(),
	}
}

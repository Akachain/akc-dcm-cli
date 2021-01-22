package check

import (
	"aka-dcm-cli/commands/common"
	"aka-dcm-cli/utilities"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"path/filepath"
	"sort"
	"time"
)

func NewExpireCommand() *cobra.Command {
	c := ExpireCommand{}

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Checking expiration date",
		Long:  "Checking expiration date of certificate",
		Args:  c.ParseArgs(),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			if err := c.Validate(); err != nil {
				return err
			}

			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return c.Run()
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&c.CertPath, "cert-path", "c", "", "Path to your certificate")
	flags.StringVarP(&c.FolderPath, "folder-cert", "f", "", "Path to folder have certificates")

	return cmd
}

type ExpireCommand struct {
	common.Command
	CertPath   string
	FolderPath string
}

// Validate checks the required parameters for run
func (c *ExpireCommand) Validate() error {
	if len(c.CertPath) == 0 && len(c.FolderPath) == 0 {
		return errors.New("File certificate or Folder have certificate is required")
	}

	return nil
}

// Run executes the command
func (c *ExpireCommand) Run() error {
	if len(c.CertPath) > 0 {
		return checkExpireCert(c.CertPath)
	} else {
		certificateFiles, err := filepath.Glob(filepath.Join(c.FolderPath, "*.pem"))
		if err != nil {
			return errors.WithMessage(err, "Failed to scan certificate dir")
		}
		sort.Strings(certificateFiles)
		for _, file := range certificateFiles {
			err := checkExpireCert(file)
			if err != nil {
				fmt.Println(fmt.Sprintf("Unable to check expire date of \"%s\" certificate", file))
				continue
			}
		}
	}
	return nil
}

func checkExpireCert(certPath string) error {
	cert, err := utilities.ParseCertificate(certPath)
	if err != nil {
		return err
	}

	fileName := filepath.Base(certPath)
	fmt.Println("Certificate", color.YellowString("%s", fileName), "will be expire at", color.YellowString("%s", cert.NotAfter.String()))
	if cert.NotAfter.Before(time.Now()) {
		fmt.Println("Certificate", color.RedString("%s was expired!!!", fileName))
	} else {
		fmt.Println("Certificate", color.GreenString("%s is good today.", fileName))
	}

	return nil
}

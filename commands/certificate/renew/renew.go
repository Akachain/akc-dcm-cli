package renew

import (
	"akc-dcm-cli/commands/common"
	"akc-dcm-cli/glossary"
	"akc-dcm-cli/utilities"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"time"
)

func NewReNewCertificateCommand() *cobra.Command {
	c := ReNewCommand{}

	cmd := &cobra.Command{
		Use:   "renew",
		Short: "Renew expiration date",
		Long:  "Renew expiration date of certificate",
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
	flags.StringVar(&c.ParenCertPath, "parent-cert", "", "Path to parent certificate (CA or ICA)")
	flags.StringVar(&c.ParentPrivKeyPath, "parent-private-key", "", "Path to private key of parent certificate (CA or ICA)")
	flags.StringVar(&c.OldCertPath, "old-cert", "", "Path to certificate of old certificate that need to renew")
	flags.StringVar(&c.PrivKeyOldCertPath, "old-private-key", "", "Path to private key of old certificate")
	flags.IntVarP(&c.Day, "days", "d", 1, "Number date expire of new certificate")
	flags.StringVarP(&c.OutputPath, "output", "o", ".dcm/output/renew-cert.pem", "Path to output file of new certificate")

	return cmd
}

type ReNewCommand struct {
	common.Command
	ParenCertPath      string
	ParentPrivKeyPath  string
	OldCertPath        string
	PrivKeyOldCertPath string
	Day                int
	OutputPath         string
}

// Validate checks the required parameters for run
func (c *ReNewCommand) Validate() error {
	if len(c.ParenCertPath) == 0 {
		return errors.New("CA Certificate path is required")
	}

	if len(c.ParentPrivKeyPath) == 0 {
		return errors.New("Private key of CA path is required")
	}

	if len(c.OldCertPath) == 0 {
		return errors.New("Old Certificate path is required")
	}

	if len(c.PrivKeyOldCertPath) == 0 {
		return errors.New("Private key of Old Certificate path is required")
	}

	if c.Day <= 0 {
		return errors.New("Number of date expire must be greater than zero")
	}

	return nil
}

// Run executes the command
func (c *ReNewCommand) Run() error {
	// parse ca information
	parentPrivKey, err := utilities.ParsePrivateKey(c.ParentPrivKeyPath)
	if err != nil {
		return err
	}

	parentCert, _, err := utilities.ParseCertificate(c.ParenCertPath)
	if err != nil {
		return err
	}

	// verify private/public key pair of parent cert
	if err := utilities.VerifyKey(parentPrivKey, parentCert.PublicKey); err != nil {
		return err
	}

	// parse old certificate information
	privateKey, err := utilities.ParsePrivateKey(c.PrivKeyOldCertPath)
	if err != nil {
		return err
	}

	oldCert, isJson, err := utilities.ParseCertificate(c.OldCertPath)
	if err != nil {
		return err
	}

	// renew certificate
	oldCert.NotAfter = time.Now().AddDate(0, 0, c.Day)

	// verify private/public key pair of children cert
	if err := utilities.VerifyKey(privateKey, oldCert.PublicKey); err != nil {
		return err
	}

	// work around for fabric extension
	for _, ext := range oldCert.Extensions {
		if ext.Id.Equal(glossary.FabricComment) {
			oldCert.ExtraExtensions = append(oldCert.ExtraExtensions, ext)
		}
	}

	// generate new certificate
	newCert, err := utilities.GenerateCertificate(oldCert, parentCert, privateKey, parentPrivKey)
	if err != nil {
		return err
	}

	// handle support json format cert
	if !isJson {
		err = utilities.WriteFileToLocal(c.OutputPath, newCert.Bytes())
		if err != nil {
			return err
		}
	} else {
		jCert, err := utilities.ParseJsonCert(c.OldCertPath)
		if err != nil {
			return err
		}
		jCert.Enrollment.Identity.Certificate = newCert.String()
		err = utilities.WriteJsonFileToLocal(c.OutputPath, jCert)
		if err != nil {
			return err
		}
	}

	color.Green("Renew Certificate Success...")

	return nil
}

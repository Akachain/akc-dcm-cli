package utilities

import (
	"akc-dcm-cli/glossary"
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// ParsePrivateKey return crypto.PrivateKey from file
func ParsePrivateKey(privateKeyPath string) (crypto.PrivateKey, error) {
	privKeyByte, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, errors.WithMessage(err, "ParsePrivateKey - Read private key file error")
	}

	privKeyBlock, _ := pem.Decode(privKeyByte)
	if privKeyBlock == nil {
		return nil, errors.New("ParsePrivateKey - Unable to convert block of private key")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privKeyBlock.Bytes); err != nil {
			return nil, errors.WithMessage(err, "ParsePrivateKey - Unable to parse private key")
		}
	}

	return parsedKey, nil
}

// ParsePrivateKey return Certificate from file
func ParseCertificate(certPath string) (*x509.Certificate, error) {
	certByte, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, errors.WithMessage(err, "ParseCertificate - Read certificate file error")
	}

	blockCert, _ := pem.Decode(certByte)
	if blockCert == nil {
		return nil, errors.New("ParseCertificate - Unable to convert block of certificate")
	}

	cert, err := x509.ParseCertificate(blockCert.Bytes)
	if err != nil {
		return nil, errors.WithMessage(err, "ParseCertificate - Parse certificate error")
	}
	return cert, nil
}

// GenerateCertificate return new certificate from template
func GenerateCertificate(template *x509.Certificate, parentCert *x509.Certificate,
	cryptoPrivate crypto.PrivateKey, parentPrivateKey crypto.PrivateKey) (output *bytes.Buffer, err error) {
	var certBytes []byte
	switch cryptoPrivate.(type) {
	case *rsa.PrivateKey:
		privateKey, ok := cryptoPrivate.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("GenerateCertificate - Unable to convert to type RSA")
		}
		certBytes, err = x509.CreateCertificate(rand.Reader, template, parentCert, privateKey.Public(), parentPrivateKey)
		if err != nil {
			return nil, errors.WithMessage(err, "Unable to generate certificate")
		}
		break
	case *ecdsa.PrivateKey:
		privateKey, ok := cryptoPrivate.(*ecdsa.PrivateKey)
		if !ok {
			return nil, errors.New("GenerateCertificate - Unable to convert to type RSA")
		}
		certBytes, err = x509.CreateCertificate(rand.Reader, template, parentCert, privateKey.Public(), parentPrivateKey)
		if err != nil {
			return nil, errors.WithMessage(err, "Unable to generate certificate")
		}
		break
	default:
		return nil, errors.New("GenerateCertificate - Unsupported type private key")
	}

	output = new(bytes.Buffer)
	err = pem.Encode(output, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	return output, err
}

func WriteFileToLocal(filePath string, data []byte) error {
	usr, err := user.Current()
	if err != nil {
		return errors.WithMessage(err, "WriteFileToLocal - Unable to get current user")
	}

	if _, err := os.Stat(usr.HomeDir + "/" + glossary.DefaultOutputPath); os.IsNotExist(err) {
		err := os.MkdirAll(usr.HomeDir+"/"+glossary.DefaultOutputPath, 0755)
		return errors.WithMessage(err, "WriteFileToLocal - Unable to make default output folder")
	}

	if strings.Contains(filePath, ".dcm") {
		filePath = usr.HomeDir + "/" + filePath
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return errors.WithMessage(err, "Unable to write certificate")
	}
	return nil
}

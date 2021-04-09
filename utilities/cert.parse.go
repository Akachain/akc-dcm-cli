package utilities

import (
	"akc-dcm-cli/glossary/jsonCert"
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
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
func ParseCertificate(certPath string) (cert *x509.Certificate, isJson bool, err error) {
	isJson = false
	var certByte []byte
	if IsJsonCert(certPath) {
		isJson = true
		jCert, err := ParseJsonCert(certPath)
		if err != nil {
			return nil, isJson, errors.WithMessage(err, "ParseCertificate - Parse json cert error")
		}
		certByte = []byte(jCert.Enrollment.Identity.Certificate)
	} else {
		certByte, err = ioutil.ReadFile(certPath)
		if err != nil {
			return nil, isJson, errors.WithMessage(err, "ParseCertificate - Read certificate file error")
		}
	}

	blockCert, _ := pem.Decode(certByte)
	if blockCert == nil {
		return nil, isJson, errors.New("ParseCertificate - Unable to convert block of certificate")
	}

	cert, err = x509.ParseCertificate(blockCert.Bytes)
	if err != nil {
		return nil, isJson, errors.WithMessage(err, "ParseCertificate - Parse certificate error")
	}
	return cert, isJson, nil
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

// ParseJsonCert parse json cert with format was supported
func ParseJsonCert(jsonCertPath string) (*jsonCert.JCert, error) {
	certByte, err := ioutil.ReadFile(jsonCertPath)
	if err != nil {
		return nil, errors.WithMessage(err, "ParseJsonCert - Read certificate file error")
	}
	js := new(jsonCert.JCert)
	err = json.Unmarshal(certByte, &js)
	return js, err
}

// IsJsonCert check cert is json format or not
func IsJsonCert(jsonCertPath string) bool {
	if filepath.Ext(strings.TrimSpace(jsonCertPath)) == ".jsonCert" {
		return true
	}
	certByte, err := ioutil.ReadFile(jsonCertPath)
	if err != nil {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal(certByte, &js) == nil
}

func VerifyKey(privateKey crypto.PrivateKey, publicKey crypto.PublicKey) error {
	switch privateKey.(type) {
	case *rsa.PrivateKey:
		privateKey, ok := privateKey.(*rsa.PrivateKey)
		if !ok {
			return errors.New("VerifyKey - Unable to convert to type RSA")
		}

		if !privateKey.Public().(*rsa.PublicKey).Equal(publicKey) {
			return errors.New("VerifyKey - Private key and public key not a pair")
		}
		break
	case *ecdsa.PrivateKey:
		privateKey, ok := privateKey.(*ecdsa.PrivateKey)
		if !ok {
			return errors.New("VerifyKey - Unable to convert to type RSA")
		}

		if !privateKey.Public().(*ecdsa.PublicKey).Equal(publicKey) {
			return errors.New("Private key and public key not match")
		}
		break
	default:
		return errors.New("VerifyKey - Unsupported type private key")
	}
	return nil
}

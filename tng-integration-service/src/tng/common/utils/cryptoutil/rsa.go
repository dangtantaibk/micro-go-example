package cryptoutil

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// ParseToPublicKey will parse pubPem string to RSA Public Key
func ParseToPublicKey(pubPem string) (interface{}, error) {

	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid rsa public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		return pub, nil
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

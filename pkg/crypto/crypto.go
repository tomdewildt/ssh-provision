package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/ssh"
)

// GenerateKeypair is used to generate a new ssh private and public
// keypair. It takes the amount of bits that should be used to generate
// the private key as input and returns the private key, the public and
// nil or nil, nil and an error if one occurred.
func GenerateKeypair(bits int) ([]byte, []byte, error) {
	if bits == 0 || bits%2 != 0 {
		return nil, nil, errors.New("invalid bits count")
	}

	var private bytes.Buffer
	var public []byte

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	privatePEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(&private, privatePEM); err != nil {
		return nil, nil, err
	}

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	public = ssh.MarshalAuthorizedKey(publicKey)

	return private.Bytes(), public, nil
}

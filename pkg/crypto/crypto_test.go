package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
)

func TestGenerateKeypairInvalidBits(t *testing.T) {
	_, _, err := GenerateKeypair(0)

	assert.EqualError(t, err, "invalid bits count", "Error should be invalid bits count")

	_, _, err = GenerateKeypair(3)

	assert.EqualError(t, err, "invalid bits count", "Error should be invalid bits count")
}

func TestGenerateKeypair(t *testing.T) {
	privateKey, publicKey, err := GenerateKeypair(2048)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, privateKey, "Private key should not be nil")
	assert.NotNil(t, publicKey, "Public key should not be nil")

	_, err = ssh.ParsePrivateKey(privateKey)

	assert.Nil(t, err, "Error should be nil")

	_, _, _, _, err = ssh.ParseAuthorizedKey(publicKey)

	assert.Nil(t, err, "Error should be nil")
}

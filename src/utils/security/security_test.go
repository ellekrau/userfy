package security

import (
	"github.com/ellekrau/userfy/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	config.DataSecurity.SecurityDataKey = "r%of*34_f6yif9$3"

	var err error
	data := "this data will be encrypted and decrypted"

	var encryptedData string
	if encryptedData, err = Encrypt(data); err != nil {
		assert.Fail(t, "encrypt data error: ", err.Error())
	}

	var decryptedData string
	if decryptedData, err = Decrypt(encryptedData); err != nil {
		assert.Fail(t, "decrypt data error: ", err.Error())
	}

	assert.Equal(t, data, decryptedData)
}

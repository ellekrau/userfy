package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/ellekrau/mercafacil/config"
)

var byteArray = []byte{114, 37, 111, 102, 42, 51, 52, 95, 102, 54, 121, 105, 102, 57, 114, 116}

func Encrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(config.DataSecurity.SecurityDataKey))
	if err != nil {
		return "", err
	}

	byteArrayData := []byte(data)
	encryptedData := make([]byte, len(byteArrayData))

	encrypt := cipher.NewCFBEncrypter(block, byteArray)
	encrypt.XORKeyStream(encryptedData, byteArrayData)

	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func Decrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(config.DataSecurity.SecurityDataKey))
	if err != nil {
		return "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decryptedData := make([]byte, len(decodedData))

	encrypt := cipher.NewCFBDecrypter(block, byteArray)
	encrypt.XORKeyStream(decryptedData, decodedData)

	return string(decryptedData), nil
}

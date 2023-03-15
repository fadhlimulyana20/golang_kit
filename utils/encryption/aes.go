package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type AESEcryption struct {
	Key string
}

type AESEncryptionContract interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

func NewAESEncrypt(key string) AESEncryptionContract {
	return &AESEcryption{
		Key: key,
	}
}

func (a *AESEcryption) Encrypt(plainText string) (string, error) {
	text := []byte(plainText)
	key := []byte(a.Key)

	c, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	cipherText := gcm.Seal(nonce, nonce, text, nil)
	cipherTextb64 := base64.URLEncoding.EncodeToString(cipherText)
	return cipherTextb64, nil
}

func (a *AESEcryption) Decrypt(cipherText string) (string, error) {
	text, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	key := []byte(a.Key)

	c, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		logrus.Error("invalid nonce")
		return "", err
	}

	nonce, ciphertext := text[:nonceSize], text[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return string(plaintext), nil
}

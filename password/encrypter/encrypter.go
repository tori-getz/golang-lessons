package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"password/app/env"
)

type Encrypter struct {
	key string
}

func NewEncrypter() *Encrypter {
	envKey := env.Get(env.Encryptkey)

	if envKey == "" {
		panic("ENCRYPT_KEY variable is not set!")
	}

	return &Encrypter{
		key: envKey,
	}
}

func (encrypter *Encrypter) Encrypt(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypter.key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainString, nil)
}

func (encrypter *Encrypter) Decrypt(encryptedString []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypter.key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"embed"
	"encoding/json"
)

var DEFAULT_PORT = 8000

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MustRead(fs embed.FS, name string) []byte {
	data, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return data
}

func MustMarshal(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

//go:embed aes.key
var AesKey []byte

func Encrypt(plaintext []byte) []byte {
	for len(plaintext)%aes.BlockSize != 0 {
		plaintext = append(plaintext, 0)
	}

	blockCipher, err := aes.NewCipher(AesKey)
	PanicOnErr(err)

	cbcEncrypter := cipher.NewCBCEncrypter(blockCipher, make([]byte, aes.BlockSize))
	ciphertext := make([]byte, len(plaintext))
	cbcEncrypter.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func Decrypt(ciphertext []byte) []byte {
	blockCipher, err := aes.NewCipher(AesKey)
	PanicOnErr(err)
	cbcEncrypter := cipher.NewCBCDecrypter(blockCipher, make([]byte, aes.BlockSize))
	plaintext := make([]byte, len(ciphertext))
	cbcEncrypter.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

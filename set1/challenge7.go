package set1

import (
	"crypto/aes"
	"fmt"
)

func DecryptAES(input, key []byte) string {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating aes cipher:", err)
	}
	cipherSize := cipherBlock.BlockSize()
	decrypted := make([]byte, len(input))
	for idx := 0; idx < len(input); idx += cipherSize {
		cipherBlock.Decrypt(decrypted[idx:idx+cipherSize], input[idx:idx+cipherSize])
	}
	return string(decrypted)
}

func EncryptAES(input, key []byte) string {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating aes cipher:", err)
	}
	cipherSize := cipherBlock.BlockSize()
	encrypted := make([]byte, len(input))
	for idx := 0; idx < len(input); idx += cipherSize {
		cipherBlock.Encrypt(encrypted[idx:idx+cipherSize], input[idx:idx+cipherSize])
	}
	return string(encrypted)
}

package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func Challenge7() {
	encrypted, _ := ioutil.ReadFile("7.txt")
	encrypted, _ = base64.StdEncoding.DecodeString(string(encrypted))
	encrypted = []byte(encrypted)
	key := []byte("YELLOW SUBMARINE")
	decrypted := DecryptECB(encrypted, key)
	fmt.Printf("SOLUTION 7: %v... \n\n", string(decrypted[:80]))
	if string(EncryptECB(decrypted, key)) != string(encrypted) {
		panic("Encrypt/Decrypt ECB not symmetrical")
	}
}

func DecryptECB(input, key []byte) []byte {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating aes cipher:", err)
	}
	cipherSize := cipherBlock.BlockSize()
	decrypted := make([]byte, len(input))
	for idx := 0; idx < len(input); idx += cipherSize {
		cipherBlock.Decrypt(decrypted[idx:idx+cipherSize], input[idx:idx+cipherSize])
	}
	return decrypted
}

func EncryptECB(input, key []byte) []byte {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating aes cipher:", err)
	}
	cipherSize := cipherBlock.BlockSize()
	encrypted := make([]byte, len(input))
	for idx := 0; idx < len(input); idx += cipherSize {
		cipherBlock.Encrypt(encrypted[idx:idx+cipherSize], input[idx:idx+cipherSize])
	}
	return encrypted
}

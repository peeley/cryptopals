package set2

import (
	"cryptopals/set1"
	"fmt"
	"math/rand"
	"time"
)

func RandomBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for idx, _ := range bytes {
		bytes[idx] = byte(rand.Intn(256))
	}
	return bytes
}

func EncryptWithRandomMode(input []byte) string {
	rand.Seed(time.Now().UnixNano())
	key := RandomBytes(16)
	prependBytes := RandomBytes(rand.Intn(6) + 5)
	appendBytes := RandomBytes(rand.Intn(6) + 5)
	plainText := append(prependBytes, input...)
	plainText = append(input, appendBytes...)
	padLen := 16 - (len(plainText) % 16) + len(plainText)
	plainText = []byte(PadPKCS(plainText, padLen))
	var encrypted string
	if rand.Intn(2) == 1 {
		fmt.Println("encrypting with ECB mode")
		encrypted = set1.EncryptAES(plainText, key)
	} else {
		fmt.Println("encrypting with CBC mode")
		iv := RandomBytes(16)
		encrypted = EncryptCBC(plainText, key, iv)
	}
	return encrypted
}

func DetectAESMode(input []byte) string {
	hexStringInput := set1.BytesToHexString(input)
	if set1.DetectAES(hexStringInput) != "" {
		return "ECB"
	}
	return "CBC"
}

package set2

import (
	"cryptopals/set1"
	"fmt"
	"math/rand"
	"time"
)

func Challenge11() {
	input := []byte("YELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINE")
	mode, encrypted := EncryptWithRandomMode(input)
	detectedMode := DetectAESMode(encrypted)
	fmt.Printf("\nSOLUTION 11: detected %v mode, encrypted with %v mode \n", detectedMode, mode)
	if mode != detectedMode {
		panic("Unable to correctly guess AES mode.")
	}
}

func RandomBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for idx := range bytes {
		bytes[idx] = byte(rand.Intn(256))
	}
	return bytes
}

func EncryptWithRandomMode(input []byte) (string, []byte) {
	key := RandomBytes(16)
	prependBytes := RandomBytes(rand.Intn(6) + 5)
	appendBytes := RandomBytes(rand.Intn(6) + 5)
	plainText := append(prependBytes, input...)
	plainText = append(input, appendBytes...)
	padLen := 16 - (len(plainText) % 16) + len(plainText)
	plainText = []byte(PadPKCS(plainText, padLen))
	var encrypted []byte
	var mode string
	if rand.Intn(2) == 1 {
		mode = "ECB"
		encrypted = set1.EncryptECB(plainText, key)
	} else {
		mode = "CBC"
		iv := RandomBytes(16)
		encrypted = EncryptCBC(plainText, key, iv)
	}
	return mode, encrypted
}

func DetectAESMode(input []byte) string {
	hexStringInput := set1.BytesToHexString(input)
	isECB, _ := set1.DetectECB(hexStringInput)
	if isECB {
		return "ECB"
	}
	return "CBC"
}

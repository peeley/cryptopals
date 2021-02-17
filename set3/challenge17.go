package set3

import (
	"math/rand"
	"cryptopals/set2"
	"fmt"
)

var RandomStrings = []string{
	"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
	"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
	"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
	"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
	"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
	"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
	"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
	"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
	"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
	"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
}

var Iv = []byte("YELLOW SUBMARINE")

var RandomKey = set2.RandomBytes(16)

func Challenge17(){
	fmt.Println("\nSOLUTION 17:")

	paddingAttack()

	fmt.Println()
}

func encryptRandomString() []byte {
	randomIndex := rand.Intn(len(RandomStrings))
	randomStringBytes := []byte(RandomStrings[randomIndex])

	paddedBytes := set2.PadPKCS(randomStringBytes, set2.GetPaddedLength(randomStringBytes))
	encrypted := set2.EncryptCBC(paddedBytes, RandomKey, Iv)

	return encrypted
}

func decryptAndCheckPadding(encrypted []byte) bool {
	decrypted := set2.DecryptCBC(encrypted, RandomKey, Iv)
	validated := set2.ValidatePKCS(decrypted)

	return len(validated) != 0
}

func paddingAttack() {
	encrypted := encryptRandomString()

	testBlock := encrypted[len(encrypted)-(len(RandomKey)*2):len(encrypted)-len(RandomKey)]
	intermediate := make([]byte, len(RandomKey))
	decrypted := make([]byte, len(RandomKey))

	for blockIdx := len(RandomKey)-1; blockIdx > 0; blockIdx-- {
		// TODO iterate backwards to set previously cracked pad bytes
		padValue := byte(len(RandomKey) - blockIdx)
		for byteValue := byte(0); byteValue < 255; byteValue++ {
			testBlock[blockIdx] = byteValue
			isValidPadding := decryptAndCheckPadding(encrypted)
			if isValidPadding {
				intermediate[blockIdx] = testBlock[blockIdx] ^ padValue
				decrypted[blockIdx] = testBlock[blockIdx] ^ intermediate[blockIdx]
			}
		}
	}
	fmt.Println(decrypted)
}

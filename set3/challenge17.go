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

	paddingAttackOnString()

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

func paddingAttackOnBlock(encrypted []byte, startIdx, endIdx int) []byte {

	testBlock := encrypted[startIdx:endIdx]
	intermediateBlock := make([]byte, len(RandomKey))
	decryptedBlock := make([]byte, len(RandomKey))

	originalBlock := make([]byte, len(RandomKey))
	copy(testBlock, originalBlock)

	for blockIdx := len(RandomKey)-1; blockIdx >= 0; blockIdx-- {
		padValue := byte(len(RandomKey) - blockIdx)
		for solvedIdx := len(RandomKey)-1; solvedIdx > blockIdx; solvedIdx-- {
			testBlock[solvedIdx] = intermediateBlock[solvedIdx] ^ padValue
		}
		for byteValue := byte(0); byteValue < 255; byteValue++ {
			testBlock[blockIdx] = byteValue
			isValidPadding := decryptAndCheckPadding(encrypted)
			if isValidPadding {
				intermediateBlock[blockIdx] = testBlock[blockIdx] ^ padValue
				decryptedBlock[blockIdx] = originalBlock[blockIdx] ^ intermediateBlock[blockIdx]
				fmt.Printf("cracked byte at %v: %v\n", blockIdx, decryptedBlock[blockIdx])
			}
		}
	}
	fmt.Printf("decrypted block from %v to %v: %v\n", startIdx, endIdx, string(decryptedBlock))
	return decryptedBlock
}

func paddingAttackOnString() []byte {
	encrypted := encryptRandomString()
	var decrypted []byte

	for blockIdx := len(encrypted) - len(RandomKey); blockIdx > 0; blockIdx -= len(RandomKey) {
		decryptedBlock := paddingAttackOnBlock(
			encrypted[0:blockIdx+len(RandomKey)],
			blockIdx,
			blockIdx+len(RandomKey))
		decrypted = append(decryptedBlock, decrypted...)
	}

	fmt.Println("entire string:", string(decrypted))
	return decrypted
}

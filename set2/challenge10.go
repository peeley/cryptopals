package set2

import (
	"cryptopals/set1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func Challenge10() {
	input, _ := ioutil.ReadFile("10.txt")
	input, _ = base64.StdEncoding.DecodeString(string(input))
	key := []byte("YELLOW SUBMARINE")
	initVector := make([]byte, len(key))
	for idx := range initVector {
		initVector[idx] = '0'
	}
	solution := DecryptCBC(input, key, initVector)
	fmt.Println("\nSOLUTION 10:", string(solution[:80]))
}

func DecryptCBC(input, key, iv []byte) []byte {
	chunkSize := len(key)
	lastChunk := iv
	var decrypted []byte
	for idx := 0; idx < len(input)-chunkSize; idx += chunkSize {
		thisChunk := input[idx : idx+16]
		decryptedChunk := set1.DecryptECB(thisChunk, key)
		xordChunk := set1.XORBytes(decryptedChunk, lastChunk)
		decrypted = append(decrypted, xordChunk...)
		lastChunk = thisChunk
	}
	return decrypted
}

func EncryptCBC(input, key, iv []byte) []byte {
	chunkSize := len(key)
	lastChunk := iv
	var encrypted []byte
	for idx := 0; idx < len(input)-chunkSize; idx += chunkSize {
		thisChunk := input[idx : idx+16]
		xordChunk := set1.XORBytes(thisChunk, lastChunk)
		encryptedChunk := set1.EncryptECB(xordChunk, key)
		encrypted = append(encrypted, encryptedChunk...)
		lastChunk = thisChunk
	}
	return encrypted
}

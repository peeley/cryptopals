package set2

import "cryptopals/set1"

func DecryptCBC(input []byte, key []byte, iv []byte) string {
	chunkSize := len(key)
	lastChunk := iv
	var decrypted []byte
	for idx := 0; idx < len(input)-chunkSize; idx += chunkSize {
		thisChunk := input[idx : idx+16]
		decryptedChunk := []byte(set1.DecryptAES(thisChunk, key))
		xordChunk := set1.XORBytes(decryptedChunk, lastChunk)
		decrypted = append(decrypted, xordChunk...)
		lastChunk = thisChunk
	}
	return string(decrypted)
}

func EncryptCBC(input []byte, key []byte, iv []byte) string {
	chunkSize := len(key)
	lastChunk := iv
	var encrypted []byte
	for idx := 0; idx < len(input)-chunkSize; idx += chunkSize {
		thisChunk := input[idx : idx+16]
		xordChunk := set1.XORBytes(thisChunk, lastChunk)
		encryptedChunk := []byte(set1.EncryptAES(xordChunk, key))
		encrypted = append(encrypted, encryptedChunk...)
		lastChunk = thisChunk
	}
	return string(encrypted)
}

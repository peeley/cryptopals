package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

var oracleKey = RandomBytes(16)

func Challenge12() {
	fmt.Println("\nSOLUTION 12:")
	blockSize := FindBlockSize()
	fmt.Println("  detecting a block size of ", blockSize)
	revealed := ByteAtATimeDecrypt(blockSize)
	fmt.Printf("  cracked: \n'%v'\n", string(revealed))
}

func FindBlockSize() int {
	dummyInput := []byte{}
	initialLength := len(EncryptionOracle(dummyInput))
	for {
		dummyInput = append(dummyInput, 'A')
		thisLength := len(EncryptionOracle(dummyInput))
		lengthDiff := thisLength - initialLength
		if lengthDiff >= 1 {
			return lengthDiff
		}
	}
}

func EncryptionOracle(userInput []byte) []byte {
	unknown, _ := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	combined := append(userInput, unknown...)
	return PadAndEncryptECB(combined, oracleKey)
}

func ByteAtATimeDecrypt(blockSize int) []byte {
	oracleLength := len(EncryptionOracle([]byte{}))
	var cracked []byte

	for length := 1; length < oracleLength; length++ {
		prefixLength := (blockSize - (len(cracked) % blockSize)) - 1
		prefix := make([]byte, prefixLength)

		for prefixIdx := range prefix {
			prefix[prefixIdx] = 'A'
		}

		realOutput := EncryptionOracle(prefix)
		testBlock := append(prefix, cracked...)

		for possibleByte := byte(0); possibleByte < 255; possibleByte++ {
			block := append(testBlock, possibleByte)
			possibleOut := EncryptionOracle(block)

			if bytes.Equal(possibleOut[:len(block)],
							realOutput[:len(block)]) {
				cracked = append(cracked, possibleByte)
				break
			}
		}
	}
	return cracked
}

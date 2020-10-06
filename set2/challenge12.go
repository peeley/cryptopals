package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func Challenge12() {
	fmt.Println("SOLUTION 12:")
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
	rand.Seed(69) // create consistent but unknown key
	key := RandomBytes(16)
	rand.Seed(time.Now().UnixNano()) // reset for future keys to be random
	unknown, _ := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	combined := append(userInput, unknown...)
	return PadAndEncryptECB(combined, key)
}

func ByteAtATimeDecrypt(blockSize int) []byte {
	oracleLength := len(EncryptionOracle([]byte{}))
	var cracked []byte
	for length := 0; length < oracleLength; length++ {
		prefixLength := (blockSize - (len(cracked) % blockSize)) - 1
		prefix := make([]byte, prefixLength)
		for prefixIdx := range prefix {
			prefix[prefixIdx] = 'A'
		}
		shortOutput := EncryptionOracle(prefix)
		testBlock := append(prefix, cracked...)
		for possibleByte := byte(0); possibleByte < 255; possibleByte++ {
			block := append(testBlock, possibleByte)
			oracleOut := EncryptionOracle(block)
			if bytes.Equal(oracleOut[:prefixLength+len(cracked)+1], shortOutput[:prefixLength+len(cracked)+1]) {
				cracked = append(cracked, possibleByte)
				break
			}
		}
	}
	return cracked
}

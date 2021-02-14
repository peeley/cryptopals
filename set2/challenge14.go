package set2

import (
	"fmt"
	"math/rand"
	"bytes"
)

var randomPrefix []byte = RandomBytes(rand.Intn(48))

func Challenge14(){
	fmt.Println("\nSOLUTION 14:")

	fmt.Printf("random prefix: %v, len: %v\n",
		string(randomPrefix),
		len(randomPrefix))
	blockSize := FindBlockSize2()
	prefixLength := FindRandomPrefixLength(blockSize)

	fmt.Println("blockSize:", blockSize)
	fmt.Println("cracked prefix length: ", prefixLength)

	cracked := ByteAtATimeDecrypt2(blockSize, prefixLength)
	fmt.Println("cracked", string(cracked))

	fmt.Println()
}

func EncryptionOracle2(input []byte) []byte {
	userInput := append(randomPrefix, input...)
	return EncryptionOracle(userInput)
}

func FindBlockSize2() int {
	dummyInput := []byte{}
	initialLength := len(EncryptionOracle2(dummyInput))
	for {
		dummyInput = append(dummyInput, 'A')
		thisLength := len(EncryptionOracle2(dummyInput))
		lengthDiff := thisLength - initialLength
		if lengthDiff >= 1 {
			return lengthDiff
		}
	}
}

func FindRandomPrefixLength(blockSize int) int {
	dummyInput := make([]byte, blockSize * 3, 'A')
	oracled := EncryptionOracle2(dummyInput)
	for i := 0; i < len(oracled); i++ {
		firstBlock := oracled[i:i+blockSize]
		secondBlock := oracled[i+blockSize:i+(blockSize*2)]
		if bytes.Equal(firstBlock, secondBlock) {
			return i
		}
	}
	return -1
}

func ByteAtATimeDecrypt2(blockSize, prefixLength int) []byte {
	oracleLength := len(input)
	var cracked []byte
	for length := 0; length < oracleLength; length++ {
		prefixLength := (blockSize - (len(cracked) % blockSize)) - 1
		prefix := make([]byte, prefixLength)
		for prefixIdx := range prefix {
			prefix[prefixIdx] = 'A'
		}
		shortOutput := EncryptionOracle2(prefix)
		testBlock := append(prefix, cracked...)
		for possibleByte := byte(0); possibleByte < 255; possibleByte++ {
			block := append(testBlock, possibleByte)
			oracleOut := EncryptionOracle2(block)
			if bytes.Equal(oracleOut[:prefixLength+len(cracked)+1], shortOutput[:prefixLength+len(cracked)+1]) {
				cracked = append(cracked, possibleByte)
				break
			}
		}
	}
	return cracked
}

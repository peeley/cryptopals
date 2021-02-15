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
		randomPrefix,
		len(randomPrefix))
	blockSize := FindBlockSize2()
	paddingStart, endBlockSize := FindRandomPrefixLength(blockSize)
	fmt.Printf("padding starts at %v, third block of %v long pulls in one secret char\n", paddingStart, endBlockSize)

	cracked := ByteAtATimeDecrypt2(blockSize, paddingStart, endBlockSize)
	fmt.Println("cracked: ", string(cracked))

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

func FindRandomPrefixLength(blockSize int) (int, int) {
	dummyInput := make([]byte, (blockSize * 3))
	oracled := EncryptionOracle2(dummyInput)
	var paddingStart int
	var paddingEnd int
	for i := 0; i < len(oracled) - (blockSize * 2); i++ {
		firstBlock := oracled[i:i+blockSize]
		secondBlock := oracled[i+blockSize:i+(blockSize*2)]
		if bytes.Equal(firstBlock, secondBlock) {
			paddingStart = i
			paddingEnd = i+(blockSize*2)
		}
	}

	if paddingStart == 0 && paddingEnd == 0 {
		return -1, -1
	}

	for i := blockSize; i > 0; i-- {
		testInput := make([]byte, (blockSize * 2) + i)
		testOracle := EncryptionOracle2(testInput)
		if !bytes.Equal(testOracle[paddingStart:paddingEnd], oracled[paddingStart:paddingEnd]) {
			fmt.Println(testOracle[paddingStart:paddingStart+blockSize],
				testOracle[paddingStart+blockSize:paddingStart+(blockSize*2)])
			fmt.Println(oracled[paddingStart:paddingStart+blockSize],
				oracled[paddingStart+blockSize:paddingStart+(blockSize*2)])
			return paddingStart, i
		}
	}
	return -1, -1
}

func ByteAtATimeDecrypt2(blockSize, paddingStart, endBlockSize int) []byte {
	var cracked []byte
	oracleLength := len(EncryptionOracle2([]byte{}))
	for length := 0; length < oracleLength; length++ {

		oneShortPrefix := make([]byte, blockSize*10 - len(cracked) + endBlockSize)

		shortOutput := EncryptionOracle2(oneShortPrefix)
		testBlock := append(oneShortPrefix, cracked...)
		fmt.Println(testBlock)

		for possibleByte := byte(0); possibleByte < 255; possibleByte++ {
			block := append(testBlock, possibleByte)
			possibleOutput := EncryptionOracle2(block)

			if bytes.Equal(possibleOutput[:len(block)],
							shortOutput[:len(block)]) {

				cracked = append(cracked, possibleByte)
				break
			}
		}
	}
	return cracked
}

package set2

import (
	"fmt"
	"math/rand"
)

var randomPrefix []byte

func Challenge14(){
	randomPrefix = RandomBytes(rand.Intn(16))
	fmt.Println("\nSOLUTION 14:")

	fmt.Println(FindBlockAndPrefixSize())

	fmt.Println()
}

func EncryptionOracle2(input []byte) []byte {
	userInput := append(input, randomPrefix...)
	return EncryptionOracle(userInput)
}

func FindBlockAndPrefixSize() (int, int) {
	dummyInput := []byte{}
	initialLength := len(EncryptionOracle2(dummyInput))
	for {
		dummyInput = append(dummyInput, 'A')
		thisLength := len(EncryptionOracle2(dummyInput))
		lengthDiff := thisLength - initialLength
		if lengthDiff >= 1 {
			return lengthDiff, lengthDiff - len(dummyInput)
		}
	}
}

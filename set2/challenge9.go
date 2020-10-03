package set2

import (
	"cryptopals/set1"
	"fmt"
)

func Challenge9() {
	input := []byte("YELLOW SUBMARINE")
	length := 20
	fmt.Println("SOLUTION 9:", string(PadPKCS(input, length)))
}

func PadPKCS(input []byte, length int) []byte {
	var diff byte = byte(length - len(input))
	var output []byte
	output = append(output, input...)
	for i := byte(0); i < diff; i++ {
		output = append(output, diff)
	}
	return output
}

func GetPaddedLength(buffer []byte) int {
	return 16 - (len(buffer) % 16) + len(buffer)
}

func PadAndEncryptECB(input, key []byte) []byte {
	padded := PadPKCS(input, GetPaddedLength(input))
	return set1.EncryptECB(padded, key)
}

package set2

import "fmt"

func Challenge9() {
	input := []byte("YELLOW SUBMARINE")
	length := 20
	fmt.Println("SOLUTION 9:", string(PadPKCS(input, length)))
	fmt.Println()
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

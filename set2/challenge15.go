package set2

import (
	"fmt"
)

func Challenge15(){
	fmt.Println("\nSOLUTION 15:")

	padded := PadPKCS([]byte("ICE ICE BABY"), 16)
	validated := string(validatePKCS(padded))
	invalid := []byte("ICE ICE BABY\x01\x02\x03\x04")
	invalidated := string(validatePKCS(invalid))
	fmt.Println(validated)
	fmt.Println(invalidated)
	fmt.Println()
}

func validatePKCS(string []byte) []byte {
	if len(string) == 0 {
		return []byte{}
	}

	padValue := string[len(string)-1]
	numPadValues := byte(0)

	var index int
	for index = len(string) - 1; index >= 0; index-- {
		if string[index] != padValue {
			break
		}
		numPadValues += 1
	}

	if padValue == numPadValues {
		return string[:index+1]
	}

	return []byte{}
}

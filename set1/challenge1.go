package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Challenge1() {
	decodedBase64 := HexStringToBase64(
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println("SOLUTION 1:", decodedBase64)
}

func HexStringToBase64(input string) string {
	rawBytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(rawBytes, []byte(input))
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(rawBytes)
}

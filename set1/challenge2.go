package set1

import (
	"encoding/hex"
	"fmt"
)

func Challenge2() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	fmt.Println("SOLUTION 2:", XORHexStrings(input1, input2), "\n")
}

func XORHexStrings(input1 string, input2 string) string {
	rawBytes1 := HexStringToBytes(input1)
	rawBytes2 := HexStringToBytes(input2)
	return BytesToHexString(XORBytes(rawBytes1, rawBytes2))
}

func XORBytes(input1, input2 []byte) []byte {
	xorBytes := make([]byte, len(input2))
	for idx, _ := range input1 {
		xorBytes[idx] = input1[idx] ^ input2[idx]
	}
	return xorBytes
}

func HexStringToBytes(input string) []byte {
	rawBytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(rawBytes, []byte(input))
	if err != nil {
		fmt.Println("Error decoding hex string to bytes:", err)
	}
	return rawBytes
}

func BytesToHexString(input []byte) string {
	output := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(output, input)
	return string(output)
}

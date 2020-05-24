package set1

import (
	"encoding/hex"
	"fmt"
)

func Challenge2(input1 string, input2 string) string {
	rawBytes1 := HexStringToBytes(input1)
	rawBytes2 := HexStringToBytes(input2)
	xorBytes := make([]byte, len(rawBytes2))
	for idx, _ := range rawBytes1 {
		xorBytes[idx] = rawBytes1[idx] ^ rawBytes2[idx]
	}
	return BytesToHexString(xorBytes)
}

func HexStringToBytes(input string) []byte {
	rawBytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(rawBytes, []byte(input))
	if err != nil {
		fmt.Println(err)
	}
	return rawBytes
}

func BytesToHexString(input []byte) string {
	output := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(output, input)
	return string(output)
}

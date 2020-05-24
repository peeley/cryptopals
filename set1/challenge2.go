package set1

import (
	"encoding/hex"
	"fmt"
)

func Challenge2(input1 string, input2 string) string {
	rawBytes1 := hexStringToBytes(input1)
	rawBytes2 := hexStringToBytes(input2)
	xorBytes := make([]byte, len(rawBytes2))
	for idx, _ := range rawBytes1 {
		xorBytes[idx] = rawBytes1[idx] ^ rawBytes2[idx]
	}
	output := make([]byte, hex.EncodedLen(len(xorBytes)))
	hex.Encode(output, xorBytes)
	return string(output)
}

func hexStringToBytes(input string) []byte {
	rawBytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(rawBytes, []byte(input))
	if err != nil {
		fmt.Println(err)
	}
	return rawBytes
}

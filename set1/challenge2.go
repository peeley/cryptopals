package set1

import (
	"encoding/hex"
	"fmt"
)

func Challenge2(input1 string, input2 string) string {
	rawBytes1 := make([]byte, hex.DecodedLen(len(input1)))
	_, err := hex.Decode(rawBytes1, []byte(input1))
	if err != nil {
		fmt.Println(err)
	}
	rawBytes2 := make([]byte, hex.DecodedLen(len(input2)))
	_, err = hex.Decode(rawBytes2, []byte(input2))
	if err != nil {
		fmt.Println(err)
	}
	xorBytes := make([]byte, len(rawBytes2))
	for idx, _ := range rawBytes1 {
		xorBytes[idx] = rawBytes1[idx] ^ rawBytes2[idx]
	}
	output := make([]byte, hex.EncodedLen(len(xorBytes)))
	hex.Encode(output, xorBytes)
	return string(output)
}

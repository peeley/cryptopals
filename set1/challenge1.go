package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexStringToBase64(input string) string {
	rawBytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(rawBytes, []byte(input))
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(rawBytes)
}

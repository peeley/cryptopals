package set2

import (
	"fmt"
	"strings"
	"regexp"
)

var Key = []byte("YELLOW SUBMARINE")
var Iv =  []byte("0000000000000000")

const (
	Prepend = "comment1=cooking%20MCs;userdata="
	Append = ";comment2=%20like%20a%20pound%20of%20bacon"
)

func Challenge16(){
	fmt.Println("\nSOLUTION 16:")

	input := []byte("asdf;admin=true")
	encrypted := padAndEncrypt(input)
	isAdmin := decryptAndSearchForAdmin(encrypted)
	fmt.Println("user is admin: ", isAdmin)
	fmt.Println()
}

func padAndEncrypt(input []byte) []byte {
	stringInput := string(input)
	quotedInput := strings.ReplaceAll(stringInput, ";", "\";\"")
	quotedInput = strings.ReplaceAll(quotedInput, "=", "\"=\"")
	quotedInputBytes := []byte(quotedInput)

	prepended := append([]byte(Prepend), quotedInputBytes...)
	appended := append(prepended, []byte(Append)...)

	preparedInput := string(appended)

	preparedInputBytes := []byte(preparedInput)
	paddedLength := GetPaddedLength(preparedInputBytes)
	padded := PadPKCS(preparedInputBytes, paddedLength)
	return EncryptCBC(padded, Key, []byte(Iv))
}

func decryptAndSearchForAdmin(encrypted []byte) bool {
	decrypted := DecryptCBC(encrypted, Key, []byte(Iv))
	matches, err := regexp.Match(";admin=true;", decrypted)

	if err != nil {
		fmt.Println("Could not match decrypted with regex")
		return false
	}

	return matches
}

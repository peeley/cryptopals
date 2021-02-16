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

	fmt.Println("bitflipped CBC encryption input:")
	bitflipAttack()

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
	return EncryptCBC(padded, Key, Iv)
}

func decryptAndSearchForAdmin(encrypted []byte) bool {
	decrypted := DecryptCBC(encrypted, Key, Iv)
	matches, err := regexp.Match(";admin=true;", decrypted)

	if err != nil {
		fmt.Println("Could not match decrypted with regex")
		return false
	}

	return matches
}

func bitflipAttack() {
	input := make([]byte, len(Key) * 2)
	for i := range(input) {
		input[i] = 0
	}

	startOfInjectedData := len(Prepend)
	startOfCorruptedData := len(Prepend) + len(Key)
	targetData := []byte(";admin=true;")
	encrypted := padAndEncrypt(input)

	for idx := 0; idx < len(targetData); idx++ {
		for byteOffset := 0; byteOffset < 255; byteOffset ++ {
			encrypted[idx + startOfInjectedData] = byte(byteOffset)
			decrypted := DecryptCBC(encrypted, Key, Iv)
			if decrypted[startOfCorruptedData + idx] == targetData[idx] {
				break
			}
		}
	}
	isAdmin := decryptAndSearchForAdmin(encrypted)
	if isAdmin {
		fmt.Println(string(DecryptCBC(encrypted, Key, Iv)))
	}
}

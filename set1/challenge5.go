package set1

import "fmt"

func Challenge5() {
	poetry := []byte("Burning 'em, if you ain't quick and nimble")
	key := []byte("ICE")
	encrypted := EncryptRotatingXOR(poetry, key)
	fmt.Println("\nSOLUTION 5:", BytesToHexString(encrypted))
}

func EncryptRotatingXOR(input, key []byte) []byte {
	inputBytes := []byte(input)
	keyBytes := []byte(key)
	encrypted := make([]byte, len(inputBytes))
	for idx := 0; idx < len(inputBytes); idx++ {
		encrypted[idx] = inputBytes[idx] ^ keyBytes[idx%len(keyBytes)]
	}
	return encrypted
}

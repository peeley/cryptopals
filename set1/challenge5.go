package set1

func EncryptRotatingXOR(input string, key string) string {
	inputBytes := []byte(input)
	keyBytes := []byte(key)
	encrypted := make([]byte, len(inputBytes))
	for idx := 0; idx < len(inputBytes); idx++ {
		encrypted[idx] = inputBytes[idx] ^ keyBytes[idx%len(keyBytes)]
	}
	return string(BytesToHexString(encrypted))
}

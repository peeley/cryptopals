package set2

func PadPKCS(input []byte, length int) string {
	var diff byte = byte(length - len(input))
	var output []byte
	output = append(output, input...)
	for i := byte(0); i < diff; i++ {
		output = append(output, diff)
	}
	return string(output)
}

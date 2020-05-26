package set1

import (
	"strings"
)

func DetectAES(input string) string {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		lineBytes := HexStringToBytes(line)
		chunk := lineBytes[:16]
		for lineIdx := 1; lineIdx < (len(lineBytes) - 16); lineIdx++ {
			if HammingDistance(lineBytes[lineIdx:lineIdx+16], chunk) == 0 {
				return string(HexStringToBytes(line))
			}
		}
	}
	return "no aes detected"
}

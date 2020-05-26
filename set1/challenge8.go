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
		for xIdx := 16; xIdx < (len(lineBytes) - 16); xIdx += 16 {
			thisChunk := lineBytes[xIdx : xIdx+16]
			for yIdx := xIdx + 16; yIdx < (len(lineBytes) - 16); yIdx += 16 {
				comparedChunk := lineBytes[yIdx : yIdx+16]
				if HammingDistance(thisChunk, comparedChunk) == 0 {
					return line
				}
			}
		}
	}
	return "no aes detected"
}

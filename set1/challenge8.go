package set1

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Challenge8() {
	contents, _ := ioutil.ReadFile("8.txt")
	_, line := DetectECB(string(contents))
	fmt.Println("SOLUTION 8:\n", line)
}

func DetectECB(input string) (bool, string) {
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
					return true, line
				}
			}
		}
	}
	return false, ""
}

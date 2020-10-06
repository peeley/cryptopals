package set1

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Challenge4() {
	contents, err := ioutil.ReadFile("4.txt")
	if err != nil {
		fmt.Println("Cannot open file", "4.txt")
	}
	fmt.Println("\nSOLUTION 4:", FindDecrypted(string(contents)))
}

func FindDecrypted(lines string) string {
	var bestCandidate string
	var bestScore float32 = -1.0
	for _, line := range strings.Split(lines, "\n") {
		decrypted, _ := DecryptXOR(HexStringToBytes(line))
		score := FrequencyScore(decrypted)
		if score > bestScore {
			bestScore = score
			bestCandidate = string(decrypted)
		}
	}
	return bestCandidate
}

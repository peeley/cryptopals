package set1

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FindDecrypted(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot open file", filename)
	}
	lines := string(contents)
	var bestCandidate string
	var bestScore float32 = -1.0
	for _, line := range strings.Split(lines, "\n") {
		decrypted, _ := DecryptXOR(line)
		score := FrequencyScore([]byte(decrypted))
		if score > bestScore {
			bestScore = score
			bestCandidate = decrypted
		}
	}
	return bestCandidate
}

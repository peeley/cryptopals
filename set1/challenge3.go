package set1

var english_frequencies = map[byte]float32{
	'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610,
	'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490, 'm': 0.0202124, 'n': 0.0564513,
	'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357, 'u': 0.0225134,
	'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692, 'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182}

func Challenge3(input string) string {
	rawBytes := HexStringToBytes(input)
	var cipherChar byte
	var highestScore float32 = -1.0
	var bestCandidate []byte
	for cipherChar = 0; cipherChar < 255; cipherChar++ {
		decoded := make([]byte, len(rawBytes))
		for idx, char := range rawBytes {
			decoded[idx] = char ^ cipherChar
		}
		cipherScore := frequencyScore(decoded)
		if cipherScore > highestScore {
			highestScore = cipherScore
			bestCandidate = decoded
		}
	}
	return string(bestCandidate)
}

func frequencyScore(candidate []byte) float32 {
	var score float32 = 0
	for _, char := range candidate {
		score += english_frequencies[char]
	}
	return score
}

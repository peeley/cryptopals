package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
)

const (
	keySizeLowerBound = 2
	keySizeUpperBound = 40
)

func DecryptRotatingXOR(filename string) string {
	base64Bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot open file", filename)
	}
	inputBytes := make([]byte, base64.StdEncoding.DecodedLen(len(base64Bytes)))
	_, err = base64.StdEncoding.Decode(inputBytes, base64Bytes)
	if err != nil {
		fmt.Println("Error decoding base 64", err)
	}
	keySize := findKeySize(inputBytes)
	blocks := splitIntoBlocks(keySize, inputBytes)
	var cipher []byte
	for _, block := range blocks {
		_, cipherChar := DecryptXOR(BytesToHexString(block))
		cipher = append(cipher, cipherChar)
	}
	fmt.Println("found key:", cipher)
	return EncryptRotatingXOR(string(inputBytes), string(cipher))
}

func HammingDistance(input1 []byte, input2 []byte) float64 {
	distance := 0.0
	for byteIdx, _ := range input1 {
		byte1 := input1[byteIdx]
		byte2 := input2[byteIdx]
		for bitIdx := 0; bitIdx < 8; bitIdx++ {
			mask := byte(1 << uint(bitIdx))
			if (byte1 & mask) != (byte2 & mask) {
				distance += 1
			}
		}
	}
	return distance
}

func findKeySize(inputBytes []byte) int {
	var distance float64
	var smallestDistance float64 = math.Inf(1)
	var bestKeySize int
	for keySize := keySizeLowerBound; keySize < keySizeUpperBound; keySize++ {
		firstChunk := inputBytes[:keySize]
		secondChunk := inputBytes[keySize : keySize*2]
		distance = HammingDistance(firstChunk, secondChunk) / float64(keySize)
		if distance < smallestDistance {
			smallestDistance = distance
			bestKeySize = keySize
		}
	}
	return bestKeySize
}

func splitIntoBlocks(keySize int, bytes []byte) [][]byte {
	blocks := make([][]byte, keySize)
	for idx, byte := range bytes {
		blocks[idx%keySize] = append(blocks[idx%keySize], byte)
	}
	return blocks
}

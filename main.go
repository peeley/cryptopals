package main

import (
	"cryptopals/set1"
	"fmt"
)

func main() {
	base64 := set1.HexStringToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println("solution to challenge 1:", base64, "\n")

	chal2input1 := "1c0111001f010100061a024b53535009181c"
	chal2input2 := "686974207468652062756c6c277320657965"
	fmt.Println("solution to challenge 2:", set1.XORStrings(chal2input1, chal2input2), "\n")

	chal3sol, _ := set1.DecryptXOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println("solution to challenge 3:", chal3sol, "\n")

	fmt.Println("solution to challenge 4:", set1.FindDecrypted("4.txt"))

	poetry := "Burning 'em, if you ain't quick and nimble"
	fmt.Println("solution to challenge 5:", set1.EncryptRotatingXOR(poetry, "ICE"), "\n")

	fmt.Println("solution to challenge 6:", set1.DecryptRotatingXOR("6.txt"))
}

package main

import (
	"cryptopals/set1"
	"fmt"
)

func main() {
	base64 := set1.Challenge1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println("solution to challenge 1: ", base64, "\n")

	chal2input1 := "1c0111001f010100061a024b53535009181c"
	chal2input2 := "686974207468652062756c6c277320657965"
	fmt.Println("solution to challenge 2:", set1.Challenge2(chal2input1, chal2input2))
}

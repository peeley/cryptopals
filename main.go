package main

import (
	"cryptopals/set1"
	"fmt"
)

func main() {
	base64 := set1.Challenge1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(base64)
}

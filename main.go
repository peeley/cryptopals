package main

import (
	"cryptopals/set1"
	"cryptopals/set2"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	set1.Challenge1()
	set1.Challenge2()
	set1.Challenge3()
	set1.Challenge4()
	set1.Challenge5()
	set1.Challenge6()
	set1.Challenge7()
	set1.Challenge8()

	fmt.Println("\nSET 2\n-----")
	set2.Challenge9()
	set2.Challenge10()
	set2.Challenge11()
	set2.Challenge12()
	set2.Challenge13()
	set2.Challenge14()
	set2.Challenge15()
}

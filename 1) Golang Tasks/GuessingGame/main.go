package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guessNb int
	var randomNb int = 0

	fmt.Println("Please enter a guess number : ")
	fmt.Scan(&guessNb)

	for guessNb != randomNb {
		rand.Seed(time.Now().UnixNano())
		randomNb = rand.Intn(100) + 1
		if randomNb == guessNb {
			fmt.Println("Guess is successful")
		} else {
			fmt.Println("Your are wrong. Random value is ", randomNb)
		}
	}
}

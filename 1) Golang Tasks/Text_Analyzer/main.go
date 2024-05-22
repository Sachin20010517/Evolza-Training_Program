package main

import (
	"fmt"
)

func main() {
	var text string

	fmt.Print("Please Enter your text: ")
	fmt.Scan(&text)

	var nbWords int = len(text)
	var characters int
	var uniqeLetters int

	for i := 0; i < len(text); i++ {
		characters++
	}

	for i := 0; i < len(text); i++ {
		uniqeLetters++
	}

	fmt.Printf("\nNumber of counts = %v", nbWords)

}

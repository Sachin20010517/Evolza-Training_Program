package main

import (
	"fmt"
)

func main() {
	var width int
	var length int

	fmt.Print("Please enter the length  of the rectangle: ")
	fmt.Scan(&length)
	fmt.Print("\nPlease enter the width  of the rectangle: ")
	fmt.Scan(&width)
	fmt.Print("\n\nThe Area of the rectangle is ", length*width)
}

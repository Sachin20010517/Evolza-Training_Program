package main

import (
	"fmt"
)

func main() {
	var number1 int
	var number2 int
	var symbol string

	fmt.Print("Please enter num1 : ")
	fmt.Scan(&number1)
	fmt.Print("\nPlease enter num2 : ")
	fmt.Scan(&number2)
	fmt.Print("\nPlease enter the symbol : ")
	fmt.Scan(&symbol)

	if symbol == "+" {
		fmt.Println("The total is ", number1+number2)
	} else if symbol == "-" {
		fmt.Println("The value is ", number1-number2)
	} else if symbol == "*" {
		fmt.Println("The value is ", number1*number2)
	} else {
		fmt.Println("The value is ", number1/number2)
	}
}

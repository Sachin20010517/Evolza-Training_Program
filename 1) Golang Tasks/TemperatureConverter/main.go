package main

import (
	"fmt"
)

func main() {
	var celcius_temp int
	var fahrenheit_temp float32

	fmt.Print("\nPlease enter the temperature in Celcius : ")
	fmt.Scan(&celcius_temp)
	fahrenheit_temp = float32(celcius_temp)*1.8 + 32
	fmt.Printf("Temperature is %vf", fahrenheit_temp)
}

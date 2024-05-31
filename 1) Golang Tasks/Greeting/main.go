package main

import (
	"fmt"
)

func main() {
	var userName string
	fmt.Print("Please Enter Your Name : ")
	fmt.Scan(&userName)
	fmt.Printf("\n\n Welcome You %v to the Go Global Comunity", userName)
}

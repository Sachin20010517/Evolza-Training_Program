package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Create a new reader for standard input
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user to enter a password
		fmt.Print("Enter a password: ")

		// Read the input from the user
		password, _ := reader.ReadString('\n')

		// Remove the newline character from the input
		password = password[:len(password)-1]

		// Check if the password meets the criteria
		if isValidPassword(password) {
			fmt.Println("Password is valid.")
			break
		} else {
			fmt.Println("Password is invalid. It must be at least 8 characters long, contain both uppercase and lowercase letters, and include at least one number.")
		}
	}
}

// Function to check if the password is valid based on specified criteria
func isValidPassword(password string) bool {
	// Criteria for a valid password
	var hasUpper, hasLower, hasNumber bool
	var length int

	// Iterate over each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
		length++
	}

	// Password must be at least 8 characters long and meet all other criteria
	if length >= 8 && hasUpper && hasLower && hasNumber {
		return true
	}

	return false
}

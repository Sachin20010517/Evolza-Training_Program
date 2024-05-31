package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// Variables to store line count and word count
	lineCount := 0
	wordCount := 0

	// Process each line of the input file
	for scanner.Scan() {
		// Increment line count
		lineCount++

		// Split the line into words
		words := strings.Fields(scanner.Text())

		// Increment word count by the number of words in the line
		wordCount += len(words)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Open the output file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the results to the output file
	_, err = fmt.Fprintf(outputFile, "   Hi There! I'm Sachin and this is the result of input.txt file\n\nLine Count: %d\nWord Count: %d", lineCount, wordCount)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("File processing complete.")
}

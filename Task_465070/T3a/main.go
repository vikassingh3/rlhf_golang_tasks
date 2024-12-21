package main

import (
	"fmt"
	"strconv"
)

func readIntFromFile(filename string) (int, error) {
	// Simulate opening a file and reading a line
	input := "123"

	// Simulate parsing the input to an integer
	number, err := strconv.Atoi(input)
	return number, err
}

func main() {
	filename := "example.txt"

	// Call the function with multiple return values
	number, err := readIntFromFile(filename)

	// Check the error
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// If no error, use the result
	fmt.Println("Read number:", number)
}
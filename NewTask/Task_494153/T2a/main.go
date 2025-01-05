package main

import (
	"errors"
	"fmt"
)

// Function that can return multiple values, including an error
func getData() (string, int, error) {
	if randomError() {
		return "", 0, errors.New("random error occurred")
	}
	return "Hello", 42, nil
}

// Simulate a random error
func randomError() bool {
	// Replace with actual random logic
	return true // For demonstration purposes, let's always return an error
}

func main() {
	// Call the function and retrieve values
	str, num, err := getData()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("String: %s, Number: %d\n", str, num)
}

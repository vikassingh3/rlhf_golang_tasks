package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	// Validate input: Check if the name contains only alphanumeric characters and spaces
	validNamePattern := regexp.MustCompile("^[A-Za-z0-9 ]+$")
	if validNamePattern.MatchString(name) {
		// Print the user's name safely using fmt.Printf
		fmt.Printf("Hello, %s! It's nice to meet you.\n", name)
	} else {
		fmt.Println("Invalid input. Please enter a valid name (alphanumeric characters and spaces only).")
	}
}
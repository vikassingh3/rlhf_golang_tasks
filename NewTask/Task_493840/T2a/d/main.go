package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidAge(age string) bool {
	num, err := strconv.Atoi(age)
	return err == nil && num >= 0
}

func main() {
	// User input
	var age string
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	age = strings.TrimSpace(age)

	// Validation
	if !isValidAge(age) {
		fmt.Println("Invalid age format. Please enter a non-negative integer.")
		return
	}

	// Process valid age
	num, _ := strconv.Atoi(age)
	fmt.Printf("Your age is: %d\n", num)
}
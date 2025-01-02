package main

import (
	"fmt"
)

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Iterate through the slice using a range loop
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
}
package main

import (
	"fmt"
)

// Function that returns multiple values
func sumAndDifference(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	// Call the function that returns multiple values
	result1, result2 := sumAndDifference(10, 5)

	// Using type assertion to extract values from an interface{}
	var iface interface{} = result1
	sum, ok := iface.(int)

	if ok {
		fmt.Printf("Sum: %d\n", sum)
	} else {
		fmt.Println("Type assertion failed for sum")
	}

	iface = result2
	difference, ok := iface.(int)

	if ok {
		fmt.Printf("Difference: %d\n", difference)
	} else {
		fmt.Println("Type assertion failed for difference")
	}
}
package main

import (
	"fmt"
)

// Function returning multiple value types, including an error
func getData() (interface{}, interface{}, error) {
	return "Hello", 42, nil
}

func main() {
	// Call the function and retrieve values
	result1, result2, err := getData()

	// Use type assertion to extract specific types
	str, ok1 := result1.(string)
	num, ok2 := result2.(int)

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if ok1 && ok2 {
		fmt.Printf("String: %s, Number: %d\n", str, num)
	} else {
		fmt.Println("Type assertion failed")
	}
}

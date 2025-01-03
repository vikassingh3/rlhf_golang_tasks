package main

import (
	"fmt"
)

// CreateAndReturnSlice creates a slice of integers from 1 to 5 and returns it
func CreateAndReturnSlice() []int {
	// Initialize a slice of integers with values from 1 to 5
	slice := []int{1, 2, 3, 4, 5}
	
	// Return the slice
	return slice
}

func main() {
	// Call the function to get the slice
	numbers := CreateAndReturnSlice()
	
	// Print the slice
	fmt.Println("The slice returned is:", numbers)
}
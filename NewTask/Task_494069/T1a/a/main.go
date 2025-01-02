package main

import (
	"fmt"
)

func addAndEdit(originalSlice []int) []int {
	// This function is intended to add a new element to the slice and edit one existing element.
	updatedSlice := append(originalSlice, 5) // Append a new element
	
	// EDITING STRATEGY ISSUE: Editing an element by index
	updatedSlice[0] = 100 // Change the first element to 100
	return updatedSlice
}

func main() {
	numbers := []int{1, 2, 3, 4} // Original slice
	
	// We are going to add an element to the original slice and change the first one
	newNumbers := addAndEdit(numbers)
	
	// PRINT SLICES
	fmt.Println("Original Slice:", numbers) // Expectation: {1, 2, 3, 4} but change shows {100, 2, 3, 4}
	fmt.Println("New Slice:", newNumbers)   // Expectation: {100, 2, 3, 4, 5}
}




package main

import "fmt"

func addAndEdit(originalSlice []int) []int {
	// Copy the original slice to ensure a new underlying array
	updatedSlice := make([]int, len(originalSlice), len(originalSlice)+1)
	copy(updatedSlice, originalSlice)

	// Append a new element to the slice
	updatedSlice = append(updatedSlice, 5)

	// Edit an element by index
	updatedSlice[0] = 100
	return updatedSlice
}

func main() {
	numbers := []int{1, 2, 3, 4} // Original slice

	// We are going to add an element to the original slice and change the first one
	newNumbers := addAndEdit(numbers)

	// PRINT SLICES
	fmt.Println("Original Slice:", numbers) // Now this remains {1, 2, 3, 4}
	fmt.Println("New Slice:", newNumbers)   // This shows the expected output {100, 2, 3, 4, 5}
}
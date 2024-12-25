package main

import (
	"fmt"
)

func main() {
	var numbers []int = []int{1, 2, 3, 4, 5}

	// Accessing an element within bounds
	if index := 3; index >= 0 && index < len(numbers) {
		fmt.Println("Accessing element at index", index, ":", numbers[index])
	} else {
		fmt.Println("Index out of bounds")
	}

	// Accessing an element out of bounds
	if index := 5; index >= 0 && index < len(numbers) {
		fmt.Println("Accessing element at index", index, ":", numbers[index])
	} else {
		fmt.Println("Index out of bounds")
	}
}


// func main() {
// 	var numbers []int

// 	// Attempting to access a nil slice
// 	if numbers != nil {
// 		fmt.Println("Accessing element at index 0:", numbers[0])
// 	} else {
// 		fmt.Println("Slice is nil")
// 	}

// 	// Assigning values to the slice
// 	numbers = []int{1, 2, 3, 4, 5}

// 	// Accessing an element of a non-nil slice
// 	if numbers != nil {
// 		fmt.Println("Accessing element at index 0:", numbers[0])
// 	} else {
// 		fmt.Println("Slice is nil")
// 	}
// }
package main

import (
	"fmt"
)

// func main() {
// 	data := []int{1, 2, 3, 4, 5}

// 	// Example 1: Using range for readability
// 	sum := 0
// 	for i, v := range data {
// 		sum += v
// 		fmt.Printf("Index: %d, Value: %d\n", i, v)
// 	}
// 	fmt.Println("Sum:", sum)
// }

// package main

// import (
// 	"fmt"
// )

func main() {
	data := []int{1, 2, 3, 4, 5}
	
	// Example 2: Using an explicit loop for removing elements
	length := len(data)
	for i := 0; i < length; i++ {
		if data[i] % 2 == 0 {
			data = append(data[:i], data[i+1:]...)
			i-- // Decrease index to re-check the current element
			length--
		}
	}
	fmt.Println("Modified data:", data)
}
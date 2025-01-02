package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	
	// Iterate over elements only
	for _, value := range numbers {
		fmt.Println(value)
	}
	
	// Iterate over indices and elements
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
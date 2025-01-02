package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Iterate over indices and values
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

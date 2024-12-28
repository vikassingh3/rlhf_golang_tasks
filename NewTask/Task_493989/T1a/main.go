package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	start := time.Now()

	// Add elements from 0 to a large number (e.g., 10 million)
	const numElements = 10000000
	for i := 0; i < numElements; i++ {
		m[i] = i
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken to add %d elements to the map: %s\n", numElements, elapsed)
}

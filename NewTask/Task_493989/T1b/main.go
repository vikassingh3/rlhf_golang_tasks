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
	fmt.Printf("Time taken to add %d elements to the map (with resizing)   : %s\n", numElements, elapsed)

        start = time.Now()
	mm := make(map[int]int, 10000000)
	// Add elements from 0 to a large number (e.g., 10 million)
	for i := 0; i < numElements; i++ {
		mm[i] = i
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken to add %d elements to the map (without resizing): %s\n", numElements, elapsed)

}
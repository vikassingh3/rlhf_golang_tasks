package main

import (
	"fmt"
	"time"
)

// Simulate some CPU-intensive computation
func dummyComputing(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	// Number of iterations for testing
	n := 1_000_000

	// Measure execution time of dummyComputing
	start := time.Now()
	dummyComputing(n)
	end := time.Now()
	duration := end.Sub(start)

	fmt.Printf("dummyComputing(%d) took: %s\n", n, duration)
}
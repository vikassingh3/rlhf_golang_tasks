package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	n := 1_000_000
	var sum int
	
	// Example 1: Range loop
	for i, _ := range make([]int, n) {
		sum += i
	}
	fmt.Println("Range loop time:", time.Since(start))
	start = time.Now()
	
	// Example 2: Traditional for loop
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println("Traditional for loop time:", time.Since(start))
}
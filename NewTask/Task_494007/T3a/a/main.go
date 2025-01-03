package main

import (
	"fmt"
	"time"
)

func sumDigits(n int) int {
	s := 0
	for i := n; i > 0; i /= 10 {
		s += i % 10
	}
	return s
}

func main() {
	var total int
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		total += sumDigits(i)
	}
	duration := time.Since(start)
	fmt.Printf("Total: %d, Duration: %s\n", total, duration)
}
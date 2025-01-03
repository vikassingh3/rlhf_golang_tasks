package main

import (
	"fmt"
	"time"
)

func main() {
	var total int
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		s := 0
		for j := i; j > 0; j /= 10 {
			s += j % 10
		}
		total += s
	}
	duration := time.Since(start)
	fmt.Printf("Total: %d, Duration: %s\n", total, duration)
}
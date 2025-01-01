package main

import (
	"fmt"
	"time"
)

func processBatch(data []int) []int {
	var results []int
	
	for _, d := range data {
		result := d * d
		results = append(results, result)
	}
	return results
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	
	// Simulate batch processing time
	fmt.Println("Batch processing started...")
	time.Sleep(time.Second)
	
	results := processBatch(data)
	fmt.Println("Batch processing completed:", results)
}
package main

import (
	"fmt"
	"sync"
	"time"
)

// Field represents agricultural data for a single field.
type Field struct {
	Name string
	Acres float64
	Yield float64
}

// CalculateTotalYield calculates the total yield from a slice of fields.
func CalculateTotalYield(fields []Field) float64 {
	var totalYield float64
	for _, field := range fields {
		totalYield += field.Yield
	}
	return totalYield
}

func main() {
	// Simulated agricultural data.
	fields := make([]Field, 10000)
	for i := range fields {
		fields[i] = Field{
			Name: fmt.Sprintf("Field %d", i+1),
			Acres: float64(i + 1),
			Yield: float64(i + 1) * 10,
		}
	}

	// Timing the sequential calculation.
	start := time.Now()
	totalYield := CalculateTotalYield(fields)
	end := time.Now()

	fmt.Printf("Sequential calculation total yield: %.2f, Time taken: %v\n", totalYield, end.Sub(start))

	// Using concurrency to parallelize the calculation.
	var wg sync.WaitGroup
	var parallelTotalYield float64
	partSize := len(fields) / 4

	var mu sync.Mutex // Add a mutex for thread-safe updates

	for i := 0; i < 4; i++ {
		startIndex := i * partSize
		endIndex := min((i+1)*partSize, len(fields))

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var partialYield float64
			for _, field := range fields[start:end] {
				partialYield += field.Yield
			}
			mu.Lock()
			parallelTotalYield += partialYield
			mu.Unlock()
		}(startIndex, endIndex)
	}

	wg.Wait()

	fmt.Printf("Parallel calculation total yield: %.2f, Time taken: %v\n", parallelTotalYield, end.Sub(start))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
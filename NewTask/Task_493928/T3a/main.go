package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Field struct {
	Name  string
	Acres float64
	Yield float64
}

// ProcessField processes a single field by calculating some metrics.
func ProcessField(field *Field) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate processing time
	field.Yield *= 1.1 // Simulate some processing
}

func main() {
	// Simulated agricultural data
	fields := make([]Field, 10)
	for i := range fields {
		fields[i] = Field{
			Name:  fmt.Sprintf("Field %d", i+1),
			Acres: float64(i + 1),
			Yield: float64(i+1) * 10,
		}
	}

	// Split data into smaller batches to reduce shared resource contention
	numWorkers := 4
	batchSize := (len(fields) + numWorkers - 1) / numWorkers // Ensure proper batch distribution
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * batchSize
		end := min(start+batchSize, len(fields))

		// Process each batch in parallel using goroutines
		wg.Add(1)
		go func(batch []Field) {
			defer wg.Done()
			for j := range batch {
				ProcessField(&batch[j])
			}
		}(fields[start:end])
	}

	wg.Wait()

	// Output the processed data
	for _, field := range fields {
		fmt.Printf("Field: %s, Acres: %.2f, Yield: %.2f\n", field.Name, field.Acres, field.Yield)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sync"
)

type Field struct {
	Name  string
	Yield float64
}

type Result struct {
	Field string
	Yield float64
}

func main() {
	// Simulate agricultural fields
	fields := make([]Field, 10000)
	for i := range fields {
		fields[i] = Field{Name: fmt.Sprintf("Field-%d", i), Yield: float64(i + 1)}
	}

	resultCh := make(chan []Result)
	var wg sync.WaitGroup

	// Start worker goroutines
	const numWorkers = 4
	batchSize := len(fields) / numWorkers
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * batchSize
		end := start + batchSize
		if i == numWorkers-1 {
			end = len(fields) // Handle any remaining fields in the last batch
		}

		go func(fieldsBatch []Field) {
			defer wg.Done()
			calculateYieldConcurrentlyBatch(fieldsBatch, resultCh)
		}(fields[start:end])
	}

	// Close the channel once all workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Gather and process results
	totalYield := 0.0
	for batch := range resultCh {
		for _, result := range batch {
			totalYield += result.Yield
		}
	}

	fmt.Printf("Total yield: %.2f\n", totalYield)
}

func calculateYieldConcurrentlyBatch(fields []Field, resultCh chan []Result) {
	const batchSize = 100
	var batch []Result

	for _, field := range fields {
		result := Result{Field: field.Name, Yield: field.Yield * 10}
		batch = append(batch, result)

		if len(batch) == batchSize {
			resultCh <- batch
			batch = make([]Result, 0, batchSize)
		}
	}

	// Send the remaining batch, if any
	if len(batch) > 0 {
		resultCh <- batch
	}
}

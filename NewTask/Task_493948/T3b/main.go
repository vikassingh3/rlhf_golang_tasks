package main

import (
	"fmt"
	"math/rand"
)  
type AggregationResult struct {
    Total int
    Error error
}
func aggregateData(data []int, result chan AggregationResult) {  
    total := 0
    for _, value := range data {
        total += value
    }
    result <- AggregationResult{Total: total, Error: nil}
}  
func main() {  
    numTasks := 3
    dataSize := 1000000
    data := fetchData(dataSize)
    results := make(chan AggregationResult, numTasks)

    // Start workers
    for i := 0; i < numTasks; i++ {
        chunkSize := len(data) / numTasks
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(data) {
            end = len(data)
        }
        go aggregateData(data[start:end], results)
    }

    // Collect results
    total := 0
    var errors []error
    for i := 0; i < numTasks; i++ {
        result := <-results
        if result.Error != nil {
            errors = append(errors, result.Error)
        } else {
            total += result.Total
        }
    }

    // Handle errors
    if len(errors) > 0 {
        fmt.Println("Errors occurred during aggregation:")
        for _, err := range errors {
            fmt.Println(err)
        }
        return
    }

    fmt.Println("Final Aggregated Result:", total)
}  
func fetchData(size int) []int {
    data := make([]int, size)
    for i := 0; i < size; i++ {
        data[i] = rand.Intn(100)
    }
    return data
}
  
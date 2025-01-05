package main

import (
    "fmt"
    "sync"
    "time"
)

// Simulate a serverless function that processes data
func handleRequest(data []int, wg *sync.WaitGroup) {
    defer wg.Done()
    // Simulate processing time
    time.Sleep(100 * time.Millisecond)
    total := 0
    for _, value := range data {
        total += value
    }
    // Simulate sending the result back
    fmt.Printf("Aggregated result for data: %v is %d\n", data, total)
}

func main() {
    const numRequests = 10
    var wg sync.WaitGroup

    // Create random data sets for aggregation
    dataSets := make([][]int, numRequests)
    for i := 0; i < numRequests; i++ {
        dataSets[i] = make([]int, 1000)
        for j := 0; j < 1000; j++ {
            dataSets[i][j] = j + 1
        }
    }

    // Start processing requests concurrently
    for _, data := range dataSets {
        wg.Add(1)
        go handleRequest(data, &wg)
    }

    // Wait for all requests to complete
    wg.Wait()

    fmt.Println("All requests processed.")
}
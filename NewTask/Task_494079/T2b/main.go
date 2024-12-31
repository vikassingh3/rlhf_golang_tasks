package main

import (
	"fmt"
	"sync"
)

// Simulate analyzing a chunk of data
func analyzeChunk(data []int, analysisResults *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {  
    defer wg.Done()
    // Simulate analysis work
    result := analyzeData(data)
    fmt.Printf("Analyzed chunk: %v, Result: %d\n", data, result)

    // Acquire lock before updating analysisResults
    mu.Lock()
    defer mu.Unlock()
    *analysisResults = append(*analysisResults, result)
}

// Aggregate results from all chunks
func aggregateResults(analysisResults []int) int {
    total := 0
    for _, result := range analysisResults {
        total += result
    }
    return total
}

func analyzeData(data []int) int {
    // Simulate simple analysis (e.g., sum of elements)
    sum := 0
    for _, value := range data {
        sum += value
    }
    return sum
}

func main() {  
    // Create a large dataset
    dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    
    chunkSize := 5
    numChunks := len(dataset) / chunkSize
    chunks := make([][]int, numChunks)
    for i := range chunks { 
        start := i * chunkSize
        end := start + chunkSize
        chunks[i] = dataset[start:end]
    } 
    
    var wg sync.WaitGroup 
    var mu sync.Mutex // Mutex to synchronize access to analysisResults
    analysisResults := make([]int, 0)
    
    for _, chunk := range chunks {      
        wg.Add(1)
        go analyzeChunk(chunk, &analysisResults, &wg, &mu) 
    }   
    
    wg.Wait() 
    
    // Use the mutex to ensure thread safety during aggregation
    mu.Lock()
    aggregatedResult := aggregateResults(analysisResults)
    mu.Unlock()
    
    fmt.Println("All chunks have been analyzed.")
    fmt.Println("Aggregated Result:", aggregatedResult) 
}  
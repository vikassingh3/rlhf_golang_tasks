package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Define an error type specific to our pipeline
type PipelineError struct {
    message string
}

func (e *PipelineError) Error() string {
    return fmt.Sprintf("Pipeline Error: %s", e.message)
}

// Function to generate random data
func dataGenerator(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        data := rand.Intn(10) - 5
        ch <- data
    }
    close(ch)
}

// Function to process the data. This function can fail if it encounters a negative number.
func dataProcessor(input chan int, output chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for num := range input {
        if num < 0 {
            // err := &PipelineError{"Encountered negative number"}
            output <- -1  // Use -1 to indicate an error occurred
            return
        }
        result := num * 2
        output <- result
    }
}

func main() {
    dataChan := make(chan int)
    errorChan := make(chan int)
    var wg sync.WaitGroup

    // Start the data generator stage
    wg.Add(1)
    go dataGenerator(dataChan, &wg)

    // Start the data processor stage
    wg.Add(1)
    go dataProcessor(dataChan, errorChan, &wg)
    
    // Process the results
    for result := range errorChan {
        if result == -1 {
            // Error occurred
            fmt.Println("Error in data processing:", <-errorChan)
        } else {
            fmt.Println("Processed data:", result)
        }
    }

    wg.Wait()
} 
 
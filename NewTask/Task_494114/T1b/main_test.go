package main

import (
	"sync"
	"testing"
)

func TestDataPipelineErrorHandling(t *testing.T) {
    // Test case 1: Expect an error when negative number is encountered
    dataChan := make(chan int)
    errorChan := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)
    go dataGenerator(dataChan, &wg)
    wg.Add(1)
    go dataProcessor(dataChan, errorChan, &wg)

    // Simulate sending a negative number to the pipeline
    dataChan <- -2

    // Wait for the pipeline to complete
    wg.Wait()

    expectedError := "Encountered negative number"
    actualError := <-errorChan

    if actualError == -1 {
        t.Errorf("Expected error: '%s', but got -1", expectedError)
    } 
} 
 
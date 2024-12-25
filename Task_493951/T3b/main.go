package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataMiningTask struct {
	Name string
}

type TaskResult struct {
	Result interface{}
	Error  error
}

func (task DataMiningTask) Run(wg *sync.WaitGroup, dataIn <-chan interface{}, dataOut chan<- interface{}, results chan<- TaskResult) {
	defer wg.Done()
	fmt.Printf("Starting task: %s\n", task.Name)

	// Simulate task execution time
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	data, ok := <-dataIn
	if !ok {
		fmt.Printf("Task %s failed to receive data.\n", task.Name)
		results <- TaskResult{Error: fmt.Errorf("task %s: data channel closed", task.Name)}
		return
	}

	processedData, err := task.processData(data)
	if err != nil {
		fmt.Printf("Task %s encountered an error: %v\n", task.Name, err)
		results <- TaskResult{Error: err}
		return
	}

	dataOut <- processedData
	results <- TaskResult{Result: processedData}
	fmt.Printf("Task %s completed!\n", task.Name)
}

func (task DataMiningTask) processData(data interface{}) (interface{}, error) {
	// Simulate potential errors or missed data in the "Data Retrieval" task
	if task.Name == "Data Retrieval" && rand.Intn(2) == 0 {
		return nil, errors.New("data retrieval failed")
	}

	// Perform actual task processing here
	// For demonstration purposes, we return the data unchanged.
	return fmt.Sprintf("%s processed data", task.Name), nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	tasks := []DataMiningTask{
		{"Data Retrieval"},
		{"Data Preprocessing"},
		{"Feature Extraction"},
		{"Model Training"},
		{"Result Analysis"},
	}

	// Initial input data channel
	dataIn := make(chan interface{}, 1)
	dataIn <- "Initial Data" // Seed the pipeline with initial data
	// close(dataIn)            // Ensure no further writes to the initial input channel

	results := make(chan TaskResult, len(tasks))

	// Maintain a pipeline of data flow
	var previousOut chan interface{}
	previousOut = dataIn

	for _, task := range tasks {
		dataOut := make(chan interface{}, 1) // Create a new channel for each task's output

		wg.Add(1)
		go task.Run(&wg, previousOut, dataOut, results)

		// Close the previousOut channel to signal no further writes after each task
		go func(ch chan interface{}) {
			wg.Wait()
			close(ch)
		}(previousOut)

		previousOut = dataOut // Set dataOut as the input for the next task
	}

	// Wait for all tasks to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Waiting for all tasks to finish...")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error in task: %v\n", result.Error)
		} else {
			fmt.Println("Task result:", result.Result)
		}
	}

	fmt.Println("All tasks completed. Research data mining process is done!")
}

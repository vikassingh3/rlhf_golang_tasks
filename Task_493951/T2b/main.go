package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DataMiningTask represents a single data mining task
type DataMiningTask struct {
	Name string
}

// Run performs the data mining task, optionally returning an error
func (task DataMiningTask) Run(wg *sync.WaitGroup, dataIn <-chan interface{}, dataOut chan<- interface{}) {
	defer wg.Done()
	fmt.Printf("Starting task: %s\n", task.Name)

	// Simulate task execution time
	time.Sleep(time.Second)

	// Receive data from the previous task
	data, ok := <-dataIn
	if !ok {
		fmt.Printf("Task %s failed to receive data.\n", task.Name)
		return
	}

	// Process the data
	processedData, err := task.processData(data)
	if err != nil {
		fmt.Printf("Task %s encountered an error: %v\n", task.Name, err)
		return
	}

	// Send the processed data to the next task
	dataOut <- processedData
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

	// Create the initial data channel
	dataIn := make(chan interface{}, 1)
	dataIn <- "Initial Data" // Seed the pipeline with initial data

	for i, task := range tasks {
		dataOut := make(chan interface{}, 1) // Create a new channel for the output of each task

		wg.Add(1)
		go task.Run(&wg, dataIn, dataOut)

		// Pass the output channel to the next task as input
		dataIn = dataOut

		// Close the output channel only after the last task
		if i == len(tasks)-1 {
			go func(ch chan interface{}) {
				wg.Wait()
				close(ch)
			}(dataOut)
		}
	}

	fmt.Println("Waiting for all tasks to finish...")
	wg.Wait()
	fmt.Println("All tasks completed. Research data mining process is done!")

	// Process the final result (if available)
	if result, ok := <-dataIn; ok {
		fmt.Println("Final Result:", result)
	} else {
		fmt.Println("No final result available due to errors or missed data.")
	}
}

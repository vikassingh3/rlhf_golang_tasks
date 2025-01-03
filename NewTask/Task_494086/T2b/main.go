package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Simulate task orchestration
	taskOrchestrator()
}

func taskOrchestrator() {
	var wg sync.WaitGroup
	tasks := []func(){
		dataProcessorTask,
		notificationServiceTask,
	}

	// Add the number of tasks to the wait group
	wg.Add(len(tasks))

	// Start the tasks concurrently
	for _, task := range tasks {
		go func(t func()) {
			t()
			wg.Done()
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("All tasks completed. Sending notification.")
}

func dataProcessorTask() {
	// Simulate data processing work
	time.Sleep(2 * time.Second)
	fmt.Println("Data processing task completed.")
}

func notificationServiceTask() {
	// Simulate sending notification
	time.Sleep(3 * time.Second)
	fmt.Println("Notification sent.")
} 
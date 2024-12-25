package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DataMiningTask represents a single data mining task
type DataMiningTask struct {
	Name      string
	Attempts  int
	MaxAttempts int
}

// Run performs the data mining task with retry mechanism
func (task DataMiningTask) Run(wg *sync.WaitGroup, results *[]string) {
	defer wg.Done()
	for attempt := 1; attempt <= task.MaxAttempts; attempt++ {
		task.Attempts++
		fmt.Printf("Starting task: %s (Attempt %d)\n", task.Name, attempt)
		// Simulate task execution with potential error
		err := simulateTask(task.Name)
		if err == nil {
			*results = append(*results, task.Name)
			fmt.Printf("Task %s completed successfully!\n", task.Name)
			break
		}
		fmt.Printf("Task %s failed: %v. Retrying...\n", task.Name, err)
		time.Sleep(1 * time.Second) // Add a delay before retrying
	}
	if task.Attempts == task.MaxAttempts {
		fmt.Printf("Task %s failed after %d attempts.\n", task.Name, task.Attempts)
	}
}

func simulateTask(taskName string) error {
	// Randomly simulate failure with a 30% chance
	if rand.Intn(10) < 3 {
		return fmt.Errorf("Failed to retrieve data for %s", taskName)
	}
	// Simulate task execution time
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	var wg sync.WaitGroup
	tasks := []DataMiningTask{
		{"Data Fetching", 0, 3},
		{"Data Preprocessing", 0, 3},
		{"Feature Extraction", 0, 3},
		{"Model Training", 0, 3},
		{"Result Analysis", 0, 3},
	}

	var results []string

	for _, task := range tasks {
		wg.Add(1)
		go task.Run(&wg, &results)
	}

	fmt.Println("Waiting for all tasks to finish...")
	wg.Wait()

	fmt.Println("All tasks completed.")
	fmt.Printf("Results collected: %v\n", results)

	// Check for consistency (e.g., verify that all expected tasks are in the results)
	if len(results) != len(tasks) {
		fmt.Println("Warning: Not all tasks completed successfully.")
	}
}
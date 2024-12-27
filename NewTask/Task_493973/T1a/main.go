package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func slowTask(wg *sync.WaitGroup, taskName string) {
	defer wg.Done()
	fmt.Println("Starting slow task:", taskName)
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	fmt.Println("Finished slow task:", taskName)
}

func fastTask(wg *sync.WaitGroup, taskName string) {
	defer wg.Done()
	fmt.Println("Starting fast task:", taskName)
	time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
	fmt.Println("Finished fast task:", taskName)
}

func main() {
	// Create a WaitGroup for slow tasks
	slowWg := &sync.WaitGroup{}
	// Create a WaitGroup for fast tasks
	fastWg := &sync.WaitGroup{}

	// Add tasks to the slow WaitGroup
	slowTasks := []string{"Task A", "Task B", "Task C"}
	for _, task := range slowTasks {
		slowWg.Add(1)
		go slowTask(slowWg, task)
	}

	// Add tasks to the fast WaitGroup
	fastTasks := []string{"Task D", "Task E", "Task F", "Task G"}
	for _, task := range fastTasks {
		fastWg.Add(1)
		go fastTask(fastWg, task)
	}

	// Wait for all slow tasks to complete
	fmt.Println("Waiting for all slow tasks to complete...")
	slowWg.Wait()
	fmt.Println("All slow tasks completed.")

	// Wait for all fast tasks to complete
	fmt.Println("Waiting for all fast tasks to complete...")
	fastWg.Wait()
	fmt.Println("All fast tasks completed.")

	fmt.Println("All tasks completed.")
}
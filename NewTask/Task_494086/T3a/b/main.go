package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Task 1 completed")
	// Signal that task 2 should start
}

func task2() {
	// Wait for task 1 to signal
	time.Sleep(1 * time.Second)
	fmt.Println("Task 2 completed")
}

var released = false

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task1(&wg)

	// Wait for task 2 to finish executing
	task2()
	wg.Wait()

	fmt.Println("All tasks completed")
}
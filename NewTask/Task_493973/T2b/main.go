
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const maxWorkers = 4

func worker(wg *sync.WaitGroup, tasks chan int) {
	defer wg.Done()
	for task := range tasks {
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println("Worker processed task:", task)
	}
}

func main() {
	runtime.GOMAXPROCS(maxWorkers) // Limit the number of OS threads

	tasks := make(chan int, 100) // Buffered channel to store tasks
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(&wg, tasks)
	}

	// Generate tasks and send them to the channel
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}
  
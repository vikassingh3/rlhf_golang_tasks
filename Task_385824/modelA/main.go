package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numWorkers = 5
)

type task struct {
	message string
}

var (
	taskQueue = make(chan task, 10) // Buffered channel to prevent deadlocks
	wg        sync.WaitGroup
	done      int32 // Atomic flag to signal workers to stop
)

func worker(workerID int) {
	defer wg.Done()
	for {
		select {
		case t, ok := <-taskQueue:
			// Check if the channel is closed
			if !ok {
				fmt.Printf("Worker %d exiting: task queue closed\n", workerID)
				return
			}
			fmt.Printf("Worker %d processing task: %s\n", workerID, t.message)
			time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // Simulate processing
		default:
			// Check if the `done` flag is set
			if atomic.LoadInt32(&done) == 1 {
				fmt.Printf("Worker %d exiting: shutdown signal received\n", workerID)
				return
			}
			time.Sleep(100 * time.Millisecond) // Avoid busy-waiting
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i + 1)
	}

	// Generate and add tasks to the queue
	for i := 0; i < 20; i++ {
		taskQueue <- task{message: fmt.Sprintf("Task %d", i+1)}
	}

	// Allow some time for workers to process tasks
	time.Sleep(5 * time.Second)

	// Signal workers to stop and close the queue
	atomic.StoreInt32(&done, 1)
	close(taskQueue)

	// Wait for workers to complete
	wg.Wait()

	fmt.Println("All tasks processed.")
}

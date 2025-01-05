package main  
    
import (  
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)  

// Given function
func processTask(i int, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// Randomly introduce an error
	if rand.Intn(10) == 0 {
		errCh <- fmt.Errorf("task %d failed", i)
		return
	}
	fmt.Println("Task", i, "completed successfully.")
}

func main() {  
	// Number of tasks
	const numTasks = 20
	var wg sync.WaitGroup
	errCh := make(chan error, numTasks)

	// Adding tasks to the waitgroup
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go processTask(i, &wg, errCh)
	}

	// wait for all tasks to finish or error occurs
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Handling errors
	var firstErr error
	for err := range errCh {
		if firstErr == nil {
			firstErr = err
		}
		log.Println("Error:", err)
		// You can handle errors here, for example, by retrying the failed task or canceling all other tasks.
	}

	if firstErr != nil {
		// Handle the overall failure if any error occurred.
		fmt.Println("Failed due to error:", firstErr)
	} else {
		fmt.Println("All tasks completed successfully.")
	}
}  
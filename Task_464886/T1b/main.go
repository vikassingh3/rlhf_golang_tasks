package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func eventLoop(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		// Simulate some work being done here.
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	const numWorkers = 10

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go eventLoop(ch, wg)
	}

	// Generate and send events
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// Signal workers to stop
	close(ch)

	// Wait for all workers to complete
	wg.Wait()
}
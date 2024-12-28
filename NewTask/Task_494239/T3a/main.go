package main

import (
	"fmt"
	"sync"
	"time"
)

var initialCapacity = 10

func worker(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	data := make([]int, 0, initialCapacity)

	for v := range input {
		data = append(data, v)
		if len(data) == initialCapacity {
			fmt.Println("Worker Data:", data)
			data = data[:0]  // Reset data
		}
	}

	if len(data) > 0 {
		fmt.Println("Worker Data:", data)
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 4
	input := make(chan int)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, input)
	}

	start := time.Now()
	for i := 0; i < 1000; i++ {
		input <- i
	}
	close(input)

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Total time elapsed:", duration)
}




// data := make([]int, 0, initialCapacity)
// data = data[:0]  // Reset length, retain capacity


// channel := make(chan int)
// go producer(channel)
// go consumer(channel)

// import "sync/atomic"

// var counter int32

// func increment() {
//     atomic.AddInt32(&counter, 1)
// }


// var rwMutex sync.RWMutex

// func read() {
//     rwMutex.RLock()
//     defer rwMutex.RUnlock()
//     // Perform read operations
// }

// func write() {
//     rwMutex.Lock()
//     defer rwMutex.Unlock()
//     // Perform write operations
// }
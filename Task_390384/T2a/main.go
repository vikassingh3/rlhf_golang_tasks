package main

import (
	"fmt"
	"sync"
)

// generate sends numbers 1 to n to the output channel
func generate(n int, out chan<- int) {
	defer close(out)
	for i := 1; i <= n; i++ {
		out <- i
	}
}

// doubleWorker reads from the input channel, doubles the numbers, and sends them to the output channel
func doubleWorker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		out <- num * 2
	}
}

// print reads numbers from the input channel and prints them
func print(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	const (
		numItems   = 10 // Number of integers to generate
		numWorkers = 4  // Number of workers for the double stage
		bufferSize = 10 // Buffer size for channels
	)

	// Create channels
	genChan := make(chan int, bufferSize) // Buffered for better performance
	doubleChan := make(chan int, bufferSize)

	// Start the generate stage
	go generate(numItems, genChan)

	// Launch worker pool for the double stage
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doubleWorker(genChan, doubleChan, &wg)
	}

	// Close doubleChan once all workers finish
	go func() {
		wg.Wait()
		close(doubleChan)
	}()

	// Start the print stage
	print(doubleChan)
}

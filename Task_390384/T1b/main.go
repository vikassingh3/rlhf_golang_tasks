package main

import (
	"fmt"
	"sync"
)

// generate produces numbers from 1 to n and sends them to the given channel
func generate(n int, ch chan<- int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch) // Signal that no more values will be sent
}

// double takes an int from a channel, doubles it, and sends it to the next channel
func double(in <-chan int, out chan<- int) {
	for num := range in {
		doubled := num * 2
		out <- doubled
	}
	close(out) // Signal that no more values will be sent
}

// print takes an int from a channel and prints it
func print(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	// Create channels for the pipeline
	genChan := make(chan int)
	doubleChan := make(chan int)
	// printChan := make(chan int)

	// Launch goroutines for each pipeline stage
	var wg sync.WaitGroup
	wg.Add(3) // Increment for generate, double, and print

	go func() {
		generate(10, genChan)
		wg.Done() // Decrement when generate is complete
	}()

	go func() {
		double(genChan, doubleChan)
		wg.Done() // Decrement when double is complete
	}()

	go func() {
		print(doubleChan)
		wg.Done() // Decrement when print is complete
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

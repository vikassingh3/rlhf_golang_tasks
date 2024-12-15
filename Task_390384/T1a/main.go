package main

import (
	"fmt"
	"sync"
)

// generator produces a sequence of numbers and sends them to the output channel.
func generator(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i // Send number to output channel
	}
	close(out) // Close the channel when done
}

// doubler receives numbers from the input channel, doubles them, and sends them to the output channel.
func doubler(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when it returns
	for num := range in {
		out <- num * 2 // Double the number and send it to the output channel
	}
	close(out) // Close the output channel when done
}

// printer receives numbers from the input channel and prints them.
func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num) // Print the number
	}
}

func main() {
	numChan := make(chan int)    // Channel for numbers
	doubleChan := make(chan int) // Channel for doubled numbers
	var wg sync.WaitGroup        // WaitGroup to wait for goroutines to finish

	wg.Add(1)                            // Add one goroutine for the doubler
	go doubler(numChan, doubleChan, &wg) // Start the doubler goroutine

	go generator(numChan)  // Start the generator goroutine
	go printer(doubleChan) // Start the printer goroutine

	wg.Wait() // Wait for the doubler to finish
}

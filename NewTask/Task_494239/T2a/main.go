package main

import (
	"fmt"
	"sync"
)

var slicePool sync.Pool

func init() {
	// Initialize the pool with slices of a certain capacity
	for i := 0; i < 10; i++ {
		slicePool.Put(make([]int, 100))
	}
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// Get a slice from the pool
		slice := slicePool.Get().([]int)
		defer slicePool.Put(slice)

		// Use the slice
		slice = slice[:0] // Clear the slice, keeping its capacity
		slice = append(slice, i, i+1, i+2)
		fmt.Println("Worker added:", slice)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)

	wg.Wait()
}
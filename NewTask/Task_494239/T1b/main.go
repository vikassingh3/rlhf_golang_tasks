package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var sharedSlice []int

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutex.Lock()
		// sharedSlice = append(sharedSlice, i)
        sharedSliceCopy := make([]int, len(sharedSlice))
        copy(sharedSliceCopy, sharedSlice)
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	mutex.Lock()
	fmt.Println("Shared slice:", sharedSlice)
	mutex.Unlock()
}
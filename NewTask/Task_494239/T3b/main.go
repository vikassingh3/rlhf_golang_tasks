package main

import (
	"fmt"
	"sync"
)

type workerData struct {
	mutex sync.Mutex
	slice []int
}

func (wd *workerData) reset() {
	wd.mutex.Lock()
	defer wd.mutex.Unlock()
	wd.slice = wd.slice[:0]
}

func (wd *workerData) append(data int) {
	wd.mutex.Lock()
	defer wd.mutex.Unlock()
	wd.slice = append(wd.slice, data)
}

func worker(wg *sync.WaitGroup, wd *workerData) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		wd.reset()
		for j := 0; j < 10; j++ {
			wd.append(i*10 + j)
		}
		// Process the data in the slice
		fmt.Println("Worker Data:", wd.slice)
	}
}

func main() {
	var wg sync.WaitGroup
	wd := &workerData{
		slice: make([]int, 0, 100), // Allocate a large capacity slice initially
	}

	numWorkers := 4
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(&wg, wd)
	}

	wg.Wait()
}
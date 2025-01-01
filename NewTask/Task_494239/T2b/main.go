package main

import (
	"fmt"
	"sync"
)

const (
	initialCapacity = 5
)

type workerData struct {
	slice []int
}

func (wd *workerData) reset() {
	wd.slice = wd.slice[:0]
}

func (wd *workerData) append(data int) {
	wd.slice = append(wd.slice, data)
}

func worker(wg *sync.WaitGroup, wd *workerData) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
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
	numWorkers := 4
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		wd := &workerData{
			slice: make([]int, 0, initialCapacity),
		}
		go worker(&wg, wd)
	}

	wg.Wait()
}

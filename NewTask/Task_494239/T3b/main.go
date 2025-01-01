package main

import (
	"fmt"
	"sync"
)

type workerData struct {
	mu    sync.Mutex
	data  []int
}

func (wd *workerData) append(value int) {
	wd.mu.Lock()
	defer wd.mu.Unlock()
	wd.data = append(wd.data, value)
}

func (wd *workerData) reset() {
	wd.mu.Lock()
	defer wd.mu.Unlock()
	wd.data = nil
}

func (wd *workerData) print() {
	wd.mu.Lock()
	defer wd.mu.Unlock()
	fmt.Println("Worker Data:", wd.data)
}

func worker(id int, wd *workerData, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		wd.append(id*10 + i)
	}
	wd.print()
	wd.reset()
}

func main() {
	var wg sync.WaitGroup
	wd := &workerData{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, wd, &wg)
	}

	wg.Wait()
}

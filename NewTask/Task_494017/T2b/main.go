package main

import (
	"fmt"
	"sync"
)

var counter int

func increment() {
	counter++
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
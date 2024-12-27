
package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg1 *sync.WaitGroup) {
	defer wg1.Done()
	time.Sleep(time.Second * 2) // Simulate work
	fmt.Println("Task 1 completed.")
}

func task2(wg1 *sync.WaitGroup, wg2 *sync.WaitGroup) {
	defer wg1.Done()
	defer wg2.Done()

	time.Sleep(time.Second * 1) // Simulate work
	fmt.Println("Task 2 completed.")
}

func task3(wg2 *sync.WaitGroup) {
	defer wg2.Done()

	time.Sleep(time.Second * 3) // Simulate work
	fmt.Println("Task 3 completed.")
}

func main() {
	var wg1, wg2 sync.WaitGroup

	// Start worker goroutines with respective WaitGroups
	wg1.Add(1)
	go task1(&wg1)

	wg1.Add(1)
	wg2.Add(1)
	go task2(&wg1, &wg2)

	wg2.Add(1)
	go task3(&wg2)

	// Main goroutine will wait here until wg1 is complete
	wg1.Wait()

	fmt.Println("All tasks from Group 1 completed.")

	// Main goroutine will wait here until wg2 is complete
	wg2.Wait()
	
	fmt.Println("All tasks from Group 2 completed. Application exited.")
}
  
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate fetching data from a service
func fetchData(service string, wg *sync.WaitGroup) []string {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	data := []string{fmt.Sprintf("Data from %s", service)}
	return data
}

func main() {
	services := []string{"service1", "service2", "service3"}
	var wg sync.WaitGroup
	result := make([]string, 0)

	// Start fetching data from all services concurrently
	wg.Add(len(services))
	for _, service := range services {
		go func(service string) {
			data := fetchData(service, &wg)
			result = append(result, data...)
		}(service)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("All services have responded.")
	fmt.Println("Combined result:", result)
}

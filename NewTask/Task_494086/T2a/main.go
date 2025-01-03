package main

import (
	"fmt"
	"sync"
	"time"
)

type ServiceComponent struct {
	name string
}

// Simulate a delayed task in a service component
func (s *ServiceComponent) performTask() {
	time.Sleep(time.Duration(time.Second))
	fmt.Printf("Task from service %s completed\n", s.name)
}

func main() {
	var wg sync.WaitGroup

	// Create instances of multiple service components
	components := []*ServiceComponent{
		&ServiceComponent{name: "ServiceA"},
		&ServiceComponent{name: "ServiceB"},
		&ServiceComponent{name: "ServiceC"},
	}

	// Add a WaitGroup count for each service component
	for _, component := range components {
		wg.Add(1)
		go func(c *ServiceComponent) {
			defer wg.Done() // Mark the task as done when it's complete
			c.performTask() // Perform the delayed task
		}(component)
	}

	// Wait for all service components to complete their tasks
	wg.Wait()

	fmt.Println("All tasks from service components are complete.")
}
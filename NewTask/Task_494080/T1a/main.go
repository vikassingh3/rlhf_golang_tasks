package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated function from Vehicle Tracking Service
func trackVehicles(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Tracking vehicles...")
	time.Sleep(1 * time.Second)
}

// Simulated function from Route Management Service
func manageRoutes(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Managing routes...")
	time.Sleep(2 * time.Second)
}

// Function in User Request Service that handles a user request
func handleUserRequest() {
	var wg sync.WaitGroup
	
	// Start goroutines for tracking vehicles and managing routes
	wg.Add(2)
	go trackVehicles(&wg)
	go manageRoutes(&wg)
	
	// Wait for both tasks to complete before responding
	wg.Wait()
	
	fmt.Println("User request fulfilled!")
}

func main() {
	// Simulate multiple user requests
	for i := 0; i < 5; i++ {
		handleUserRequest()
	}
}
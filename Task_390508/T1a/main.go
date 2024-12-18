package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	serverAddress = "localhost:8080" // Replace with your server address
	numClients    = 100              // Number of concurrent clients
	requestCount  = 10               // Number of requests per client
	bufferSize    = 1024             // Buffer size for messages
)

var (
	totalRequests   uint64 = 0
	successRequests uint64 = 0
	failureRequests uint64 = 0
	startTime              = time.Now()
)

func client(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the WaitGroup counter is decremented

	for i := 0; i < requestCount; i++ {
		conn, err := net.Dial("tcp", serverAddress)
		if err != nil {
			atomic.AddUint64(&failureRequests, 1)
			fmt.Printf("Client %d: Error connecting to server: %v\n", id, err)
			continue
		}

		// Explicitly close the connection after use
		func() {
			defer conn.Close()

			// Send a request to the server
			message := fmt.Sprintf("Request from client %d, request number %d\n", id, i)
			_, err := conn.Write([]byte(message))
			if err != nil {
				atomic.AddUint64(&failureRequests, 1)
				fmt.Printf("Client %d: Error writing to server: %v\n", id, err)
				return
			}

			// Read response from the server
			response := make([]byte, bufferSize)
			n, err := conn.Read(response)
			if err != nil {
				atomic.AddUint64(&failureRequests, 1)
				fmt.Printf("Client %d: Error reading from server: %v\n", id, err)
				return
			}

			fmt.Printf("Client %d: Received response: %s\n", id, string(response[:n]))
			atomic.AddUint64(&successRequests, 1)
		}()

		// Increment total requests atomically
		atomic.AddUint64(&totalRequests, 1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// Create multiple client connections
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go client(i, &wg)
	}

	// Wait for all clients to complete
	wg.Wait()

	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	// Ensure no division by zero
	totalReq := atomic.LoadUint64(&totalRequests)
	successReq := atomic.LoadUint64(&successRequests)
	if totalReq == 0 {
		fmt.Println("\nNo requests were made.")
		return
	}

	// Display load test results
	fmt.Printf("\nLoad Test Results:\n")
	fmt.Printf("Total requests: %d\n", totalReq)
	fmt.Printf("Success requests: %d\n", successReq)
	fmt.Printf("Failure requests: %d\n", atomic.LoadUint64(&failureRequests))
	fmt.Printf("Total duration: %v\n", totalDuration)
	fmt.Printf("Request rate: %.2f requests/second\n", float64(totalReq)/totalDuration.Seconds())
	fmt.Printf("Success rate: %.2f%%\n", (float64(successReq)/float64(totalReq))*100)
}

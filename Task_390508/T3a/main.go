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
	serverAddress = "localhost:8080"
	numClients    = 100
	requestCount  = 10
	bufferSize    = 1024
	maxWorkers    = 100 // Number of worker goroutines
)

var (
	totalRequests   uint64 = 0
	successRequests uint64 = 0
	failureRequests uint64 = 0
	startTime              = time.Now()
)

// Worker function to handle client tasks
func clientWorker(task <-chan int, results chan<- time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for clientID := range task {
		for i := 0; i < requestCount; i++ {
			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				atomic.AddUint64(&failureRequests, 1)
				fmt.Printf("Client %d: Error connecting to server: %v\n", clientID, err)
				continue
			}

			// Use explicit connection closure
			func() {
				defer conn.Close()

				// Send a request to the server
				message := fmt.Sprintf("Request from client %d, request number %d\n", clientID, i)
				start := time.Now()
				_, err = conn.Write([]byte(message))
				if err != nil {
					atomic.AddUint64(&failureRequests, 1)
					fmt.Printf("Client %d: Error writing to server: %v\n", clientID, err)
					return
				}

				// Read response from the server
				response := make([]byte, bufferSize)
				n, err := conn.Read(response)
				if err != nil {
					atomic.AddUint64(&failureRequests, 1)
					fmt.Printf("Client %d: Error reading from server: %v\n", clientID, err)
					return
				}

				fmt.Printf("Client %d: Received response: %s\n", clientID, string(response[:n]))

				duration := time.Since(start)
				results <- duration

				atomic.AddUint64(&successRequests, 1)
				atomic.AddUint64(&totalRequests, 1)
			}()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	task := make(chan int, numClients)
	results := make(chan time.Duration, numClients*requestCount)

	// Start worker goroutines
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go clientWorker(task, results, &wg)
	}

	// Assign clients to workers
	for i := 0; i < numClients; i++ {
		task <- i
	}

	close(task)

	// Wait for all workers to complete
	wg.Wait()

	close(results)

	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	var totalResponseTime time.Duration
	for duration := range results {
		totalResponseTime += duration
	}

	totalReq := atomic.LoadUint64(&totalRequests)
	successReq := atomic.LoadUint64(&successRequests)
	if totalReq == 0 {
		fmt.Println("\nNo requests were made.")
		return
	}

	fmt.Printf("\nLoad Test Results:\n")
	fmt.Printf("Total requests: %d\n", totalReq)
	fmt.Printf("Success requests: %d\n", successReq)
	fmt.Printf("Failure requests: %d\n", atomic.LoadUint64(&failureRequests))
	fmt.Printf("Total duration: %v\n", totalDuration)
	fmt.Printf("Request rate: %.2f requests/second\n", float64(totalReq)/totalDuration.Seconds())
	fmt.Printf("Success rate: %.2f%%\n", float64(successReq)/float64(totalReq)*100)
	fmt.Printf("Average response time: %v\n", totalResponseTime/time.Duration(totalReq))
}

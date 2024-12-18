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
	rateLimit     = 100 * time.Millisecond // Rate limit of 100 requests per second
)

var (
	totalRequests   uint64 = 0
	successRequests uint64 = 0
	failureRequests uint64 = 0
	startTime              = time.Now()
)

func client(id int, ch chan<- time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(rateLimit)
	defer ticker.Stop()

	for i := 0; i < requestCount; i++ {
		<-ticker.C

		conn, err := net.Dial("tcp", serverAddress)
		if err != nil {
			atomic.AddUint64(&failureRequests, 1)
			fmt.Printf("Client %d: Error connecting to server: %v\n", id, err)
			continue
		}

		// Ensure connection is explicitly closed after each use
		func() {
			defer conn.Close()

			// Send a request to the server
			message := fmt.Sprintf("Request from client %d, request number %d\n", id, i)
			start := time.Now()
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

			duration := time.Since(start)
			ch <- duration

			atomic.AddUint64(&successRequests, 1)
			atomic.AddUint64(&totalRequests, 1)
		}()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	ch := make(chan time.Duration, numClients*requestCount)

	// Create multiple client connections
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go client(i, ch, &wg)
	}

	// Close the channel after all clients are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	var totalResponseTime time.Duration
	for duration := range ch {
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
	fmt.Printf("Success rate: %.2f%%\n", (float64(successReq)/float64(totalReq))*100)
	fmt.Printf("Average response time: %v\n", totalResponseTime/time.Duration(totalReq))
}

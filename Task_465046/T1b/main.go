package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5" // Importing chi package
)

func batchAPIRequests(wg *sync.WaitGroup, userIDs []int) {
	defer wg.Done() // Decrement the wait group counter when the function returns

	// Make API requests for each userID in the batch
	for _, userID := range userIDs {
		// Simulate making an API request by introducing a delay
		time.Sleep(2 * time.Second)
		fmt.Printf("Fetched user %d\n", userID)
	}
}

func main() {
	r := chi.NewRouter()

	// Set up a route to handle the batch API requests
	r.Post("/batch", func(w http.ResponseWriter, r *http.Request) {
		// Create a new WaitGroup
		var wg sync.WaitGroup

		// User IDs to fetch from the API
		userIDs := []int{1, 2, 3, 4, 5}

		// Set the number of Goroutines equal to the number of batches
		batchSize := 2
		wg.Add(batchSize)

		// Start separate Goroutines to make API requests in batches
		for i := 0; i < len(userIDs); i += batchSize {
			end := i + batchSize
			if end > len(userIDs) {
				end = len(userIDs)
			}
			go batchAPIRequests(&wg, userIDs[i:end])
		}

		// Wait for all Goroutines to finish their execution
		wg.Wait()

		// Send a response once all batches have been processed
		w.Write([]byte("All API requests completed!"))
	})

    fmt.Println("running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

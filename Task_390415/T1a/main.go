package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// mockDBCall simulates a long-running database call.
func mockDBCall(ctx context.Context) (string, error) {
	select {
	case <-time.After(2 * time.Second): // Simulating a database call
		return "data retrieved from DB", nil
	case <-ctx.Done(): // Listen for context cancellation
		return "", ctx.Err()
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with a timeout (5 seconds) for this request.
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel() // Ensures that cancel is called to clean up

	select {
	case result := <-func() chan string {
		out := make(chan string)
		go func() {
			data, err := mockDBCall(ctx)
			if err != nil {
				out <- err.Error()
			} else {
				out <- data
			}
		}()
		return out
	}():
		fmt.Fprintf(w, "Result: %s\n", result)
	case <-ctx.Done(): // Timeout or cancellation
		http.Error(w, ctx.Err().Error(), http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/fetchData", requestHandler)
	fmt.Println("Server is running on localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// APIResponse represents the structure of the API response
type APIResponse struct {
	Message string `json:"message"`
}

func getJSON(ctx context.Context, url string, target interface{}) error {
	client := &http.Client{
		Timeout: 5 * time.Second, // Set a timeout for the HTTP client
	}

	// Create a request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	// Perform the GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for successful response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get data: status code %d", resp.StatusCode)
	}

	// Decode the response into the target struct
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var response APIResponse
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the API request
	err := getJSON(ctx, url, &response)
	if err != nil {
		// Check if the error is a timeout
		if err == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else {
			fmt.Printf("Request failed: %v\n", err)
		}
		return
	}

	// Print the response data
	fmt.Println(response)
}

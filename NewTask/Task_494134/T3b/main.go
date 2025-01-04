package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"net"
)

// APIResponse is a struct representing the response from the API.
type APIResponse struct {
	Message string `json:"message"`
}

func getJSON(url string, target interface{}, timeout time.Duration) error {
	// Set the HTTP timeout.
	client := &http.Client{Timeout: timeout}

	// Make the GET request.
	resp, err := client.Get(url)
	if err != nil {
		// Check if the error is a timeout.
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return fmt.Errorf("get: timeout")
		}
		return err
	}
	defer resp.Body.Close()

	// Check the status code.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	// Decode the response into the target struct.
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return err
	}

	return nil
}

func main() {
	var response APIResponse
	apiURL := "https://jsonplaceholder.typicode.com/posts/1"

	// Set the desired timeout duration. In this case, it's 5 seconds.
	timeout := 5 * time.Second

	err := getJSON(apiURL, &response, timeout)
	if err != nil {
		switch err.Error() {
		case "get: timeout":
			fmt.Println("Error: Request timed out.")
		default:
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	// Print the decoded response.
	fmt.Println(response)
}

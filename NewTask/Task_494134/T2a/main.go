package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIResponse is a struct that represents the response from the API.
// The `json:"message,omitempty"` tag allows the field to be optional.
type APIResponse struct {
	Message string `json:"message,omitempty"`
	ID      int    `json:"id,omitempty"`
}

// getJSON makes a GET request to the specified URL and decodes the response into the target struct.
func getJSON(url string, target interface{}) error {
	// Make the GET request.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the status code.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	// Decode the response into the target struct.
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber() // Optionally, use Number to handle unexpected JSON types
	err = decoder.Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var response APIResponse

	// Make the GET request and decode the response.
	err := getJSON("https://jsonplaceholder.typicode.com/posts/1", &response)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Check for missing fields or unexpected JSON types
	if response.ID == 0 {
		fmt.Println("Warning: ID field is missing from the response.")
	}

	// Print the decoded response.
	fmt.Println(response)
}
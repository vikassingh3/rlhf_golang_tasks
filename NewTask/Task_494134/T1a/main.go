package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User represents a user structure.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// fetchUser retrieves a user from the specified URL and returns a User struct.
func fetchUser(url string) (*User, error) {
	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Decode the response body into a User struct
	var user User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return &user, nil
}

func main() {
	// Define the URL to fetch the user from
	userURL := "https://api.example.com/users/1"

	// Fetch the user
	user, err := fetchUser(userURL)
	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)
		return
	}

	// Print the user information
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("User Name: %s\n", user.Name)
	fmt.Printf("User Email: %s\n", user.Email)
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func fetchUsersFromAPI() ([]User, error) {
	// Define the API endpoint
	apiURL := "https://api.example.com/users"

	// Make an HTTP GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the HTTP response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the JSON response into a slice of User structs
	var users []User
	err = json.Unmarshal(bodyBytes, &users)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return users, nil
}

func displayUsers(users []User) {
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func main() {
	// Fetch users from the API
	users, err := fetchUsersFromAPI()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		os.Exit(1)
	}

	// Display the fetched users
	displayUsers(users)
}
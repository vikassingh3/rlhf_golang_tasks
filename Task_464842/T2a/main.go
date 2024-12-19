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

func updateUsers(users []User) ([]User, error) {
	// Define the new data to be added to the slice
	newData := []byte(`
		[
			{ "id": 4, "name": "User Four", "email": "user4@example.com" },
			{ "id": 5, "name": "User Five", "email": "user5@example.com" }
		]
	`)

	// Parse the new JSON data into a slice of User structs
	var newUsers []User
	err := json.Unmarshal(newData, &newUsers)
	if err != nil {
		return nil, fmt.Errorf("error parsing new JSON response: %v", err)
	}

	// Append the new users to the existing slice
	users = append(users, newUsers...)

	return users, nil
}

func main() {
	// Fetch initial users from the API
	users, err := fetchUsersFromAPI()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		os.Exit(1)
	}

	// Display the initial fetched users
	fmt.Println("Initial Users:")
	displayUsers(users)

	// Update the slice with new data
	users, err = updateUsers(users)
	if err != nil {
		fmt.Println("Error updating users:", err)
		os.Exit(1)
	}

	// Display the updated users
	fmt.Println("\nUpdated Users:")
	displayUsers(users)
}
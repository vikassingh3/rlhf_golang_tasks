package main

import (
	"encoding/json"
	"fmt"
)

// User struct definition remains the same

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func updateUsers(users []User) ([]User, error) {
	// Simulating fetching new data from API
	newData := []byte(`[{"id":4,"name":"User Four"}]`) // Missing the "email" field in the first user

	// Parse the new JSON data into a slice of User structs
	var newUsers []User
	if err := json.Unmarshal(newData, &newUsers); err != nil {
		// Check if the error is due to invalid JSON format
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			return nil, fmt.Errorf("error parsing new JSON response: field '%s' has unexpected type %s", jsonErr.Field, jsonErr.Type.String())
		}
		return nil, fmt.Errorf("error parsing new JSON response: %v", err)
	}

	// Append the new users to the existing slice
	users = append(users, newUsers...)
	return users, nil
}

func main() {
	// rest of the code remains the same
}

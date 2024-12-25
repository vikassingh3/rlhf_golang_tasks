package main

import (
	"errors"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(id int) (*User, error) {
	// Simulate error scenario for user ID not found
	if id == 101 {
		return nil, errors.New("user not found")
	}

	// Simulate successful user retrieval
	user := &User{
		ID:    id,
		Name:  "Sample User",
		Email: "sample@example.com",
	}
	return user, nil
}

func main() {
	user, err := GetUser(102)
	if err != nil {
		println("Error retrieving user:", err)
		return
	}

	println("User retrieved successfully:", user)
}
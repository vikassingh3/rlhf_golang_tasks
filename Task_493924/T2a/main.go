package main

import (
	"fmt"
	"errors"
)

// Struct for user data
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

// Function to retrieve a user's profile from a database
func GetUserProfile(userID int) (*User, error) {
	// Simulate database lookup
	users := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
		2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
	}

	user, found := users[userID]
	if !found {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Function to retrieve a user's profile with phone number
func GetUserProfileWithPhone(userID int) (*UserWithPhone, error) {
	// If we needed to adjust the logic to fetch and include phone numbers, we could do so here
	// for simplicity, we'll just return the existing user with a placeholder phone number
	user, err := GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	return &UserWithPhone{ID: user.ID, Name: user.Name, Email: user.Email, Phone: "+123456789"}, nil
}

// Struct for user data with phone number
type UserWithPhone struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func main() {
	// Example usage
	user, err := GetUserProfile(2)
	if err != nil {
		fmt.Println("Error retrieving user profile:", err)
		return
	}

	fmt.Printf("User Profile: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)

	// Adding more functionality: exporting JSON
	// You need to use "encoding/json" to actually export JSON
	// For now, we will just print the user directly
	fmt.Printf("User in JSON: %+v\n", user)

	// Now let's say we need to add a new field (e.g., "PhoneNumber") to the User struct
	userWithPhone, err := GetUserProfileWithPhone(2)
	if err != nil {
		fmt.Println("Error retrieving user profile with phone:", err)
		return
	}

	fmt.Printf("User Profile with Phone: ID=%d, Name=%s, Email=%s, Phone=%s\n", userWithPhone.ID, userWithPhone.Name, userWithPhone.Email, userWithPhone.Phone)
}

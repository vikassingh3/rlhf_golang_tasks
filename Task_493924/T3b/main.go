package main

import (
	"fmt"
	"errors"
	"sync"
	"encoding/json"
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

// Struct for custom error type
type DivideError struct {
	Message string
	Divisor float64
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("divisor (%f) cannot be zero: %s", e.Divisor, e.Message)
}

// Function that returns a struct (with no error) and an error
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &DivideError{Message: "cannot divide by zero", Divisor: divisor}
	}
	if dividend < 0 {
		return 0, errors.New("dividend cannot be negative")
	}
	result := dividend / divisor
	return result, nil
}

func main() {
	// Example usage for GetUserProfile and GetUserProfileWithPhone
	user, err := GetUserProfile(2)
	if err != nil {
		fmt.Println("Error retrieving user profile:", err)
		return
	}

	// Convert user to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling user:", err)
		return
	}

	fmt.Printf("User Profile: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
	fmt.Printf("User in JSON: %s\n", userJSON)

	userWithPhone, err := GetUserProfileWithPhone(2)
	if err != nil {
		fmt.Println("Error retrieving user profile with phone:", err)
		return
	}

	fmt.Printf("User Profile with Phone: ID=%d, Name=%s, Email=%s, Phone=%s\n", userWithPhone.ID, userWithPhone.Name, userWithPhone.Email, userWithPhone.Phone)

	// Example usage for the divide function
	result, err := divide(10.0, 0.0)
	if err != nil {
		if divideErr, ok := err.(*DivideError); ok {
			fmt.Printf("Error: divisor (%f) cannot be zero\n", divideErr.Divisor)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Result:", result)
	}

	// Handle the error scenario separately
	result, err = divide(-5.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Concurrent slice processing with sync
	var wg sync.WaitGroup
	var slice []int
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice = append(slice, i)
		}(i)
	}
	wg.Wait()
	for i, value := range slice {
		fmt.Println(i, value)
	}
}

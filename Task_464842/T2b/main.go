package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func fetchUsersFromAPI() ([]User, error) {
	// Initialize an empty slice of User structs
	var users []User

	// Simulate multiple API calls and add new users to the slice dynamically
	for i := 1; i <= 3; i++ {
		newUsers, err := fetchUsersFromAPIHelper()
		if err != nil {
			return nil, err
		}
		users = append(users, newUsers...) // Note the '...' after newUsers. This is the slice expansion operator
	}
	return users, nil
}

// Helper function to fetch users from the API (simplified for demonstration)
func fetchUsersFromAPIHelper() ([]User, error) {
	// Create a list of new users with random data
	newUsers := []User{
		{
			ID:     10 + rand.Intn(100), // Generate random ID between 10 and 109
			Name:   fmt.Sprintf("User %d", rand.Intn(100)), // Generate random name
			Email:  fmt.Sprintf("user%d@example.com", rand.Intn(100)), // Generate random email
		},
	}

	// Sleep for 1 second to introduce a delay between API calls (simulated)
	time.Sleep(1 * time.Second)
	return newUsers, nil
}

func main() {
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Fetch users from the simulated API
	users, err := fetchUsersFromAPI()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	// Print the fetched users
	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

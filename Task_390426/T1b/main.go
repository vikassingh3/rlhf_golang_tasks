package main

import (
	"fmt"
	"my_project/auth"
	"my_project/session"
)

func main() {
	userID := "test_user"

	// Create a new session
	s := session.NewSession(userID)
	if s == nil {
		fmt.Println("Failed to create session: Invalid user")
		return
	}

	fmt.Printf("Session created for user: %s\n", s.UserID)

	// Verify user in the auth package
	isValid := auth.VerifyUser(userID, &auth.Session{UserID: s.UserID})
	fmt.Printf("User verification status: %v\n", isValid)
}

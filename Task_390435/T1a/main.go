package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Simulated JWT validation function
func validateJWT(tokenString string) (bool, error) {
	// Simulate JWT parsing and validation (this is where your actual validation logic goes)
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Normally, you'd retrieve and return the public key for validation
		return nil, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil // Assume the token is valid for this example
}

// Middleware to validate JWT asynchronously
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Create a channel to receive validation result
		resultChan := make(chan error, 1)

		// Validate JWT asynchronously
		go func() {
			isValid, _ := validateJWT(token)
			if !isValid {
				resultChan <- fmt.Errorf("invalid token")
			} else {
				resultChan <- nil
			}
		}()

		select {
		case err := <-resultChan:
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// Token is valid, proceed to the next handler
			next.ServeHTTP(w, r)
		case <-ctx.Done():
			http.Error(w, "JWT validation timeout", http.StatusInternalServerError)
		}
	})
}

func main() {
	http.Handle("/protected", jwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the protected route!"))
	})))

	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

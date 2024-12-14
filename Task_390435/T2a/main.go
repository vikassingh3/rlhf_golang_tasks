package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// Simulated public key cache
	publicKey *rsa.PublicKey
	keyOnce   sync.Once
)

// Function to load or cache the public key (for example, from a remote service)
func loadPublicKey() (*rsa.PublicKey, error) {
	// Simulate loading public key logic (replace with actual loading logic)
	return &rsa.PublicKey{}, nil // Load your public key here
}

// ValidateJWT validates a JWT token asynchronously
func validateJWT(tokenString string) (bool, error) {
	keyOnce.Do(func() {
		var err error
		publicKey, err = loadPublicKey()
		if err != nil {
			log.Fatalf("Failed to load public key: %v", err)
		}
	})

	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm, ensuring it's what you expect
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return false, err
	}

	// Additional claims validation can be done here
	return true, nil
}

// JWT Middleware to handle token validation
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Create a channel for the validation result
		resultChan := make(chan error, 1)

		// Validate JWT asynchronously
		go func() {
			isValid, err := validateJWT(token)
			if err != nil {
				fmt.Println("Error validation token", err.Error())
				return
			}
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
			// Token valid, proceed
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

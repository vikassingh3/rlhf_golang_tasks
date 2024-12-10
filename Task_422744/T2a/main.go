package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Constants
const (
	secretKey           = "your_super_secret_key" // Replace this with a strong and secure secret key
	tokenExpirationTime = time.Second * 10        // Tokens will expire in 10 seconds for demonstration purposes
)

// Function to generate JWT tokens
func generateJWT(claims jwt.MapClaims) (string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Function to validate incoming token
func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return []byte(secretKey), nil
	})

	// Handle parsing error
	if err != nil {
		return nil, err
	}

	// Check token validity
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

// HTTP handler to rotate JWT tokens
func rotateJWT(w http.ResponseWriter, r *http.Request) {
	// Extract token from request headers
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "token not provided", http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " prefix
	if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
		http.Error(w, "invalid token format", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[7:]

	// Validate token
	token, err := validateToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Assert claims type
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusUnauthorized)
		return
	}

	// Check token expiration
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if expirationTime.Sub(time.Now()) < 5*time.Second { // Rotate token if it's about to expire
		// Update claims for new token
		claims["exp"] = time.Now().Add(tokenExpirationTime).Unix()

		// Generate new token
		newTokenString, err := generateJWT(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send response with the new token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Token rotated",
			"token":   newTokenString,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		// Token is valid, no need to rotate
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Token is still valid",
		}
		json.NewEncoder(w).Encode(response)
	}
}

// Main function to set up routes and server
func main() {
	// Set up route for token rotation
	http.HandleFunc("/rotate-jwt", rotateJWT)

	// Start HTTP server
	fmt.Println("JWT Rotation Server Started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

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
	secretKey        = "your_secret_key" // Replace this with a strong, secure secret key
	rotationInterval = time.Second * 5   // Tokens will expire in 5 seconds for demonstration purposes
)

// Struct for JWT token response
type tokenResponse struct {
	Token string `json:"token"`
}

// Function to generate a JWT
func generateJWT(claims jwt.MapClaims) (string, error) {
	// Set token expiration in the claims
	claims["exp"] = time.Now().Add(rotationInterval).Unix() // Set expiration

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token string
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Handler for rotating JWT
func rotateJWT(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	claims := jwt.MapClaims{}

	// Populate the claims with necessary user information
	claims["sub"] = "user123" // Replace this with the actual user ID

	// Generate the token
	token, err := generateJWT(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the token as a response
	response := tokenResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Register the route
	http.HandleFunc("/rotate-jwt", rotateJWT)

	// Start the HTTP server
	fmt.Println("JWT Rotation Server Started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

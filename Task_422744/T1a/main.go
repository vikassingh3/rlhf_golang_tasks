package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "dfghjklkjhgfhjbkn"

func init() {
	if secretKey == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
	}
}

// Generates a JWT token
func generateJWT(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// HTTP handler for JWT rotation
func rotateJWT(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	claims := jwt.MapClaims{
		"sub": "user123",
		"iat": time.Now().Unix(),
	}

	token, err := generateJWT(claims)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not generate token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode token in response
	response := map[string]string{
		"token": token,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("could not encode response: %v", err)
	}
}

func main() {
	// Set up route
	http.HandleFunc("/rotate-jwt", rotateJWT)

	// Start the server
	fmt.Println("JWT Rotation Server Started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

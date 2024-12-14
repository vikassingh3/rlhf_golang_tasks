package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claim defines the structure of JWT claims
type Claim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func main() {
	http.HandleFunc("/protected", protectedHandler)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token required", http.StatusUnauthorized)
		return
	}

	// Validate the token
	claims, err := validateToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Respond with a personalized message
	fmt.Fprintf(w, "Hello, %s!\n", claims.UserID)
}

func validateToken(tokenString string) (*Claim, error) {
	// Parse the token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Provide the secret key
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Verify token validity and extract claims
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		// Asynchronous validation (e.g., against a blacklist or database)
		go asyncValidateToken(claims)
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func asyncValidateToken(claims *Claim) {
	// Simulate a delay for asynchronous token validation
	time.Sleep(2 * time.Second)
	fmt.Printf("Asynchronously validated token for user: %s\n", claims.UserID)
}

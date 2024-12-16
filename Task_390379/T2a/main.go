package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Constants for JWT
const (
	secretKey = "your-secret-key" // Replace with a secure key
)

// Context key type to avoid key collisions
type contextKey string

const requestIDKey contextKey = "requestID"

// issueToken issues a new JWT token with its subject in the context
func issueToken(ctx context.Context, subject string) (string, error) {
	// Create a new JWT token
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    "your-issuer",
		Audience:  []string{"your-audience"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(), // Unique token ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Log token issuance
	log.Printf("[INFO] Token Issued for Subject '%s' (JTI: %s) - RequestID: %s", claims.Subject, claims.ID, ctx.Value(requestIDKey))
	return tokenString, nil
}

// validateToken validates a JWT token and logs the outcome
func validateToken(ctx context.Context, tokenString string) error {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Log token validation error
		log.Printf("[ERROR] Token Validation Failed - Error: %v - RequestID: %s", err, ctx.Value(requestIDKey))
		return err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Log successful token validation
		log.Printf("[INFO] Token Validated for Subject '%s' (JTI: %s) - RequestID: %s", claims.Subject, claims.ID, ctx.Value(requestIDKey))
		return nil
	}

	// Log token validation error
	log.Printf("[ERROR] Token Validation Failed - Invalid Token - RequestID: %s", ctx.Value(requestIDKey))
	return fmt.Errorf("invalid token")
}

// generateRequestID generates a new UUID for request identification
func generateRequestID() string {
	return uuid.New().String()
}

// handler is an HTTP handler that manages the request lifecycle and logging
func handler(w http.ResponseWriter, r *http.Request) {
	// Create a new context with a unique RequestID
	ctx := context.WithValue(r.Context(), requestIDKey, generateRequestID())

	// Example of issuing a token
	token, err := issueToken(ctx, "example-subject")
	if err != nil {
		log.Printf("[ERROR] Failed to issue token - RequestID: %s - Error: %v", ctx.Value(requestIDKey), err)
		http.Error(w, "Failed to issue token", http.StatusInternalServerError)
		return
	}

	// Log the newly issued token
	log.Printf("[INFO] Issued Token: %s - RequestID: %s", token, ctx.Value(requestIDKey))

	// Validate the token
	if err := validateToken(ctx, token); err != nil {
		log.Printf("[ERROR] Token Validation Failed - RequestID: %s - Error: %v", ctx.Value(requestIDKey), err)
		http.Error(w, "Failed to validate token", http.StatusUnauthorized)
		return
	}

	// Send response
	fmt.Fprintf(w, "Token issued and validated successfully.")
}

func main() {
	http.HandleFunc("/token", handler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}

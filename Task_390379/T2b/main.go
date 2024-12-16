package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Constants for JWT
const (
	secretKey = "your-secret-key" // Replace with a secure key
)

// ContextKey is a custom type for context keys to avoid collisions
type ContextKey string

const requestIDKey ContextKey = "requestID"

// issueToken issues a new JWT token and logs the operation
func issueToken(ctx context.Context, subject string) (string, error) {
	// Extract request ID from context
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return "", fmt.Errorf("request ID not found in context")
	}

	// Create JWT claims
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    "your-issuer",
		Audience:  []string{"your-audience"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(), // Unique token ID
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the encoded token string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Log token issuance
	log.Printf("[INFO] Token Issued - Subject: %s, JTI: %s, RequestID: %s", claims.Subject, claims.ID, requestID)
	return tokenString, nil
}

// validateToken validates a JWT token and logs the outcome
func validateToken(ctx context.Context, tokenString string) error {
	// Extract request ID from context
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return fmt.Errorf("request ID not found in context")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Log validation failure
		log.Printf("[ERROR] Token Validation Failed - Error: %v, RequestID: %s", err, requestID)
		return err
	}

	// Extract claims and verify validity
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Log successful validation
		log.Printf("[INFO] Token Validated - Subject: %s, JTI: %s, RequestID: %s", claims.Subject, claims.ID, requestID)
		return nil
	}

	// Log invalid token
	log.Printf("[ERROR] Token Validation Failed - Invalid Token, RequestID: %s", requestID)
	return fmt.Errorf("invalid token")
}

// handleRequest handles a request and logs token operations
func handleRequest(ctx context.Context) {
	// Generate a unique request ID
	requestID := uuid.New().String()
	// Add the request ID to the context
	ctx = context.WithValue(ctx, requestIDKey, requestID)

	// Issue a token
	token, err := issueToken(ctx, "example-subject")
	if err != nil {
		log.Fatalf("[FATAL] Failed to issue token - Error: %v, RequestID: %s", err, requestID)
	}

	// Validate the issued token
	if err := validateToken(ctx, token); err != nil {
		log.Fatalf("[FATAL] Failed to validate token - Error: %v, RequestID: %s", err, requestID)
	}

	log.Printf("[INFO] Token lifecycle completed successfully - RequestID: %s", requestID)
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Handle the request within the context
	handleRequest(ctx)
}

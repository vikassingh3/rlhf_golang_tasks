package main

import (
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

func issueToken(subject string) (string, error) {
	// Create a new JWT token
	claims := &jwt.StandardClaims{
		Subject:   subject,
		Issuer:    "your-issuer",
		Audience:  "your-audience",
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        uuid.New().String(), // Unique token ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Log token issuance
	log.Printf("[INFO] Token Issued for Subject '%s' (JTI: %s) - RequestID: %s", claims.Subject, claims.Id, generateRequestID())
	return tokenString, nil
}

func validateToken(tokenString string) error {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Log token validation error
		log.Printf("[ERROR] Token Validation Failed - Error: %v - RequestID: %s", err, generateRequestID())
		return err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		// Log successful token validation
		log.Printf("[INFO] Token Validated for Subject '%s' (JTI: %s) - RequestID: %s", claims.Subject, claims.Id, generateRequestID())
		return nil
	}

	// Log token validation error
	log.Printf("[ERROR] Token Validation Failed - Invalid Token - RequestID: %s", generateRequestID())
	return fmt.Errorf("invalid token")
}

func generateRequestID() string {
	return uuid.New().String()
}

func main() {
	// Example usage
	token, err := issueToken("example-subject")
	if err != nil {
		log.Fatalf("Failed to issue token: %v", err)
	}

	err = validateToken(token)
	if err != nil {
		log.Fatalf("Failed to validate token: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

const secretKey = "your_secret_key"

// Create a logger
var log = logrus.New()

// Custom claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// IssueToken issues a new JWT token
func IssueToken(ctx context.Context, username string) (string, error) {
	log.WithContext(ctx).Info("Issuing token")

	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myapp",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to sign token")
		return "", err
	}

	log.WithContext(ctx).Info("Successfully issued token")
	return tokenString, nil
}

// ValidateToken validates a JWT token
func ValidateToken(ctx context.Context, tokenString string) (*Claims, error) {
	log.WithContext(ctx).Info("Validating token")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.WithContext(ctx).Error("Unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Token validation failed")
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.WithContext(ctx).Info("Token successfully validated")
		return claims, nil
	}
	log.WithContext(ctx).Error("Invalid token")
	return nil, fmt.Errorf("invalid token")
}

// Handler for issuing tokens
func IssueHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.URL.Query().Get("username")
	token, err := IssueToken(ctx, username)
	if err != nil {
		http.Error(w, "Failed to issue token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Issued Token: %s", token)
}

// Handler for validating tokens
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenString := r.URL.Query().Get("token")
	claims, err := ValidateToken(ctx, tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Token is valid for user: %s", claims.Username)
}

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	http.HandleFunc("/issue", IssueHandler)
	http.HandleFunc("/validate", ValidateHandler)

	log.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

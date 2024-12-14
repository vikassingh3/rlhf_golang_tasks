package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Simulated token cache
var tokenCache = &sync.Map{}

// Validate JWT asynchronously
func validateJWTAsync(tokenString string) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Check token cache
	if cached, ok := tokenCache.Load(tokenString); ok {
		if cacheResult, ok := cached.(bool); ok {
			log.Printf("Token found in cache: %v", cacheResult)
			return cacheResult, nil
		}
	}

	// Simulate fetching a public key or other necessary setup
	pubKeyPEM := []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECiQzT7E1N1Q5N0LaPv7gGuiEkO5Rl
Gj/ZU9F8UNmGR3BXF9+4QamYszjjeC3LnLEhv+KqwGvAqkrsjj1XhjLFQ==
-----END PUBLIC KEY-----`)

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyPEM)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Validate JWT
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		log.Printf("Token validated successfully: %+v", claims)

		// Cache the validated token
		tokenCache.Store(tokenString, true)
		return true, nil
	}

	return false, fmt.Errorf("token validation failed")
}

// Middleware to handle JWT validation
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid or missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		// Validate JWT asynchronously
		isValid, err := validateJWTAsync(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isValid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/protected", jwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the protected route!"))
	})))

	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

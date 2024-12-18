package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CustomClaims struct includes standard claims and custom data
type CustomClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}

// User struct represents the user data
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a new user
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Signing key
	signingKey := []byte("your_secret_key")

	// Create claims with user data
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "your_domain.com",
			Subject:   "user",
			Audience:  "your_application",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated Token:", tokenString)

	// Validate the token
	validateToken(tokenString, signingKey)
}

func validateToken(tokenString string, signingKey []byte) {
	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// Extract and validate the claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println("Token is valid")
		fmt.Printf("User ID: %d\n", claims.User.ID)
		fmt.Printf("User Name: %s\n", claims.User.Name)
		fmt.Printf("User Email: %s\n", claims.User.Email)
	} else {
		fmt.Println("Token is invalid")
	}
}

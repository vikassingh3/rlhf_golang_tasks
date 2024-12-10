package main

import (
	// ... (existing imports)
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{
	"alice": "$2y$10$cLvSmYqatY6aKbZH/PSmxOQjmnvz0tnMVvMc6aMU81Uj5hZcLYmga", // bcrypt hash of "password"
	"bob":   "$2y$10$XsGeGqmlsDhv9x/5ezXOaeas3DdyKXudS9Sk5U8Vt7pf9LoRqB122", // bcrypt hash of "password"
}

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func authenticate(w http.ResponseWriter, r *http.Request) (*tokenClaims, error) {
	// Get the auth token from the request header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("token not provided")
	}

	parts := strings.SplitN(tokenString, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid token format")
	}

	token, err := jwt.ParseWithClaims(parts[1], &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil // Use a more secure way to manage keys in production
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, ok := users[req.Username]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		Username: req.Username,
	}

	token.Claims = claims
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   3600 * 24, // 24 hours
	})

	w.Write([]byte("Logged in successfully. You can now connect to the WebSocket."))
}

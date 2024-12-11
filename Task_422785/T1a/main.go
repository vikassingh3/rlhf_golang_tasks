package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/scrypt"
)

var (
	salt   = []byte("super-secret-salt")
	key    = make([]byte, 32)
	jwtKey = []byte("your-super-secret-jwt-key") // Replace this with a secure random key
)

// generateKey generates a new AES key using scrypt
func generateKey() error {
	var err error
	key, err = scrypt.Key([]byte("your-master-password"), salt, 16384, 8, 1, 32)
	return err
}

func main() {
	if err := generateKey(); err != nil {
		log.Fatal("Error generating key:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/login", loginHandler).Methods("POST")
	r.HandleFunc("/api/data", protectedHandler).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// User struct represents a user
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// tokenClaims holds user claims in JWT
type tokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// loginHandler handles login requests
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := decodeRequestBody(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simulate user authentication
	if user.Email == "example@example.com" && user.ID == "example123" {
		// Generate JWT token
		token, err := generateToken(user.ID)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Send token in response
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Login successful")
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

// generateToken generates a JWT token for the given user ID
func generateToken(userID string) (string, error) {
	claims := &tokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// protectedHandler handles protected routes that require authentication
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
		return
	}

	token, err := parseToken(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
		return
	}

	// Simulate fetching user data from a database
	userData := fmt.Sprintf("User ID: %s, Email: example@example.com", claims.UserID)

	// Encrypt the user data
	encryptedData, err := encryptData([]byte(userData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Encrypted User Data: %s\n", base64.StdEncoding.EncodeToString(encryptedData))
}

// parseToken parses a JWT token string and validates it
func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Method)
		}
		return jwtKey, nil
	})
}

// encryptData encrypts the data using AES-GCM
func encryptData(data []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// decodeRequestBody decodes the request body into a struct
func decodeRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(v)
}

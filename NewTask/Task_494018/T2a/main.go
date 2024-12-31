package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DataItem struct to hold the data
type DataItem struct {
	Value int `json:"value"`
}

// Define the encryption key
const encryptionKey = "mySuperSecretKey123" // Must be 16, 24, or 32 bytes long

// Encrypt data using AES-256
func encryptData(data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Middleware to enforce access controls
func requireAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer mySuperSecretToken" { // Replace with actual token verification logic
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// Process data handler
func processDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data []DataItem
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	encryptedData, err := encryptData([]byte(fmt.Sprintf("%v", data)))
	if err != nil {
		http.Error(w, "Error encrypting data", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"encryptedData": encryptedData}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/process-data", requireAuthorization(processDataHandler))
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

// Define the secret key used for signing JWT tokens
var secretKey = []byte("your_secret_key") // Keep this secret!

// Function to generate a new JWT token
func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Function to validate the incoming JWT token
func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket handler function
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from query parameters
	tokenString := r.URL.Query().Get("token")

	// Validate the token
	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", conn.RemoteAddr().String())

	// Handling incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Respond back to the client
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	// Example of generating a token for a user
	token, err := generateToken("user1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated Token:", token)

	http.HandleFunc("/ws", websocketHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

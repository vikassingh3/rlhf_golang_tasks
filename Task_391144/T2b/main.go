package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4" // Use the updated JWT library
	"github.com/gorilla/websocket"
)

// Define the JWT secret key (replace this with a secure secret in production)
var secretKey = []byte("your_secret_key_here")

// WebSocket upgrade handler
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

// Validate a JWT token
func validateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	return token, err
}

// WebSocket handler
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) <= len("Bearer ") {
		http.Error(w, "Missing or invalid authorization token", http.StatusUnauthorized)
		return
	}

	// Remove the "Bearer " prefix
	tokenString := authHeader[len("Bearer "):]

	// Validate the token
	token, err := validateJWT(tokenString)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired authorization token", http.StatusUnauthorized)
		return
	}

	// If the token is valid, upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", conn.RemoteAddr().String())

	// Handle incoming WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		action, ok := msg["action"].(string)
		if !ok {
			log.Println("Invalid or missing 'action' field in message")
			continue
		}

		log.Printf("Received message: %s, Action: %s", string(message), action)

		// Example response
		response := map[string]string{"status": "ok", "message": "Received your message"}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshaling response:", err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", websocketHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

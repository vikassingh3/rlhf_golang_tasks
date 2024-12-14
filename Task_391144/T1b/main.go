package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Message represents a message structure that can be sent over the WebSocket
type Message struct {
	Text string `json:"text"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // For simplicity, allow any origin. You should restrict this in production.
	},
}

// AuthenticateToken verifies the token provided by the client
func AuthenticateToken(token string) bool {
	// In a real application, you would validate the token against a JWT or another token system.
	// For example, using fernet for symmetric encryption or JWT for asymmetric.

	// Simulating token validation
	return token == "your_valid_token"
}

// WebSocketHandler upgrades the connection and handles WebSocket messages
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from Authorization header (if available)
	token := r.Header.Get("Authorization")
	if token == "" || !AuthenticateToken(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Loop to handle incoming WebSocket messages
	for {
		// Read a message from the WebSocket connection
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			// Handle error and check if it's due to connection closure
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Unexpected WebSocket close error:", err)
			}
			break
		}

		// Switch to handle different types of messages
		switch msgType {
		case websocket.TextMessage:
			var message Message
			// Unmarshal the incoming JSON message into the Message struct
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Println("Error unmarshaling message:", err)
				continue
			}
			log.Printf("Received message: %+v\n", message)

			// Echo the message back to the client
			reply, err := json.Marshal(message)
			if err != nil {
				log.Println("Error marshaling reply:", err)
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, reply); err != nil {
				log.Println("Error writing reply:", err)
				break
			}

		default:
			log.Println("Unexpected message type:", msgType)
		}
	}
}

func main() {
	// Handle WebSocket connections on the /ws endpoint
	http.HandleFunc("/ws", WebSocketHandler)

	log.Println("Starting WebSocket server on :8080")
	// Start the server and listen for incoming connections
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

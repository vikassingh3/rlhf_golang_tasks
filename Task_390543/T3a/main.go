package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Time allowed to perform the read operations for the pong message.
	pingPeriod = (pongWait * 9) / 10
	// Server WebSocket URL.
	wsURLString = "ws://echo.websocket.org" // Use your own server URL for testing
)

var (
	// An upgrader is used to upgrade HTTP requests to WebSocket connections.
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// connect establishes a WebSocket connection and returns the client connection.
func connect(url string) (*websocket.Conn, error) {
	parsedURL, err := url.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	wsURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", wsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	ws, _, err := upgrader.Upgrade(ctx, req, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade connection: %v", err)
	}

	return ws, nil
}

// run maintains a WebSocket connection and sends messages periodically.
func run(url string) {
	log.Printf("Connecting to WebSocket server at %s...", url)

	for {
		ws, err := connect(url)
		if err != nil {
			log.Printf("Failed to connect: %v. Retrying...", err)
			time.Sleep(time.Second * 5)
			continue
		}

		// Enter a loop to keep the connection alive by sending ping messages.
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		pongChan := make(chan *websocket.Message, 1)

		for {
			select {
			case msg := <-pongChan:
				// If we receive a pong, reset the write deadline.
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if msg != nil {
					log.Println("Pong received")
				} else {
					log.Println("Pong received with nil message")
				}
			case <-ws.Pong():
				// For WebSocket version >= 13.
				log.Println("Pong received")
			case <-ticker.C:
				// If it's time to send a ping, do so.
				err := ws.WriteControl(websocket.PingMessage, nil, time.Now().Add(pongWait))
				if err != nil {
					log.Printf("Failed to send ping: %v. Closing connection.", err)
					ws.Close()
					break
				}
			case err := <-ws.CloseNotify():
				log.Printf("Connection closed: %v. Retrying...", err)
				time.Sleep(time.Second * 5)
				goto connectLoop
			}
		}

	connectLoop:
	}
}

func main() {
	run(wsURLString)
}

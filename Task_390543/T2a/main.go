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
	wsURL = "ws://localhost:8080/ws"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func handleWebSocketConnection(ctx context.Context) {
	conn, _, err := upgrader.Upgrade(context.WithCancel(ctx), &http.Request{URL: &http.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}}, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v\n", err)
		return
	}
	defer conn.Close()

	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Context done, closing connection.")
			return
		case <-pingTicker.C:
			if err := writeMessage(conn, websocket.PingMessage, []byte("ping")); err != nil {
				log.Printf("Failed to send ping: %v\n", err)
				return
			}
		default:
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Failed to read message: %v\n", err)
				return
			}

			switch msgType {
			case websocket.PongMessage:
				log.Printf("Received pong message: %v\n", message)
			case websocket.TextMessage:
				log.Printf("Received text message: %s\n", message)
			default:
				log.Printf("Unexpected message type: %d\n", msgType)
			}
		}
	}
}

func writeMessage(conn *websocket.Conn, messageType int, message []byte) error {
	conn.SetWriteDeadline(time.Now().Add(writeWait))
	return conn.WriteMessage(messageType, message)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	go handleWebSocketConnection(ctx)

	select {
	case <-time.After(5 * time.Second):
		log.Println("Shutting down connection after 5 seconds.")
		cancel()
	case <-ctx.Done():
	}
}
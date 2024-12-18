package main

import (
	"context"
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

func measureLatency() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := upgrader.Upgrade(context.WithCancel(ctx), &http.Request{URL: &http.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}}, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
		log.Fatalf("Failed to send ping message: %v", err)
	}

	var pingMessage []byte
	if _, _, err := conn.ReadMessage(); err != nil {
		log.Fatalf("Failed to receive pong message: %v", err)
	}

	if string(pingMessage) != "pong" {
		log.Fatal("Expected pong message")
	}

	// Measure round-trip latency
	startTime := time.Now()
	conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := conn.WriteMessage(websocket.TextMessage, []byte("measure-latency")); err != nil {
		log.Fatalf("Failed to send measure-latency message: %v", err)
	}

	// var latencyMessage []byte
	if _, _, err := conn.ReadMessage(); err != nil {
		log.Fatalf("Failed to receive latency message: %v", err)
	}

	endTime := time.Now()
	latency := endTime.Sub(startTime)
	log.Printf("Latency: %s\n", latency)
}

func main() {
	measureLatency()
}

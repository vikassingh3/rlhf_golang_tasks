package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	// "github.com/gorilla/websocket"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	wsURL      = "ws://echo.websocket.org" // Public WebSocket echo server for testing
)

type wsClient struct {
	conn   *websocket.Conn
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

func (c *wsClient) start(ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.writeControl(websocket.PingMessage, nil)
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}
}

func (c *wsClient) writeControl(mt int, payload []byte) {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := c.conn.WriteControl(mt, payload, time.Now().Add(writeWait)); err != nil {
		log.Printf("Write control error: %v", err)
	}
}

func (c *wsClient) writeMessage(mt int, payload []byte) {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := c.conn.WriteMessage(mt, payload); err != nil {
		log.Printf("Write message error: %v", err)
	}
}

func (c *wsClient) close() {
	c.cancel()
	c.conn.Close()
	c.wg.Wait()
}

func measureLatency() (int, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse WebSocket URL
	u, err := url.Parse(wsURL)
	if err != nil {
		return 0, fmt.Errorf("invalid WebSocket URL: %w", err)
	}

	// Create WebSocket connection
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	client := &wsClient{conn: conn, cancel: cancel}
	defer client.close()

	// Start reading messages in a goroutine
	go client.start(ctx)

	// Send latency measurement message
	start := time.Now()
	client.writeMessage(websocket.TextMessage, []byte("measure-latency"))

	// Wait for the echo response
	messageReceived := make(chan struct{})
	go func() {
		for {
			_, message, err := client.conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			if string(message) == "measure-latency" {
				close(messageReceived)
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("timeout during latency measurement")
	case <-messageReceived:
	}

	// Measure latency
	return int(time.Since(start).Milliseconds()), nil
}

func main() {
	latency, err := measureLatency()
	if err != nil {
		log.Printf("Error during WebSocket communication: %v", err)
		return
	}
	log.Printf("Latency: %d ms\n", latency)
}

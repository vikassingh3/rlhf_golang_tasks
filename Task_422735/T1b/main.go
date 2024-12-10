package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Chat struct {
	clients    map[*websocket.Conn]*Client
	register   chan *Client
	broadcast  chan []byte
	unregister chan *Client
}

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	chat     *Chat
	username string
}

func (c *Chat) run() {
	for {
		select {
		case client := <-c.register:
			c.clients[client.conn] = client
		case message := <-c.broadcast:
			for _, client := range c.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(c.clients, client.conn)
				}
			}
		case client := <-c.unregister:
			if _, ok := c.clients[client.conn]; ok {
				close(client.send)
				delete(c.clients, client.conn)
			}
		}
	}
}

func (c *Client) write() {
	defer c.conn.Close()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// Channel closed, terminate connection
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("write error:", err)
				return
			}
		}
	}
}

func (c *Client) read() {
	defer func() {
		c.chat.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		msg := &Message{Username: c.username, Message: string(message)}
		jsonMessage, err := json.Marshal(msg)
		if err != nil {
			log.Println("json marshal error:", err)
			continue
		}
		c.chat.broadcast <- jsonMessage
	}
}

var chat = &Chat{
	clients:    make(map[*websocket.Conn]*Client),
	register:   make(chan *Client),
	broadcast:  make(chan []byte),
	unregister: make(chan *Client),
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for now (update for production)
		return true
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		conn:     conn,
		send:     make(chan []byte, 256),
		chat:     chat,
		username: r.URL.Query().Get("username"),
	}

	chat.register <- client

	go client.write()
	go client.read()
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	go chat.run()
	http.HandleFunc("/ws", serveWs)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

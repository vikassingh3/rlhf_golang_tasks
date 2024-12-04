package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

type Message struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Message string `json:"message"`
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Invalid JSON message:", err)
			continue
		}

		if err := validate.Struct(msg); err != nil {
			log.Println("Validation failed:", err)
			continue
		}

		fmt.Println("Received message:", msg)

		// Add your message handling logic here

		err = conn.WriteMessage(websocket.TextMessage, []byte("Message received and validated."))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", serveWs)
	fmt.Println("Listining on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

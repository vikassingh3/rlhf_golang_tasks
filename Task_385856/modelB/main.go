package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/xeipuuv/gojsonschema"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins
}

func main() {
	http.HandleFunc("/ws", serveWs)
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	// Load and compile the schema
	schemaData, err := ioutil.ReadFile("schema.json")
	if err != nil {
		log.Println("Error reading schema file:", err)
		return
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaData))
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		log.Println("Error compiling schema:", err)
		return
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			return
		}

		var msg struct {
			Message   string `json:"message"`
			Timestamp int    `json:"timestamp"`
		}

		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		// Validate the message against the schema
		valueLoader := gojsonschema.NewGoLoader(msg)
		result, err := schema.Validate(valueLoader)
		if err != nil {
			log.Println("Error during validation:", err)
			continue
		}

		if result.Valid() {
			fmt.Println("Received valid message:", string(message))
		} else {
			fmt.Println("Received invalid message:")
			for _, err := range result.Errors() {
				fmt.Println(err)
			}
		}
	}
}

package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	NumClients = 100                  // Number of clients to spawn
	ConnHost   = "localhost"          // Server hostname
	ConnPort   = "8080"               // Server port
	Message    = "Hello from client!" // Message to send to the server
)

var wg sync.WaitGroup

func Client(id int) {
	defer wg.Done()
	// Connect to the server
	conn, err := net.Dial("tcp", ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Client", id, ": Error connecting to server:", err)
		return
	}
	defer conn.Close()
	// Send a message to the server
	_, err = conn.Write([]byte(Message))
	if err != nil {
		fmt.Println("Client", id, ": Error sending message:", err)
		return
	}
	// Receive the server's response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Client", id, ": Error receiving response:", err)
		return
	}
	// Print the response
	fmt.Println("Client", id, ": Received response:", string(buffer[:n]))
}

func main() {
	fmt.Println("Load testing TCP server with", NumClients, "clients...")
	// Start multiple clients
	wg.Add(NumClients)
	for i := 0; i < NumClients; i++ {
		go Client(i)
	}
	// Wait for all clients to finish
	wg.Wait()
	fmt.Println("Load testing complete.")
}

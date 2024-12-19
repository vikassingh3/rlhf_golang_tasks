package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Create a TCP listener with a specific TCP options
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:8080")

	// Accept connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Do something with the connection
	fmt.Println("New connection established")
}

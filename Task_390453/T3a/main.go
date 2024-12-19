package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// Set TCP buffers
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	// Establish a TCP connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer conn.Close()

	// Configure buffer sizes using environment variables
	sendBufSize, err := strconv.Atoi(os.Getenv("TCP_SNDBUF"))
	if err != nil || sendBufSize <= 0 {
		sendBufSize = 65536
	}
	recvBufSize, err := strconv.Atoi(os.Getenv("TCP_RCVBUF"))
	if err != nil || recvBufSize <= 0 {
		recvBufSize = 65536
	}

	// Display the configured buffer sizes (unable to set directly in Go's net.Conn)
	fmt.Printf("Using send buffer size: %d bytes\n", sendBufSize)
	fmt.Printf("Using receive buffer size: %d bytes\n", recvBufSize)

	// Implement keep-alive using TCPConn methods
	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		err = tcpConn.SetKeepAlive(true)
		if err != nil {
			fmt.Println("Failed to enable keep-alive:", err)
		}

		err = tcpConn.SetKeepAlivePeriod(5 * time.Second)
		if err != nil {
			fmt.Println("Failed to set keep-alive period:", err)
		}
	} else {
		fmt.Println("Failed to assert connection as TCPConn")
	}

	// Example data transmission
	message := []byte("Hello from Golang! 🚀")
	if _, err := conn.Write(message); err != nil {
		fmt.Println("Failed to write message:", err)
	}

	// Example data reception
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))
}

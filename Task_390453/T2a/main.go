package main

import (
	"encoding/binary"
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

	// Set send and receive buffer sizes
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}

	// Set TCP send buffer size
	defer conn.Close()

	sendBufSize, err := strconv.Atoi(os.Getenv("TCP_SNDBUF"))
	if err != nil {
		sendBufSize = 65536
	}
	if err := conn.SetSockoptInt("tcp", syscall.SO_SNDBUF, sendBufSize); err != nil {
		fmt.Println("Failed to set send buffer size:", err)
	}

	// Set TCP receive buffer size
	recvBufSize, err := strconv.Atoi(os.Getenv("TCP_RCVBUF"))
	if err != nil {
		recvBufSize = 65536
	}
	if err := conn.SetSockoptInt("tcp", syscall.SO_RCVBUF, recvBufSize); err != nil {
		fmt.Println("Failed to set receive buffer size:", err)
	}

	// Set TCP keep-alive
	keepAliveInterval := 5 * time.Second // Interval after which the keep-alive messages start
	keepAliveCount := 5                 // Number of keep-alive messages to send

	if err := conn.SetKeepAlive(true); err != nil {
		fmt.Println("Failed to set keep-alive:", err)
	}
	if err := conn.SetKeepAlivePeriod(keepAliveInterval); err != nil {
		fmt.Println("Failed to set keep-alive interval:", err)
	}
	if err := conn.SetKeepAliveCount(keepAliveCount); err != nil {
		fmt.Println("Failed to set keep-alive count:", err)
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
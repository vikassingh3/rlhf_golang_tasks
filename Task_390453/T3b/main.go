package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}
	host := os.Args[1]
	port := os.Args[2]

	// Dial the TCP connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer conn.Close()

	// Assert the connection as a *net.TCPConn for advanced configurations
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Println("Failed to assert connection as TCPConn")
		return
	}

	// Set TCP send buffer size
	sendBufSize := 65536 // Adjust this size based on your requirements
	if err := tcpConn.SetWriteBuffer(sendBufSize); err != nil {
		fmt.Println("Failed to set send buffer size:", err)
	}

	// Set TCP receive buffer size
	recvBufSize := 65536
	if err := tcpConn.SetReadBuffer(recvBufSize); err != nil {
		fmt.Println("Failed to set receive buffer size:", err)
	}

	// Disable Nagle's algorithm
	if err := tcpConn.SetNoDelay(true); err != nil {
		fmt.Println("Failed to disable Nagle's algorithm:", err)
	}

	// Set TCP keep-alive
	keepAliveInterval := 5 * time.Second
	if err := tcpConn.SetKeepAlive(true); err != nil {
		fmt.Println("Failed to enable keep-alive:", err)
	}
	if err := tcpConn.SetKeepAlivePeriod(keepAliveInterval); err != nil {
		fmt.Println("Failed to set keep-alive period:", err)
	}

	// TCP fast open and initial handshake simulation (Write deadline)
	if err := tcpConn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
		fmt.Println("Failed to set write deadline:", err)
	}
	_, err = tcpConn.Write([]byte("\x00")) // Send initial data to simulate fast open
	if err != nil {
		fmt.Println("Failed to send initial data:", err)
		return
	}
	if err := tcpConn.SetWriteDeadline(time.Time{}); err != nil {
		fmt.Println("Failed to clear write deadline:", err)
	}

	// Example data transmission
	message := []byte("Hello from Golang! Optimized TCP Connection")
	if _, err := tcpConn.Write(message); err != nil {
		fmt.Println("Failed to write message:", err)
		return
	}

	// Example data reception
	buffer := make([]byte, 1024)
	n, err := tcpConn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))
}

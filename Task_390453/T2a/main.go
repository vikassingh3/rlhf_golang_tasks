package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"golang.org/x/sys/unix"
)

func main() {
	// Set TCP buffers
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	// Dial a TCP connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer conn.Close()

	// Get the file descriptor for the connection
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Println("Failed to get TCP connection")
		return
	}

	file, err := tcpConn.File()
	if err != nil {
		fmt.Println("Failed to get file descriptor:", err)
		return
	}
	defer file.Close()

	fd := int(file.Fd())

	// Set TCP send buffer size
	sendBufSize, err := strconv.Atoi(os.Getenv("TCP_SNDBUF"))
	if err != nil {
		sendBufSize = 65536
	}
	if err := unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_SNDBUF, sendBufSize); err != nil {
		fmt.Println("Failed to set send buffer size:", err)
	}

	// Set TCP receive buffer size
	recvBufSize, err := strconv.Atoi(os.Getenv("TCP_RCVBUF"))
	if err != nil {
		recvBufSize = 65536
	}
	if err := unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_RCVBUF, recvBufSize); err != nil {
		fmt.Println("Failed to set receive buffer size:", err)
	}

	// Enable TCP keep-alive
	keepAlive := 1
	if err := unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_KEEPALIVE, keepAlive); err != nil {
		fmt.Println("Failed to enable keep-alive:", err)
	}

	// Set keep-alive interval (TCP_KEEPINTVL)
	keepAliveInterval := 5 // in seconds
	if err := unix.SetsockoptInt(fd, unix.IPPROTO_TCP, unix.TCP_KEEPINTVL, keepAliveInterval); err != nil {
		fmt.Println("Failed to set keep-alive interval:", err)
	}

	// Set keep-alive probes (TCP_KEEPCNT)
	keepAliveCount := 5
	if err := unix.SetsockoptInt(fd, unix.IPPROTO_TCP, unix.TCP_KEEPCNT, keepAliveCount); err != nil {
		fmt.Println("Failed to set keep-alive probes:", err)
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

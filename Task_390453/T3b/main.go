
package main  
import (  
    "fmt"
    "net"
    "os"
    "syscall"
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

    //Set TCP Send Buffer Size
    sendBufSize := 65536 //Adjust this size based on your requirement
    if err := conn.SetWriteBuffer(sendBufSize); err != nil {
        fmt.Println("Failed to set send buffer size:", err)
    }
    //Set TCP Receive Buffer Size
    recvBufSize := 65536
    if err := conn.SetReadBuffer(recvBufSize); err != nil {
        fmt.Println("Failed to set receive buffer size:", err)
    }
    //Disable Nagle Algorithm
    if err := conn.SetNoDelay(true); err != nil {
        fmt.Println("Failed to disable Nagle algorithm:", err)
    }
    // Enable TCP Window Scaling
    if err := conn.SetWindowScale(14); err != nil {
        fmt.Println("Failed to enable TCP window scaling:", err)
    }
    // Set TCP Keep-Alive
    keepAliveInterval := 5 * time.Second
    if err := conn.SetKeepAlive(true); err != nil {
        fmt.Println("Failed to enable keep-alive:", err)
    }
    if err := conn.SetKeepAlivePeriod(keepAliveInterval); err != nil {
        fmt.Println("Failed to set keep-alive period:", err)
    }

   //Enable TCP Fast Open(If available)
    if err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
        fmt.Println("Failed to set write deadline:", err)
    }
    //Force TCP fast open handshake
    _, err = conn.Write([]byte("\x00"))
    if err != nil {
        fmt.Println("Failed to send initial data:", err)
        return
    }
    if err := conn.SetWriteDeadline(time.Time{}); err != nil {
        fmt.Println("Failed to clear write deadline:", err)
    }

    //Continue with your data transmission and reception logic
    message := []byte("Hello from Golang! Optimized TCP Connection")
    if _, err := conn.Write(message); err != nil {
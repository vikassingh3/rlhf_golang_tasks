package main  
import (  
    "fmt"
    "net"
)

// Define the maximum TCP packet size
const MaxTCPPacketSize = 65535  

func main() {  
    // Listen on a port for incoming TCP connections
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {  
        fmt.Println("Error listening:", err)
        return
    }  
    defer listener.Close()

    for {  
        // Accept incoming connections
        conn, err := listener.Accept()
        if err != nil {  
            fmt.Println("Error accepting connection:", err)
            continue
        }  
        defer conn.Close()

        // Handle the connection in a separate goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {  
    // Create a buffer to read TCP data
    buf := make([]byte, MaxTCPPacketSize)

    // Read data from the connection
    for {  
        n, err := conn.Read(buf)
        if err != nil {  
            fmt.Println("Error reading data:", err)
            return
        }

        // Handle the received data
        handleData(buf[:n])
    }
}

func handleData(data []byte) {  
    // In a real implementation, you would need to handle TCP segmentation and reassembly here.
    // For simplicity, let's just print the received data as a string.
    fmt.Println(string(data))
}
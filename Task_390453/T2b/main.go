package main  
import (  
    "fmt"
    "net"
    "os"
)

func main() {
    // Create a TCP listener with a specific TCP options
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8080})
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }
    defer listener.Close()

    // Set the maximum segment size (MSS) to the MTU of the path
    err = listener.SetReadBuffer(MTU)
    if err != nil {
        fmt.Println("Error setting read buffer:", err)
        os.Exit(1)
    }

    // Set the TCP keep-alive interval
    err = listener.SetKeepAlive(true)
    if err != nil {
        fmt.Println("Error setting keep-alive:", err)
        os.Exit(1)
    }
    err = listener.SetKeepAlivePeriod(3 * time.Minute)
    if err != nil {
        fmt.Println("Error setting keep-alive period:", err)
        os.Exit(1)
    }

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
}
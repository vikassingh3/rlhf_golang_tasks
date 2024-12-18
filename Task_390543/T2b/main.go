package main  
import (  
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
)

const (
    writeWait       = 10 * time.Second
    pongWait        = 60 * time.Second
    pingPeriod      = (pongWait * 9) / 10
    wsURL           = "ws://echo.websocket.org" // Using the public WebSocket echo server for testing
)

var upg = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// measureLatency returns the measured latency in milliseconds
func measureLatency(ctx context.Context, url string) (latency int, err error) {
    c, _, err := upg.Upgrade(ctx, &http.Request{URL: &http.URL{Scheme: "ws", Host: url}}, nil)
    if err != nil {
        return 0, fmt.Errorf("failed to connect: %w", err)
    }
    defer c.Close()

    // Initial ping to ensure connection
    if err = c.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait)); err != nil {
        return 0, fmt.Errorf("ping failed: %w", err)
    }

    // Write message and measure latency
    start := time.Now()
    if err = c.WriteMessage(websocket.TextMessage, []byte("measure-latency")); err != nil {
        return 0, fmt.Errorf("write message failed: %w", err)
    }

    _, _, err = c.ReadMessage()
    if err != nil {
        return 0, fmt.Errorf("read message failed: %w", err)
    }

    return int(time.Since(start).Milliseconds()), nil
} 

func main() {
    //  Set up a 5 second timeout context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    latency, err := measureLatency(ctx, wsURL)
    if err != nil {
        log.Printf("Error during WebSocket communication: %v", err)
        return
    }

    log.Printf("Latency: %d ms\n", latency)
} 
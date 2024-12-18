
package main  
import (  
    "context"
    "fmt"
    "log"
    "net/http"
    "sync"
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

type wsClient struct {
    conn   *websocket.Conn
    wg     sync.WaitGroup
    cancel context.CancelFunc
}

func (c *wsClient) start(ctx context.Context) {
    c.wg.Add(1)
    defer c.wg.Done()

    ticker := time.NewTicker(pingPeriod)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            c.writeControl(websocket.PingMessage, nil)
        default:
            mt, message, err := c.conn.ReadMessage()
            if err != nil {
                log.Printf("Read error: %v", err)
                return
            }
            log.Printf("Received message: %s", message)

            // Handle specific messages or perform other actions based on message type (mt)
        }
    }
}

func (c *wsClient) writeControl(mt int, payload []byte) {
    c.conn.SetWriteDeadline(time.Now().Add(writeWait))
    if err := c.conn.WriteControl(mt, payload, time.Now().Add(writeWait)); err != nil {
        log.Printf("Write control error: %v", err)
    }
}

func (c *wsClient) writeMessage(mt int, payload []byte) {
    c.conn.SetWriteDeadline(time.Now().Add(writeWait))
    if err := c.conn.WriteMessage(mt, payload); err != nil {
        log.Printf("Write message error: %v", err)
    }
}

func (c *wsClient) close() {
    c.cancel()
    c.conn.Close()
    c.wg.Wait()
}

func measureLatency() (latency int, err error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    client := &wsClient{}
    defer client.close()

    // Connect to WebSocket server
    conn, _, err := upg.Upgrade(ctx, &http.Request{URL: &http.URL{Scheme: "ws", Host: wsURL}}, nil)
    if err != nil {
        return 0, fmt.Errorf("failed to connect: %w", err)
    }
    client.conn = conn
    client.cancel = cancel

    go client.start(ctx)

    // Send latency measurement message
    start := time.Now()
    client.writeMessage(websocket.TextMessage, []byte("measure-latency"))

    // Wait for response or timeout
    select {
    case <-ctx.Done():
        return 0, fmt.Errorf("timeout during latency measurement")
    default:
    }

    // Measure latency
    return int(time.Since(start).Milliseconds()), nil
}

func main() {
    latency, err := measureLatency()
    if err != nil {
        log.Printf("Error during WebSocket communication: %v", err)
        return
    }
    log.Printf("Latency: %d ms\n", latency)
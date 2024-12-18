

package main  
import (  
   "fmt"
   "log"
   "net/http"
   "time"
   "github.com/gorilla/websocket"
)

func calculateLatency(url string) {
   var latency time.Duration
   c, _, err := websocket.DefaultDialer.Dial(url, nil)
   if err != nil {
       log.Fatal("dial:", err)
   }
   defer c.Close()

   pingMessage := "ping"

   // send ping message
   err = c.WriteMessage(websocket.TextMessage, []byte(pingMessage))
   if err != nil {
       log.Println("write:", err)
       return
   }

   // Measure the time taken for the pong response
   start := time.Now()
   _, message, err := c.ReadMessage()
   if err != nil {
       log.Println("read:", err)
       return
   }
   end := time.Now()

   latency = end.Sub(start)
   fmt.Printf("Latency to server: %s\n", latency)
}

func main() {
   // Provide the WebSocket server URL
   url := "ws://localhost:8080/ws"

   calculateLatency(url)
}


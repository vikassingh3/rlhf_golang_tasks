package main  
import (  
    "container/heap"
    "fmt"
    "time"
)  

// Message represents a message with urgency and timestamp.  
type Message struct {  
    Urgency  int        // Higher urgency means higher priority.
    Timestamp time.Time // Timestamp for priority if urgency is the same.
    Content   string     // Actual message content.
}  

// MessageHeap implements heap.Interface and holds Messages.  
type MessageHeap []Message  

func (h MessageHeap) Len() int           { return len(h) }  
func (h MessageHeap) Less(i, j int) bool { return ByUrgency(h).Less(i, j) }  
func (h MessageHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }  

func (h *MessageHeap) Push(x interface{}) {  
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(Message))  
}  

func (h *MessageHeap) Pop() interface{} {  
    old := *h  
    n := len(old)  
    x := old[n-1]  
    *h = old[0 : n-1]  
    return x  
}  
  
// ByUrgency implements the sort.Interface for []Message based on the Urgency field.  
type ByUrgency []Message  

func (a ByUrgency) Len() int           { return len(a) }  
func (a ByUrgency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }  
func (a ByUrgency) Less(i, j int) bool {  
    if a[i].Urgency != a[j].Urgency {  
        return a[i].Urgency > a[j].Urgency  
    }  
    return a[i].Timestamp.Before(a[j].Timestamp)  
}  

func main() {  
    // Create a message heap and initialize it with some messages.  
    messages := MessageHeap{}  
    heap.Init(&messages)  
    heap.Push(&messages, Message{Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"})  
    heap.Push(&messages, Message{Urgency: 2, Timestamp: time.Now(), Content: "Message 2"})  
    heap.Push(&messages, Message{Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"})  
    heap.Push(&messages, Message{Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"})  
    heap.Push(&messages, Message{Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"})  
  
    fmt.Println("Initial heap:")  
    for _, msg := range messages {  
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)  
    }  
  
    // Add a new message with higher urgency to simulate new message arrival.  
    newMessage := Message{Urgency: 5, Timestamp: time.Now(), Content: "New urgent message!"}  
    heap.Push(&messages, newMessage)  
  
    fmt.Println("\nHeap after adding new urgent message:")  
    for _, msg := range messages {  
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)  
    }  
  
    // Remove the highest priority message from the heap.  
    highestPriorityMessage := heap.Pop(&messages).(Message)  
    fmt.Println("\nHighest priority message removed:")  
    fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", highestPriorityMessage.Urgency, highestPriorityMessage.Timestamp.Format("15:04:05"), highestPriorityMessage.Content)  
  
    fmt.Println("\nHeap after removing highest priority message:")  
    for _, msg := range messages {  
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)  
    }  
}  
 
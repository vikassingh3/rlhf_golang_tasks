package main

import (
    "fmt"
    "sort"
    "time"
)

// Message represents a message with urgency and timestamp.
type Message struct {
    Urgency  int        // Higher urgency means higher priority.
    Timestamp time.Time // Timestamp for priority if urgency is the same.
    Content   string     // Actual message content.
}

// ByUrgency implements the sort.Interface for []Message based on the Urgency field.
type ByUrgency []Message

func (a ByUrgency) Len() int           { return len(a) }
func (a ByUrgency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUrgency) Less(i, j int) bool {
    // Sort first by urgency in descending order (higher urgency first).
    if a[i].Urgency != a[j].Urgency {
        return a[i].Urgency > a[j].Urgency
    }

    // If urgency is the same, sort by timestamp in ascending order (oldest first).
    return a[i].Timestamp.Before(a[j].Timestamp)
}

func main() {
    // Initialize an unordered slice of messages.
    messages := []Message{
        {Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"},
        {Urgency: 2, Timestamp: time.Now(), Content: "Message 2"},
        {Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"},
        {Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"},
        {Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"},
    }

    fmt.Println("Initial unsorted messages:")
    for _, msg := range messages {
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
    }
    
    // Sort the messages based on urgency and timestamp.
    sort.Sort(ByUrgency(messages))

    fmt.Println("\nSorted messages by urgency and timestamp:")
    for _, msg := range messages {
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
    }

    // Get the highest priority message (first element after sorting).
    highestPriorityMessage := messages[0]
    fmt.Println("\nHighest priority message:")
    fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", highestPriorityMessage.Urgency, highestPriorityMessage.Timestamp.Format("15:04:05"), highestPriorityMessage.Content)

    // Remove the highest priority message from the slice.
    messages = messages[1:]

    // Add a new message with higher urgency to simulate new message arrival.
    newMessage := Message{Urgency: 5, Timestamp: time.Now(), Content: "New urgent message!"}
    messages = append([]Message{newMessage}, messages...)

    fmt.Println("\nMessages after adding new urgent message:")
    for _, msg := range messages {
        fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
    }
}
package main

import (
	"fmt"
	"sort"
	"time"
)

// Define the Message struct
type Message struct {
	Content string
	Urgency int // Lower value indicates higher urgency
	Timestamp time.Time
}

// Custom sorting interface for Messages
func (m Message) Less(than Message) bool {
	// Check urgency first
	if m.Urgency < than.Urgency {
		return true
	} else if m.Urgency > than.Urgency {
		return false
	}
	// If urgency is the same, compare timestamps
	return m.Timestamp.After(than.Timestamp)
}

func main() {
	// Create a slice of messages
	messages := []Message{
		{"urgent notification", 1, time.Now().Add(time.Second * 10)},
		{"normal message", 2, time.Now().Add(time.Second * 5)},
		{"low priority", 3, time.Now()},
	}

	// Print original order
	fmt.Println("Original messages:")
	for _, message := range messages {
		fmt.Printf("Content: %s, Urgency: %d, Timestamp: %v\n", message.Content, message.Urgency, message.Timestamp)
	}

	// Sort messages by urgency (lower is more urgent) and timestamp
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Less(messages[j])
	})

	// Print sorted order
	fmt.Println("\nSorted messages:")
	for _, message := range messages {
		fmt.Printf("Content: %s, Urgency: %d, Timestamp: %v\n", message.Content, message.Urgency, message.Timestamp)
	}
}
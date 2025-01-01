package main

import (
	"container/heap"
	"fmt"
	"time"
)

// Message represents a message with urgency and timestamp.
type Message struct {
	Urgency   int       // Higher urgency means higher priority.
	Timestamp time.Time // Timestamp for priority if urgency is the same.
	Content   string    // Actual message content.
	index     int       // Index in the heap (used for efficient removal).
}

// MessageHeap implements heap.Interface and holds Messages.
type MessageHeap []*Message

func (h MessageHeap) Len() int { return len(h) }
func (h MessageHeap) Less(i, j int) bool {
	// Higher urgency comes first. If urgency is the same, older messages come first.
	if h[i].Urgency != h[j].Urgency {
		return h[i].Urgency > h[j].Urgency
	}
	return h[i].Timestamp.Before(h[j].Timestamp)
}
func (h MessageHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MessageHeap) Push(x interface{}) {
	n := len(*h)
	message := x.(*Message)
	message.index = n
	*h = append(*h, message)
}

func (h *MessageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	message := old[n-1]
	message.index = -1 // For safety
	*h = old[0 : n-1]
	return message
}

// PriorityQueue manages messages with a priority queue and supports efficient removal.
type PriorityQueue struct {
	messages MessageHeap
	lookup   map[string]*Message // Map content to message for fast removal
}

// NewPriorityQueue creates a new PriorityQueue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		messages: make(MessageHeap, 0),
		lookup:   make(map[string]*Message),
	}
}

// AddMessage adds a message to the queue.
func (pq *PriorityQueue) AddMessage(message *Message) {
	heap.Push(&pq.messages, message)
	pq.lookup[message.Content] = message
}

// RemoveMessage removes a message by content and maintains heap order.
func (pq *PriorityQueue) RemoveMessage(content string) {
	message, exists := pq.lookup[content]
	if !exists {
		return
	}
	// Remove from the heap
	heap.Remove(&pq.messages, message.index)
	delete(pq.lookup, content)
}

// PopMessage removes and returns the highest-priority message.
func (pq *PriorityQueue) PopMessage() *Message {
	if len(pq.messages) == 0 {
		return nil
	}
	message := heap.Pop(&pq.messages).(*Message)
	delete(pq.lookup, message.Content)
	return message
}

// Len returns the number of messages.
func (pq *PriorityQueue) Len() int {
	return len(pq.messages)
}

func main() {
	// Create a priority queue and add messages.
	pq := NewPriorityQueue()

	pq.AddMessage(&Message{Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"})
	pq.AddMessage(&Message{Urgency: 2, Timestamp: time.Now(), Content: "Message 2"})
	pq.AddMessage(&Message{Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"})
	pq.AddMessage(&Message{Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"})
	pq.AddMessage(&Message{Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"})

	// Remove a specific message.
	pq.RemoveMessage("Message 4")

	// Print and remove messages in priority order.
	fmt.Println("Messages in priority order:")
	for pq.Len() > 0 {
		msg := pq.PopMessage()
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
	}
}

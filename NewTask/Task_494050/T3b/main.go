package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Message represents a message with urgency and timestamp.
type Message struct {
	Urgency   int       // Higher urgency means higher priority.
	Timestamp time.Time // For priority when urgency is equal.
	Content   string    // Actual message content.
}

// MessageHeap implements heap.Interface and holds Messages.
type MessageHeap []*Message

func (h MessageHeap) Len() int { return len(h) }

func (h MessageHeap) Less(i, j int) bool {
	// Higher urgency comes first. If urgency is equal, older timestamp comes first.
	if h[i].Urgency != h[j].Urgency {
		return h[i].Urgency > h[j].Urgency
	}
	return h[i].Timestamp.Before(h[j].Timestamp)
}

func (h MessageHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MessageHeap) Push(x interface{}) {
	*h = append(*h, x.(*Message))
}

func (h *MessageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	message := old[n-1]
	*h = old[0 : n-1]
	return message
}

// PriorityQueue manages messages with a priority queue.
type PriorityQueue struct {
	mu       sync.Mutex // Ensures thread-safe access.
	messages MessageHeap
}

// NewPriorityQueue creates a new PriorityQueue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		messages: make(MessageHeap, 0),
	}
}

// AddMessage adds a message to the queue.
func (pq *PriorityQueue) AddMessage(message *Message) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	heap.Push(&pq.messages, message)
}

// PopMessage removes and returns the highest-priority message.
func (pq *PriorityQueue) PopMessage() *Message {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.messages.Len() == 0 {
		return nil // Return nil if the queue is empty.
	}

	return heap.Pop(&pq.messages).(*Message)
}

// Len returns the number of messages in the queue.
func (pq *PriorityQueue) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.messages.Len()
}

func main() {
	// Create a priority queue and add messages.
	pq := NewPriorityQueue()

	pq.AddMessage(&Message{Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"})
	pq.AddMessage(&Message{Urgency: 2, Timestamp: time.Now(), Content: "Message 2"})
	pq.AddMessage(&Message{Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"})
	pq.AddMessage(&Message{Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"})
	pq.AddMessage(&Message{Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"})

	// Print and remove messages in priority order.
	fmt.Println("Messages in priority order:")
	for pq.Len() > 0 {
		msg := pq.PopMessage()
		if msg != nil {
			fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n",
				msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
		}
	}
}

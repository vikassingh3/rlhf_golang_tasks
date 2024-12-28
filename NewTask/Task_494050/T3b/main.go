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
	Content   string    // Message content.
}

// MessageQueue implements a priority queue for Message.
type MessageQueue []*Message

// Len returns the length of the queue.
func (mq MessageQueue) Len() int { return len(mq) }

// Less determines the order of elements in the priority queue.
// Higher urgency first. If urgency is the same, earlier timestamp first.
func (mq MessageQueue) Less(i, j int) bool {
	if mq[i].Urgency == mq[j].Urgency {
		return mq[i].Timestamp.Before(mq[j].Timestamp)
	}
	return mq[i].Urgency > mq[j].Urgency
}

// Swap swaps two elements in the queue.
func (mq MessageQueue) Swap(i, j int) { mq[i], mq[j] = mq[j], mq[i] }

// Push adds an element to the queue.
func (mq *MessageQueue) Push(x interface{}) {
	*mq = append(*mq, x.(*Message))
}

// Pop removes and returns the highest-priority element from the queue.
func (mq *MessageQueue) Pop() interface{} {
	old := *mq
	n := len(old)
	item := old[n-1]
	*mq = old[:n-1]
	return item
}

func main() {
	// Create a priority queue and initialize it.
	pq := &MessageQueue{}
	heap.Init(pq)

	// Insert messages into the priority queue.
	heap.Push(pq, &Message{Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"})
	heap.Push(pq, &Message{Urgency: 2, Timestamp: time.Now(), Content: "Message 2"})
	heap.Push(pq, &Message{Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"})
	heap.Push(pq, &Message{Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"})
	heap.Push(pq, &Message{Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"})

	// Print and remove messages in priority order.
	fmt.Println("Messages in priority order:")
	for pq.Len() > 0 {
		msg := heap.Pop(pq).(*Message)
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
	}
}

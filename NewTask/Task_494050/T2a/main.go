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
}

// MessageQueue is a priority queue of Messages.
type MessageQueue []*Message

// Len returns the length of the queue.
func (mq MessageQueue) Len() int { return len(mq) }

// Less compares two messages based on urgency and timestamp.
func (mq MessageQueue) Less(i, j int) bool {
	// Higher urgency first. If urgency is the same, earlier timestamp first.
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

	// Initialize some messages.
	messages := []*Message{
		{Urgency: 3, Timestamp: time.Now().Add(-5 * time.Minute), Content: "Message 3"},
		{Urgency: 2, Timestamp: time.Now(), Content: "Message 2"},
		{Urgency: 4, Timestamp: time.Now().Add(-2 * time.Minute), Content: "Message 4"},
		{Urgency: 1, Timestamp: time.Now().Add(-3 * time.Minute), Content: "Message 1"},
		{Urgency: 3, Timestamp: time.Now().Add(-1 * time.Minute), Content: "Message 3B"},
	}

	// Push messages onto the priority queue.
	for _, msg := range messages {
		heap.Push(pq, msg)
	}

	// Print the highest-priority message.
	fmt.Println("\nCurrent highest-priority message:")
	if pq.Len() > 0 {
		highestPriorityMessage := heap.Pop(pq).(*Message)
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", highestPriorityMessage.Urgency, highestPriorityMessage.Timestamp.Format("15:04:05"), highestPriorityMessage.Content)
	}

	// Add a new message with higher urgency.
	newMessage := &Message{Urgency: 5, Timestamp: time.Now(), Content: "New urgent message!"}
	heap.Push(pq, newMessage)

	// Process the top-priority message.
	fmt.Println("\nNew highest-priority message after addition:")
	if pq.Len() > 0 {
		highestPriorityMessage := heap.Pop(pq).(*Message)
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", highestPriorityMessage.Urgency, highestPriorityMessage.Timestamp.Format("15:04:05"), highestPriorityMessage.Content)
	}

	// Print all remaining messages.
	fmt.Println("\nRemaining messages in priority queue:")
	for pq.Len() > 0 {
		msg := heap.Pop(pq).(*Message)
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
	}
}

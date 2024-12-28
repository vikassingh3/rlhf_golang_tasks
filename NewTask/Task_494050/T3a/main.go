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
	Timestamp time.Time // Timestamp for priority if urgency is the same.
	Content   string    // Actual message content.
}

// MessageHeap implements heap.Interface and holds Messages.
type MessageHeap []Message

func (h MessageHeap) Len() int { return len(h) }
func (h MessageHeap) Less(i, j int) bool {
	// Higher urgency comes first. If urgency is the same, older messages come first.
	if h[i].Urgency != h[j].Urgency {
		return h[i].Urgency > h[j].Urgency
	}
	return h[i].Timestamp.Before(h[j].Timestamp)
}
func (h MessageHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MessageHeap) Push(x interface{}) {
	*h = append(*h, x.(Message))
}

func (h *MessageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Create a priority queue and initialize it with some messages.
	messages := &MessageHeap{}
	heap.Init(messages)

	// Concurrently add messages
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			heap.Push(messages, Message{Urgency: 3, Timestamp: time.Now(), Content: fmt.Sprintf("Message 3%d", i)})
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			heap.Push(messages, Message{Urgency: 2, Timestamp: time.Now(), Content: fmt.Sprintf("Message 2%d", i)})
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			heap.Push(messages, Message{Urgency: 4, Timestamp: time.Now(), Content: fmt.Sprintf("Message 4%d", i)})
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			heap.Push(messages, Message{Urgency: 1, Timestamp: time.Now(), Content: fmt.Sprintf("Message 1%d", i)})
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			heap.Push(messages, Message{Urgency: 3, Timestamp: time.Now(), Content: fmt.Sprintf("Message 3B%d", i)})
		}
	}()

	wg.Wait()

	// Process messages in priority order.
	fmt.Println("Messages in priority order:")
	for messages.Len() > 0 {
		msg := heap.Pop(messages).(Message)
		fmt.Printf("[Urgency: %d, Timestamp: %s] %s\n", msg.Urgency, msg.Timestamp.Format("15:04:05"), msg.Content)
	}
}
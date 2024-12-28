package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Action string
}

type Counter struct {
	value int
	mu    *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		mu:    &sync.Mutex{},
	}
}

func (c *Counter) Actor(in <-chan Message) {
	for msg := range in {
		c.mu.Lock()
		// Unlock immediately after processing the message
		switch msg.Action {
		case "increment":
			c.value++
		default:
			fmt.Println("Unknown action:", msg.Action)
		}
		c.mu.Unlock() // Unlock here, not deferred
	}
}

func main() {
	counter := NewCounter()
	in := make(chan Message)

	// Start the actor
	go counter.Actor(in)

	// Simulate multiple concurrent increment requests
	for i := 0; i < 10000; i++ {
		in <- Message{Action: "increment"}
	}

	// Close the input channel to signal the actor to stop
	close(in)

	// Wait for the actor to finish processing all messages
	time.Sleep(1 * time.Second)

	fmt.Println("Final counter value:", counter.value)
}

package main

import (
	"fmt"
	"sync"
)

// Example: Simple Counter Actor
type CounterActor struct {
	count int
	ch    chan int // Channel to receive messages
}

func NewCounterActor() *CounterActor {
	return &CounterActor{
		ch: make(chan int),
	}
}

func (a *CounterActor) Start() {
	// This is the main actor loop that continuously processes messages
	for msg := range a.ch {
		a.count += msg
		fmt.Printf("Current count: %d\n", a.count)
	}
}

// Send a message to the actor
func (a *CounterActor) Send(msg int) {
	a.ch <- msg
}

func main() {
	// Create an actor
	counterActor := NewCounterActor()

	// Start the actor's message processing loop in a goroutine
	go counterActor.Start()

	// Multiple goroutines can send messages to the actor concurrently
	var wg sync.WaitGroup
	numGoRoutines := 10
	incrementAmount := 10

	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			counterActor.Send(incrementAmount)
			wg.Done()
		}()
	}

	wg.Wait()

	// Close the channel to terminate the actor's loop
	close(counterActor.ch)

	fmt.Println("Final Count: ", counterActor.count)
}

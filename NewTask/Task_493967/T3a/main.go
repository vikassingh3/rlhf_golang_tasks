package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Simple Counter Actor with timeout handling
type CounterActor struct {
	count int
	ch    chan int
	wg    *sync.WaitGroup
	ctx   context.Context
	cancel context.CancelFunc
}

func NewCounterActor(wg *sync.WaitGroup) *CounterActor {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return &CounterActor{
		ch:    make(chan int, 10), // Buffer size to prevent blocking
		wg:    wg,
		ctx:   ctx,
		cancel: cancel,
	}
}

func (a *CounterActor) Start() {
	defer a.wg.Done()
	for {
		select {
		case msg, ok := <-a.ch:
			if !ok {
				fmt.Println("Channel closed, shutting down.")
				return
			}
			a.count += msg
			fmt.Printf("Current count: %d\n", a.count)
			time.Sleep(time.Duration(msg) * time.Millisecond)
		case <-a.ctx.Done():
			fmt.Println("Actor timed out, shutting down.")
			return
		}
	}
}

func (a *CounterActor) Send(msg int) {
	select {
	case a.ch <- msg:
	default:
		fmt.Println("Channel buffer full, skipping message.")
	}
}

func main() {
	var wg sync.WaitGroup

	// Create multiple actors
	numActors := 3
	actors := make([]*CounterActor, numActors)
	for i := 0; i < numActors; i++ {
		actors[i] = NewCounterActor(&wg)
	}

	// Start all actors
	wg.Add(numActors)
	for _, actor := range actors {
		go actor.Start()
	}

	// Send messages to the actors
	messagesToSend := 10
	for _, actor := range actors {
		for i := 0; i < messagesToSend; i++ {
			actor.Send(i + 1)
		}
	}

	// Close channels to signal actors to stop
	for _, actor := range actors {
		close(actor.ch)
	}

	// Wait for all actors to finish
	wg.Wait()
	fmt.Println("All actors finished.")
}
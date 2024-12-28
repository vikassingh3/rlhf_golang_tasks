package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// CounterActor with timeout handling
type CounterActor struct {
	count  int
	ch     chan int
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func NewCounterActor(wg *sync.WaitGroup) *CounterActor {
	ctx, cancel := context.WithCancel(context.Background())
	return &CounterActor{
		ch:     make(chan int, 10),
		ctx:    ctx,
		cancel: cancel,
		wg:     wg,
	}
}

func (a *CounterActor) Start() {
	defer a.wg.Done()
	for {
		select {
		case msg, ok := <-a.ch:
			if !ok {
				fmt.Println("Channel closed, actor shutting down.")
				return
			}
			a.count += msg
			fmt.Printf("Processed message: %d, Current count: %d\n", msg, a.count)
		case <-a.ctx.Done():
			fmt.Println("Context canceled, actor shutting down.")
			return
		}
	}
}

func (a *CounterActor) Send(msg int) {
	select {
	case a.ch <- msg:
	default:
		fmt.Println("Channel full, message skipped.")
	}
}

func (a *CounterActor) Stop() {
	a.cancel()
	close(a.ch)
}

func main() {
	var wg sync.WaitGroup

	// Create and start an actor
	actor := NewCounterActor(&wg)
	wg.Add(1)
	go actor.Start()

	// Send messages
	for i := 1; i <= 15; i++ {
		actor.Send(i)
		time.Sleep(100 * time.Millisecond)
	}

	// Stop the actor
	actor.Stop()

	// Wait for actor to finish
	wg.Wait()
	fmt.Println("All actors completed.")
}

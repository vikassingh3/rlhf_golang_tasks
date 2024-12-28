package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxMessageWaitTime = 2 * time.Second // Maximum time to wait for a message response
)

// Simple Processor Actor
type ProcessorActor struct {
	ch       chan int
	wg       *sync.WaitGroup
	response chan int
}

func NewProcessorActor(wg *sync.WaitGroup) *ProcessorActor {
	return &ProcessorActor{
		ch:       make(chan int),
		wg:       wg,
		response: make(chan int),
	}
}

func (p *ProcessorActor) Start() {
	defer p.wg.Done()
	for msg := range p.ch {
		fmt.Printf("Processing message: %d\n", msg)
		// Simulate processing time
		time.Sleep(time.Duration(msg) * time.Millisecond)

		// Send response
		p.response <- msg
	}
}

func (p *ProcessorActor) Send(msg int) {
	p.ch <- msg
}

func (p *ProcessorActor) ReceiveResponse() (int, bool) {
	select {
	case response := <-p.response:
		return response, true // Received response
	case <-time.After(maxMessageWaitTime):
		return 0, false // Timed out
	}
}

// Simple Sender Actor
type SenderActor struct {
	processor *ProcessorActor
}

func NewSenderActor(processor *ProcessorActor) *SenderActor {
	return &SenderActor{
		processor: processor,
	}
}

func (s *SenderActor) SendMessage(msg int) {
	s.processor.Send(msg)
}

func (s *SenderActor) ReceiveResponse(msg int) bool {
	response, ok := s.processor.ReceiveResponse()
	if ok {
		fmt.Printf("Received response for message %d: %d\n", msg, response)
	} else {
		fmt.Printf("Timed out waiting for response to message %d\n", msg)
	}
	return ok
}

func main() {
	var wg sync.WaitGroup

	// Create actors
	processor := NewProcessorActor(&wg)
	sender := NewSenderActor(processor)

	// Start actors
	wg.Add(1)
	go processor.Start()

	// Send messages and receive responses
	messages := []int{100, 200, 300, 400, 500}
	for _, msg := range messages {
		sender.SendMessage(msg)
		sender.ReceiveResponse(msg)
	}

	// Close the processor channel to stop the actor
	close(processor.ch)

	// Wait for the processor to finish
	wg.Wait()

	fmt.Println("All messages processed.")
}

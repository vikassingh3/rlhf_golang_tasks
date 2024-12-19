package main

import (
	"fmt"
	"sync"
	"time"
)

// Custom error types for message processing
type MessageError struct {
	Step  string
	Errors error
}

func (e *MessageError) Error() string {
	return fmt.Sprintf("Message processing failed at step '%s': %s", e.Step, e.Error)
}

// State machine for message processing
type MessageProcessor struct {
	CurrentStep string
	mux         *sync.Mutex
}

func (mp *MessageProcessor) transition(step string, err error) error {
	if err != nil {
		return &MessageError{Step: step, Errors: err}
	}
	mp.mux.Lock()
	mp.CurrentStep = step
	mp.mux.Unlock()
	return nil
}

func (mp *MessageProcessor) ReceiveMessage() error {
	return mp.transition("receive", nil)
}

func (mp *MessageProcessor) ValidateMessage(msg string) error {
	if msg == "" {
		return &MessageError{Step: "validate", Errors: fmt.Errorf("message is empty")}
	}
	return mp.transition("validate", nil)
}

func (mp *MessageProcessor) SendResponse(resp string) error {
	return mp.transition("send_response", nil)
}

func (mp *MessageProcessor) MarkComplete() error {
	return mp.transition("mark_complete", nil)
}

func main() {
	mp := MessageProcessor{CurrentStep: "", mux: &sync.Mutex{}}

	err := mp.ReceiveMessage()
	if err != nil {
		fmt.Println("Error receiving message:", err)
	}

	err = mp.ValidateMessage("")
	if err != nil {
		fmt.Println("Error validating message:", err)
	}

	err = mp.SendResponse("OK")
	if err != nil {
		fmt.Println("Error sending response:", err)
	}

	err = mp.MarkComplete()
	if err != nil {
		fmt.Println("Error marking complete:", err)
	}

	time.Sleep(1 * time.Second)
}
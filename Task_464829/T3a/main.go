package main

import (
	"fmt"
	"time"
)

// NetworkStatus represents the different network states
type NetworkStatus string

const (
	Connected    NetworkStatus = "Connected"
	Disconnected NetworkStatus = "Disconnected"
	Error        NetworkStatus = "Error"
)

// StateMachine represents the state machine to manage network statuses
type StateMachine struct {
	currentStatus NetworkStatus
}

// NewStateMachine initializes a new state machine with the default status Disconnected
func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		currentStatus: Disconnected,
	}
	return sm
}

// Transition transitions the state machine to a new status based on the current status
func (sm *StateMachine) Transition(newStatus NetworkStatus) {
	switch sm.currentStatus {
	case Connected:
		switch newStatus {
		case Disconnected:
			sm.currentStatus = Disconnected
			fmt.Printf("Transitioned from %s to %s\n", Connected, Disconnected)
		case Error:
			sm.currentStatus = Error
			fmt.Printf("Transitioned from %s to %s\n", Connected, Error)
		}
	case Disconnected:
		switch newStatus {
		case Connected:
			sm.currentStatus = Connected
			fmt.Printf("Transitioned from %s to %s\n", Disconnected, Connected)
		case Error:
			sm.currentStatus = Error
			fmt.Printf("Transitioned from %s to %s\n", Disconnected, Error)
		}
	case Error:
		switch newStatus {
		case Disconnected:
			sm.currentStatus = Disconnected
			fmt.Printf("Transitioned from %s to %s\n", Error, Disconnected)
		case Connected:
			sm.currentStatus = Connected
			fmt.Printf("Transitioned from %s to %s\n", Error, Connected)
		}
	}
}

func main() {
	sm := NewStateMachine()

	// Start a routine to simulate network status changes
	go func() {
		sm.Transition(Connected)
		time.Sleep(2 * time.Second)
		sm.Transition(Error)
		time.Sleep(1 * time.Second)
		sm.Transition(Disconnected)
		time.Sleep(3 * time.Second)
		sm.Transition(Connected)
	}()

	// Main loop to demonstrate state transitions
	for {
		fmt.Printf("Current status: %s\n", sm.currentStatus)
		time.Sleep(1 * time.Second)
	}
}
package main

import (
	"fmt"
	"sync"
	"time"
)

// NetworkStatus is an enumeration of the different network states
type NetworkStatus int

const (
	Connected NetworkStatus = iota
	Disconnected
	Error
)

// StateMachine represents the state machine to manage network statuses
type StateMachine struct {
	current NetworkStatus
	cond    *sync.Cond
	done    chan struct{} // Channel to signal termination
}

// NewStateMachine initializes a new state machine with the default status Disconnected
func NewStateMachine() *StateMachine {
	return &StateMachine{
		current: Disconnected,
		cond:    sync.NewCond(&sync.Mutex{}),
		done:    make(chan struct{}), // To handle graceful shutdown
	}
}

// SetStatus transitions the state machine to a new status
func (sm *StateMachine) SetStatus(newStatus NetworkStatus) {
	sm.cond.L.Lock()
	defer sm.cond.L.Unlock()

	if sm.current == newStatus {
		return
	}

	sm.current = newStatus
	fmt.Printf("Network status changed to: %s\n", statusToString(sm.current))
	sm.cond.Broadcast()
}

// GetStatus returns the current network status
func (sm *StateMachine) GetStatus() NetworkStatus {
	sm.cond.L.Lock()
	defer sm.cond.L.Unlock()

	return sm.current
}

// WaitForStatus waits for the network status to become the specified status
func (sm *StateMachine) WaitForStatus(status NetworkStatus) {
	sm.cond.L.Lock()
	defer sm.cond.L.Unlock()

	for sm.current != status {
		select {
		case <-sm.done: // Exit if termination is signaled
			fmt.Println("Terminating wait for status...")
			return
		default:
			sm.cond.Wait() // Wait for the condition to change
		}
	}
}

// Stop gracefully shuts down the state machine
func (sm *StateMachine) Stop() {
	close(sm.done)      // Close the channel to signal termination
	sm.cond.Broadcast() // Wake all waiting goroutines
}

// statusToString converts a NetworkStatus to a human-readable string
func statusToString(status NetworkStatus) string {
	switch status {
	case Connected:
		return "Connected"
	case Disconnected:
		return "Disconnected"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

func main() {
	sm := NewStateMachine()

	// Start a routine to simulate network status changes
	go func() {
		time.Sleep(1 * time.Second)
		sm.SetStatus(Connected)
		time.Sleep(2 * time.Second)
		sm.SetStatus(Error)
		time.Sleep(1 * time.Second)
		sm.SetStatus(Disconnected)
		time.Sleep(3 * time.Second)
		sm.SetStatus(Connected)
	}()

	// Main loop to demonstrate wait functionality
	go func() {
		for {
			sm.WaitForStatus(Connected)
			sm.WaitForStatus(Error)
			sm.WaitForStatus(Disconnected)
			fmt.Println("Cycle complete.")
		}
	}()

	// Stop the state machine after 10 seconds
	time.Sleep(10 * time.Second)
	sm.Stop()

	fmt.Println("State machine stopped.")
}

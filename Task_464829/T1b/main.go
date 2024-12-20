package main  
import (  
    "fmt"
    "time"
)

// NetworkStatus represents the different states of the network connection
type NetworkStatus int

const (
    // Connected status indicates a successful network connection
    Connected NetworkStatus = iota
    // Disconnected status indicates that the network connection is lost
    Disconnected
    // Error status indicates an error in the network connection
    Error
)

// StateMachine represents the state machine for handling network statuses
type StateMachine struct {
    currentStatus NetworkStatus
    transition    chan NetworkStatus
}

// NewStateMachine creates a new instance of the StateMachine with the initial status set to Disconnected
func NewStateMachine() *StateMachine {
    return &StateMachine{
        currentStatus: Disconnected,
        transition:    make(chan NetworkStatus),
    }
}

// Run starts the state machine and handles transitions between states
func (sm *StateMachine) Run() {
    for {
        select {
        case nextStatus := <-sm.transition:
            sm.handleTransition(nextStatus)
        }
    }
}

// handleTransition handles the transition between the current and next status
func (sm *StateMachine) handleTransition(nextStatus NetworkStatus) {
    switch {
    case sm.currentStatus == Connected && nextStatus == Disconnected:
        fmt.Println("Network connection lost.")
    case sm.currentStatus == Disconnected && nextStatus == Connected:
        fmt.Println("Network connection established.")
    case sm.currentStatus == Error && nextStatus == Disconnected:
        fmt.Println("Network error resolved. Connection lost.")
    case sm.currentStatus == Disconnected && nextStatus == Error:
        fmt.Println("Network error occurred.")
    case sm.currentStatus == Connected && nextStatus == Error:
        fmt.Println("Network error occurred while connected.")
    }
    sm.currentStatus = nextStatus
}

// Transition transitions the state machine to the specified status
func (sm *StateMachine) Transition(status NetworkStatus) {
    sm.transition <- status
}

func main() {
    // Create a new state machine
    sm := NewStateMachine()
    
    // Start the state machine goroutine
    go sm.Run()
    
    // Simulate network status changes
    sm.Transition(Connected)
    time.Sleep(2 * time.Second)
    sm.Transition(Disconnected)
    time.Sleep(1 * time.Second)
    sm.Transition(Error)
    time.Sleep(2 * time.Second)
    sm.Transition(Connected)
    time.Sleep(1 * time.Second)
    sm.Transition(Disconnected)
}
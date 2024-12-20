package main  
import (  
    "fmt"
    "net"
    "time"
)

// NetworkStatus represents the different states of the network connection
type NetworkStatus int

const (
    Connected NetworkStatus = iota
    Disconnected
    Error
)

// StateMachine represents the state machine for handling network statuses
type StateMachine struct {
    currentStatus NetworkStatus
    done          chan struct{}
    statusChange  chan NetworkStatus
}

// NewStateMachine creates a new instance of the StateMachine with the initial status set to Disconnected.
func NewStateMachine() *StateMachine {
    return &StateMachine{
        currentStatus: Disconnected,
        done:          make(chan struct{}),
        statusChange:  make(chan NetworkStatus),
    }
}

// Start the state machine and listen for status changes.
func (sm *StateMachine) Start() {
    go func() {
        for {
            select {
            case nextStatus := <-sm.statusChange:
                sm.transitionTo(nextStatus)
            case <-sm.done:
                return
            }
        }
    }()
}

// Stop the state machine.
func (sm *StateMachine) Stop() {
    close(sm.done)
}

// Transition to the new status and print the transition message.
func (sm *StateMachine) transitionTo(nextStatus NetworkStatus) {
    if sm.currentStatus == nextStatus {
        return
    }
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

// CheckNetworkStatus periodically checks the network status and updates the state machine accordingly.
func (sm *StateMachine) CheckNetworkStatus() {
    ticker := time.NewTicker(2 * time.Second) // Check status every 2 seconds
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            if isConnected() {
                sm.statusChange <- Connected
            } else {
                sm.statusChange <- Disconnected
            }
        case <-sm.done:
            return
        }
    }
}

// isConnected attempts to make a simple TCP connection to check the network status.
// In a real application, you should use a more appropriate method to check the network status.
func isConnected() bool {
    conn, err := net.Dial("tcp", "example.com:80")
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

func main() {
    // Create a new state machine
    sm := NewStateMachine()
    
    // Start the state machine
    sm.Start()
    
    // Start the network status check goroutine
    go sm.CheckNetworkStatus()
    
    // Simulate some time passing
    time.Sleep(10 * time.Second)
    
    // Stop the state machine
    sm.Stop()
}
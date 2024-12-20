package main  
import (  
 "fmt"
 "time"
)

const (
	Connected = iota
	Disconnected
	Error
)

type NetworkStateMachine struct {
	state int
}

func NewNetworkStateMachine() *NetworkStateMachine {
	return &NetworkStateMachine{state: Disconnected}
}

func (s *NetworkStateMachine) GetState() int {
	return s.state
}

func (s *NetworkStateMachine) SetState(state int) {
	switch state {
	case Connected, Disconnected, Error:
		s.state = state
	default:
		panic("Invalid state")
	}
}

func (s *NetworkStateMachine) IsConnected() bool {
	return s.state == Connected
}

func main() {
	stateMachine := NewNetworkStateMachine()

	// Initial state is Disconnected
	fmt.Println("Current state:", getStateString(stateMachine.GetState()))

	// Simulate network connectivity changes
	stateMachine.SetState(Connected)
	fmt.Println("Current state:", getStateString(stateMachine.GetState()))

	// Simulate a network error
	stateMachine.SetState(Error)
	fmt.Println("Current state:", getStateString(stateMachine.GetState()))

	// Simulate network connectivity recovery
	time.Sleep(2 * time.Second)
	stateMachine.SetState(Connected)
	fmt.Println("Current state:", getStateString(stateMachine.GetState()))
}

func getStateString(state int) string {
	switch state {
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
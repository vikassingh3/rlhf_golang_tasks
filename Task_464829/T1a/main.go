package main

import (
	"fmt"
)

// State enum
type State string

const (
	StateConnected  State = "Connected"
	StateDisconnected State = "Disconnected"
	StateError        State = "Error"
)

// StateMachine struct to represent the state machine
type StateMachine struct {
	currentState State
	transitions  map[State]map[State]func()
}

// NewStateMachine creates a new state machine
func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		currentState: StateDisconnected,
		transitions: map[State]map[State]func(){
			StateConnected: {
				StateDisconnected: func() { fmt.Println("Transitioning from Connected to Disconnected") },
				StateError:        func() { fmt.Println("Transitioning from Connected to Error") },
			},
			StateDisconnected: {
				StateConnected: func() { fmt.Println("Transitioning from Disconnected to Connected") },
			},
			StateError: {
				StateConnected: func() { fmt.Println("Transitioning from Error to Connected") },
				StateDisconnected: func() { fmt.Println("Transitioning from Error to Disconnected") },
			},
		},
	}

	return sm
}

// Transition attempts to transition the state machine to a new state
func (sm *StateMachine) Transition(newState State) {
	if transitionFunc, ok := sm.transitions[sm.currentState][newState]; ok {
		transitionFunc()
		sm.currentState = newState
	} else {
		fmt.Println("Invalid transition from", sm.currentState, "to", newState)
	}
}

// CurrentState returns the current state of the state machine
func (sm *StateMachine) CurrentState() State {
	return sm.currentState
}

func main() {
	sm := NewStateMachine()

	fmt.Println("Initial State:", sm.CurrentState())

	sm.Transition(StateConnected)
	fmt.Println("Current State:", sm.CurrentState())

	sm.Transition(StateError)
	fmt.Println("Current State:", sm.CurrentState())

	sm.Transition(StateDisconnected)
	fmt.Println("Current State:", sm.CurrentState())

	sm.Transition(StateError)
	fmt.Println("Current State:", sm.CurrentState())

	sm.Transition(StateConnected)
	fmt.Println("Current State:", sm.CurrentState())
}
package main

import (
	"fmt"
)

// StateMachine represents a simple state machine
type StateMachine struct {
	CurrentState string
}

// InvalidStateError is used when an invalid state transition occurs
type InvalidStateError struct {
	Current  string
	Expected string
	message  string
}

func (e *InvalidStateError) Error() string {
	return fmt.Sprintf("%s: expected %s, but got %s", e.message, e.Expected, e.Current)
}

// TransitionError is used when an unknown transition event is encountered
type TransitionError struct {
	message string
}

func (e *TransitionError) Error() string {
	return e.message
}

// The Transition method of the StateMachine struct
func (sm *StateMachine) Transition(event string, newState string) error {
	switch event {
	case "start":
		if sm.CurrentState != "idle" {
			return &InvalidStateError{
				Current:  sm.CurrentState,
				Expected: "idle",
				message:  "cannot start from non-idle state",
			}
		}
		sm.CurrentState = newState
	case "stop":
		if sm.CurrentState != "running" {
			return &InvalidStateError{
				Current:  sm.CurrentState,
				Expected: "running",
				message:  "cannot stop from non-running state",
			}
		}
		sm.CurrentState = newState
	// ... other transition cases ...
	default:
		return &TransitionError{message: "unknown transition event"}
	}
	return nil
}

func main() {
	// Example usage of the StateMachine
	sm := &StateMachine{CurrentState: "idle"}
	err := sm.Transition("start", "running")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Current state after start:", sm.CurrentState)
	}
}

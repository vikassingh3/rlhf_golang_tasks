package main

import (
	"fmt"
)

// Custom error types
type InvalidStateError struct {
	Current  string
	Expected string
	message  string
}

func (e *InvalidStateError) Error() string {
	return fmt.Sprintf("%s: expected state '%s', but was in state '%s'", e.message, e.Expected, e.Current)
}

type TransitionError struct {
	message string
}

func (e *TransitionError) Error() string {
	return e.message
}

// State machine struct
type StateMachine struct {
	CurrentState string
}

func (sm *StateMachine) Transition(event string, newState string) error {
	switch event {
	case "start":
		if sm.CurrentState != "idle" {
			return &InvalidStateError{Current: sm.CurrentState, Expected: "idle", message: "cannot start from non-idle state"}
		}
		sm.CurrentState = newState
	case "stop":
		if sm.CurrentState != "running" {
			return &InvalidStateError{Current: sm.CurrentState, Expected: "running", message: "cannot stop from non-idle state"}
		}
		sm.CurrentState = newState
	case "pause":
		if sm.CurrentState != "running" {
			return &InvalidStateError{Current: sm.CurrentState, Expected: "running", message: "cannot pause from non-idle state"}
		}
		sm.CurrentState = newState
	case "resume":
		if sm.CurrentState != "paused" {
			return &InvalidStateError{Current: sm.CurrentState, Expected: "paused", message: "cannot resume from non-idle state"}
		}
		sm.CurrentState = newState
	default:
		return &TransitionError{message: "unknown transition event"}
	}
	return nil
}

// Function that uses the state machine and handles errors
func useStateMachine() {
	sm := StateMachine{CurrentState: "idle"}

	err := sm.Transition("start", "running")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sm.Transition("stop", "stopped")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sm.Transition("start", "running")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sm.Transition("pause", "paused")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sm.Transition("resume", "running")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sm.Transition("invalid", "unknown")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("State machine finished successfully.")
}

func main() {
	useStateMachine()
}

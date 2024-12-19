package main

import (
	"fmt"
)

// Custom error types
type InvalidStateError struct {
	Current    string
	Expected   string
	Message    string
	error
}

func (e *InvalidStateError) Error() string {
	return fmt.Sprintf("%s: expected state '%s', but was in state '%s'", e.Message, e.Expected, e.Current)
}

type TransitionError struct {
	Message string
	error
}

func (e *TransitionError) Error() string {
	return e.Message
}

// State machine struct
type StateMachine struct {
	CurrentState string
}

func (sm *StateMachine) Transition(event string, newState string) error {
	switch event {
	case "start":
		if sm.CurrentState != "idle" {
			return &InvalidStateError{
				Current:    sm.CurrentState,
				Expected:   "idle",
				Message:    "cannot start from non-idle state",
				error:      nil,
			}
		}
		sm.CurrentState = newState
	case "stop":
		if sm.CurrentState != "running" {
			return &InvalidStateError{
				Current:    sm.CurrentState,
				Expected:   "running",
				Message:    "cannot stop from non-idle state",
				error:      nil,
			}
		}
		sm.CurrentState = newState
	case "pause":
		if sm.CurrentState != "running" {
			return &InvalidStateError{
				Current:    sm.CurrentState,
				Expected:   "running",
				Message:    "cannot pause from non-idle state",
				error:      nil,
			}
		}
		sm.CurrentState = newState
	case "resume":
		if sm.CurrentState != "paused" {
			return &InvalidStateError{
				Current:    sm.CurrentState,
				Expected:   "paused",
				Message:    "cannot resume from non-idle state",
				error:      nil,
			}
		}
		sm.CurrentState = newState
	default:
		return &TransitionError{
			Message: "unknown transition event",
			error:   nil,
		}
	}
	return nil
}

// Function that uses the state machine and handles errors
func useStateMachine() {
	sm := StateMachine{CurrentState: "idle"}

	err := sm.Transition("start", "running")
	if err != nil {
		handleError(err)
		return
	}

	err = sm.Transition("stop", "stopped")
	if err != nil {
		handleError(err)
		return
	}

	err = sm.Transition("start", "running")
	if err != nil {
		handleError(err)
		return
	}

	err = sm.Transition("pause", "paused")
	if err != nil {
		handleError(err)
		return
	}

	err = sm.Transition("resume", "running")
	if err != nil {
		handleError(err)
		return
	}

	err = sm.Transition("invalid", "unknown")
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("State machine finished successfully.")
}

// Function to handle different types of errors
func handleError(err error) {
	switch e := err.(type) {
	case *InvalidStateError:
		fmt.Println("Invalid state transition:", e)
	case *TransitionError:
		fmt.Println("Transition error:", e)
	default:
		fmt.Println("An unknown error occurred:", err)
	}
}

func main() {
	useStateMachine()
}
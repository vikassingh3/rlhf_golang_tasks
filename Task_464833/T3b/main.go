package main

import (
	"fmt"
)

type InvalidTransitionError struct {
	CurrentState string
	ExpectedTransitions []string
	ActualTransition string
}

func (e *InvalidTransitionError) Error() string {
	return fmt.Sprintf("invalid transition from '%s': expected %s, got '%s'", e.CurrentState, e.ExpectedTransitions, e.ActualTransition)
}

type StateMachine struct {
	currentState string
	transitions map[string][]string
}

func (sm *StateMachine) Transition(transition string) error {
	if validTransitions, ok := sm.transitions[sm.currentState]; ok {
		for _, validTransition := range validTransitions {
			if validTransition == transition {
				sm.currentState = transition
				return nil
			}
		}
	}
	return &InvalidTransitionError{
		CurrentState:        sm.currentState,
		ExpectedTransitions: sm.transitions[sm.currentState],
		ActualTransition:     transition,
	}
}

func main() {
	sm := StateMachine{
		currentState: "start",
		transitions: map[string][]string{
			"start":  {"run", "stop"},
			"run":    {"stop", "pause"},
			"pause":  {"run", "stop"},
			"stop":   {"start"},
		},
	}

	err := sm.Transition("run")
	if err != nil {
		fmt.Println(err)
	}

	err = sm.Transition("invalid")
	if err != nil {
		fmt.Println(err)
	}
}
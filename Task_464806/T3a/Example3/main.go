package main

import (
	"fmt"
)

// Custom error types for resource management
type ResourceError struct {
	Resource string
	Errors    error
}

func (e *ResourceError) Error() string {
	return fmt.Sprintf("Resource management error with '%s': %s", e.Resource, e.Error)
}

// State machine for resource management
type ResourceManager struct {
	Resource   string
	State      string
	isAquired bool
}

func (rm *ResourceManager) transition(newState string, err error) error {
	if err != nil {
		return &ResourceError{Resource: rm.Resource, Errors: err}
	}
	rm.State = newState
	if newState == "aquired" {
		rm.isAquired = true
	} else {
		rm.isAquired = false
	}
	return nil
}

func (rm *ResourceManager) Acquire() error {
	if rm.isAquired {
		return &ResourceError{Resource: rm.Resource, Errors: fmt.Errorf("resource already aquired")}
	}
	return rm.transition("aquired", nil)
}

func (rm *ResourceManager) Release() error {
	if !rm.isAquired {
		return &ResourceError{Resource: rm.Resource, Errors: fmt.Errorf("resource not aquired")}
	}
	return rm.transition("released", nil)
}

func main() {
	rm := ResourceManager{Resource: "Database", State: "idle", isAquired: false}

	err := rm.Acquire()
	if err != nil {
		fmt.Println("Error acquiring resource:", err)
	}

	err = rm.Acquire()
	if err != nil {
		fmt.Println("Error acquiring resource:", err)
	}

	err = rm.Release()
	if err != nil {
		fmt.Println("Error releasing resource:", err)
	}
}
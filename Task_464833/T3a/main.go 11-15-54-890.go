package main

import (
	"fmt"
	"net/http"
)

// Custom error types for user registration process
type RegistrationError struct {
	Step  string
	Error error
}

func (e *RegistrationError) Error() string {
	return fmt.Sprintf("Registration failed at step '%s': %s", e.Step, e.Error)
}

// State machine for user registration
type RegistrationMachine struct {
	CurrentStep string
}

func (rm *RegistrationMachine) transition(step string, err error) (error, http.StatusCode) {
	if err != nil {
		return &RegistrationError{Step: step, Error: err}, http.StatusInternalServerError
	}
	rm.CurrentStep = step
	return nil, http.StatusOK
}

func (rm *RegistrationMachine) StartRegistration() (error, http.StatusCode) {
	return rm.transition("start", nil)
}

func (rm *RegistrationMachine) RegisterEmail(email string) (error, http.StatusCode) {
	if email == "" {
		return &RegistrationError{Step: "register_email", Error: fmt.Errorf("email is required")}, http.StatusBadRequest
	}
	return rm.transition("register_email", nil)
}

func (rm *RegistrationMachine) RegisterPassword(password string) (error, http.StatusCode) {
	if password == "" {
		return &RegistrationError{Step: "register_password", Error: fmt.Errorf("password is required")}, http.StatusBadRequest
	}
	return rm.transition("register_password", nil)
}

func (rm *RegistrationMachine) CompleteRegistration() (error, http.StatusCode) {
	return rm.transition("complete_registration", nil)
}

func main() {
	rm := RegistrationMachine{CurrentStep: ""}

	err, status := rm.StartRegistration()
	if err != nil {
		fmt.Println("Error starting registration:", err)
	} else {
		fmt.Println("Registration started successfully")
	}

	err, status = rm.RegisterEmail("")
	if err != nil {
		fmt.Println("Error registering email:", err)
	} else {
		fmt.Println("Email registered successfully")
	}

	err, status = rm.RegisterPassword("")
	if err != nil {
		fmt.Println("Error registering password:", err)
	} else {
		fmt.Println("Password registered successfully")
	}

	err, status = rm.CompleteRegistration()
	if err != nil {
		fmt.Println("Error completing registration:", err)
	} else {
		fmt.Println("Registration completed successfully")
	}
}


package main

import (
	"fmt"
	"sync"
	"time"
)

// Custom error types for message processing
type MessageError struct {
	Step  string
	Error error
}

func (e *MessageError) Error() string {
	return fmt.Sprintf("Message processing failed at step '%s': %s", e.Step, e.Error)
}

// State machine for message processing
type MessageProcessor struct {
	CurrentStep string
	mux         *sync.Mutex
}

func (mp *MessageProcessor) transition(step string, err error) error {
	if err != nil {
		return &MessageError{Step: step, Error: err}
	}
	mp.mux.Lock()
	mp.CurrentStep = step
	mp.mux.Unlock()
	return nil
}

func (mp *MessageProcessor) ReceiveMessage() error {
	return mp.transition("receive", nil)
}

func (mp *MessageProcessor) ValidateMessage(msg string) error {
	if msg == "" {
		return &MessageError{Step: "validate", Error: fmt.Errorf("message is empty")}
	}
	return mp.transition("validate", nil)
}

func (mp *MessageProcessor) SendResponse(resp string) error {
	return mp.transition("send_response", nil)
}

func (mp *MessageProcessor) MarkComplete() error {
	return mp.transition("mark_complete", nil)
}

func main() {
	mp := MessageProcessor{CurrentStep: "", mux: &sync.Mutex{}}

	err := mp.ReceiveMessage()
	if err != nil {
		fmt.Println("Error receiving message:", err)
	}

	err = mp.ValidateMessage("")
	if err != nil {
		fmt.Println("Error validating message:", err)
	}

	err = mp.SendResponse("OK")
	if err != nil {
		fmt.Println("Error sending response:", err)
	}

	err = mp.MarkComplete()
	if err != nil {
		fmt.Println("Error marking complete:", err)
	}

	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
)

// Custom error types for resource management
type ResourceError struct {
	Resource string
	Error    error
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
		return &ResourceError{Resource: rm.Resource, Error: err}
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
		return &ResourceError{Resource: rm.Resource, Error: fmt.Errorf("resource already aquired")}
	}
	return rm.transition("aquired", nil)
}

func (rm *ResourceManager) Release() error {
	if !rm.isAquired {
		return &ResourceError{Resource: rm.Resource, Error: fmt.Errorf("resource not aquired")}
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
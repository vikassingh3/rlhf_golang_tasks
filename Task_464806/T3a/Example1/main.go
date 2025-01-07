package main

import (
	"fmt"
	"net/http"
)

// Custom error types for user registration process
type RegistrationError struct {
	Step  string
	Errors error // Changed 'Errors' to 'Error'
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
		return &RegistrationError{Step: step, Error: err}, http.StatusInternalServerError // Fixed 'Errors' to 'Error'
	}
	rm.CurrentStep = step
	return nil, http.StatusOK
}

func (rm *RegistrationMachine) StartRegistration() (error, http.StatusCode) {
	return rm.transition("start", nil)
}

func (rm *RegistrationMachine) RegisterEmail(email string) (error, http.StatusCode) {
	if email == "" {
		return &RegistrationError{Step: "register_email", Error: fmt.Errorf("email is required")}, http.StatusBadRequest // Fixed 'Errors' to 'Error'
	}
	return rm.transition("register_email", nil)
}

func (rm *RegistrationMachine) RegisterPassword(password string) (error, http.StatusCode) {
	if password == "" {
		return &RegistrationError{Step: "register_password", Error: fmt.Errorf("password is required")}, http.StatusBadRequest // Fixed 'Errors' to 'Error'
	}
	return rm.transition("register_password", nil)
}

func (rm *RegistrationMachine) CompleteRegistration() (error, http.StatusCode) {
	return rm.transition("complete_registration", nil)
}

func main() {
	rm := RegistrationMachine{CurrentStep: ""}

	err, _ := rm.StartRegistration()
	if err != nil {
		fmt.Println("Error starting registration:", err)
	} else {
		fmt.Println("Registration started successfully")
	}

	err, _ = rm.RegisterEmail("")
	if err != nil {
		fmt.Println("Error registering email:", err)
	} else {
		fmt.Println("Email registered successfully")
	}

	err, _ = rm.RegisterPassword("")
	if err != nil {
		fmt.Println("Error registering password:", err)
	} else {
		fmt.Println("Password registered successfully")
	}

	err, _ = rm.CompleteRegistration()
	if err != nil {
		fmt.Println("Error completing registration:", err)
	} else {
		fmt.Println("Registration completed successfully")
	}
}

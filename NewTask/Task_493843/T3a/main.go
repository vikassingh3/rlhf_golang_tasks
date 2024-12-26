package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

type Feedback struct {
	Field      string `json:"field"`
	Message    string `json:"message"`
	IsValid    bool   `json:"is_valid"`
}

var feedbacks []Feedback
var feedbacksMu sync.Mutex

func handleFeedback(w http.ResponseWriter, r *http.Request) {
	feedbacksMu.Lock()
	defer feedbacksMu.Unlock()
	data, err := json.MarshalIndent(feedbacks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func logFeedback(field string, isValid bool, message string) {
	feedbacksMu.Lock()
	defer feedbacksMu.Unlock()
	feedbacks = append(feedbacks, Feedback{
		Field:      field,
		Message:    message,
		IsValid:    isValid,
	})
}

func validateField(field string, value string) (bool, string) {
	switch field {
	case "email":
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(value) {
			return false, "Invalid email format."
		}
	case "userID":
		if _, err := strconv.Atoi(value); err != nil {
			return false, "User ID must be a valid number."
		}
	}
	return true, ""
}

func main() {
	http.HandleFunc("/api/feedback", handleFeedback)
	http.HandleFunc("/api/validate", func(w http.ResponseWriter, r *http.Request) {
		field := r.URL.Query().Get("field")
		value := r.URL.Query().Get("value")

		isValid, message := validateField(field, value)
		logFeedback(field, isValid, message)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
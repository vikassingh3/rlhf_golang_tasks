package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "regexp"
)

// Define a simple structure to hold query parameters
type QueryParameters struct {
    UserID string `json:"userID"`
    Email  string `json:"email"`
}

// A regular expression to validate an email
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Handler function to process incoming HTTP requests
func parameterHandler(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters from the URL
    queryParams := r.URL.Query()

    var params QueryParameters
    for key, value := range queryParams {
        // Validate each parameter
        switch key {
        case "userID":
            // Example validation: UserID should be a non-empty string of digits
            if len(value) == 0 || !allDigits(value[0]) {
                logInvalidParameter(key, value[0]) // Pass the first element of the slice
                params.UserID = ""
            } else {
                params.UserID = value[0]
            }
        case "email":
            // Validate: Email should match the email regex pattern
            if !emailRegex.MatchString(value[0]) {
                logInvalidParameter(key, value[0]) // Pass the first element of the slice
                params.Email = ""
            } else {
                params.Email = value[0]
            }
        }
    }

    // Convert the validated parameters back to JSON for response
    jsonResponse, err := json.Marshal(params)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send the response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

// Function to log invalid parameters
func logInvalidParameter(key, value string) {
    fmt.Printf("Invalid Parameter: %s=%v\n", key, value)
}

// Check if all characters are digits
func allDigits(str string) bool {
    for _, char := range str {
        if char < '0' || char > '9' {
            return false
        }
    }
    return true
}

// Main function to start the server
func main() {
	// Define the route for the handler function
	http.HandleFunc("/api/parameters", parameterHandler)

	// Start the HTTP server
	fmt.Println("Server running on http://localhost:8080/api/parameters")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
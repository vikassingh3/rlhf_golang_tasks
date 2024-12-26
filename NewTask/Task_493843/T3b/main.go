package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type QueryParameters struct {
    UserID string `json:"userID"`
    Email  string `json:"email"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var userIDRegex = regexp.MustCompile(`^[0-9]+$`)

func validateParameters(params QueryParameters) (string, bool) {
    if !userIDRegex.MatchString(params.UserID) {
        return "Invalid UserID format. Please enter a valid number.", false
    }

    if !emailRegex.MatchString(params.Email) {
        return "Invalid email format. Please enter a valid email address.", false
    }

    return "", true // Valid parameters
}

func parameterHandler(w http.ResponseWriter, r *http.Request) {
    var params QueryParameters

    if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
        http.Error(w, "Invalid JSON data in request body", http.StatusBadRequest)
        return
    }

    errMsg, isValid := validateParameters(params)

    if !isValid {
        http.Error(w, errMsg, http.StatusBadRequest)
        return
    }

    // Process valid parameters here
    // ...

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Parameters validated successfully"})
}

func main() {
    http.HandleFunc("/api/parameters", parameterHandler)
    fmt.Println("Server running on http://localhost:8080/api/parameters")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
 
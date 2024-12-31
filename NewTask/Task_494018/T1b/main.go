package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTP function that processes data slices
func ProcessData(w http.ResponseWriter, r *http.Request) {
	// Parse the input JSON
	var data []int
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Process the data
	processedData := make([]int, len(data))
	for i, value := range data {
		processedData[i] = value * 2 // Example: multiply each value by 2
	}

	// Return the processed data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processedData)
}

func main() {
	// Register the HTTP handler
	http.HandleFunc("/processData", ProcessData)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

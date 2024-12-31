package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a data structure to hold your data
type DataItem struct {
	Value int `json:"value"`
}

// Define a function to process a slice of DataItems
func processDataSlice(data []DataItem) ([]DataItem, error) {
	processedData := make([]DataItem, len(data))

	for i, item := range data {
		// Simulate some processing
		processedData[i].Value *= 2
		fmt.Printf("Processing item %d: original=%d, processed=%d\n", i+1, item.Value, processedData[i].Value)
	}

	return processedData, nil
}

// Handle HTTP requests to process data
func processDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data []DataItem
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	processedData, err := processDataSlice(data)
	if err != nil {
		http.Error(w, "Error processing data", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(processedData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/process-data", processDataHandler)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
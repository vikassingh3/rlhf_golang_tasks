package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func paginate(w http.ResponseWriter, r *http.Request) {
	// Sample data for demonstration purposes (in a real application, you would fetch this from a database)
	allData := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10", "Item 11", "Item 12", "Item 13", "Item 14", "Item 15"}

	// Parse page and size query parameters
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page := 1
	size := 5

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
	}

	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			http.Error(w, "Invalid size parameter", http.StatusBadRequest)
			return
		}
	}

	// Calculate start index and end index
	startIndex := (page - 1) * size
	endIndex := startIndex + size

	// Check if the end index is greater than the total data size
	if endIndex > len(allData) {
		endIndex = len(allData)
	}

	// Calculate total pages
	totalPages := len(allData) / size
	if len(allData)%size > 0 {
		totalPages++
	}

	// Create a struct to hold the paginated data
	paginatedData := struct {
		Page        int      `json:"page"`
		Size        int      `json:"size"`
		TotalPages  int      `json:"totalPages"`
		TotalElements int      `json:"totalElements"`
		Elements    []string `json:"elements"`
	}{
		Page:        page,
		Size:        size,
		TotalPages:  totalPages,
		TotalElements: len(allData),
		Elements:    allData[startIndex:endIndex],
	}

	// Marshal the struct to JSON and return it as the response body
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paginatedData); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/data", paginate)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}  
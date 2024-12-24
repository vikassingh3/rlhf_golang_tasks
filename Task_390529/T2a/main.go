package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// Person is the struct representing the data model
type Person struct {
	ID      int64 `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
}

// Pagination contains pagination information
type Pagination struct {
	Page        int    `json:"page"`
	Size        int    `json:"size"`
	TotalPages  int    `json:"totalPages"`
	TotalElements int    `json:"totalElements"`
	Elements    []Person `json:"elements"`
}

// Handler paginates and returns a list of persons
func handler(w http.ResponseWriter, r *http.Request) {
	// Parse page and size parameters from the request query
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1 // Default page if error in parsing
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 {
		size = 10 // Default page size if error or non-positive size
	}

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./people.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Count total number of elements
	var totalElements int
	err = db.QueryRow("SELECT COUNT(*) FROM persons").Scan(&totalElements)
	if err != nil {
		fmt.Println("Error counting elements:", err)
		http.Error(w, "Error counting elements", http.StatusInternalServerError)
		return
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalElements) / float64(size)))

	// Avoid trying to access a non-existent page
	if page > totalPages {
		page = totalPages
	}

	// Execute query with pagination
	rows, err := db.Query("SELECT id, name, age, email FROM persons ORDER BY id LIMIT $1 OFFSET $2", size, (page-1)*size)
	if err != nil {
		fmt.Println("Error querying database:", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var elements []Person
	for rows.Next() {
		var person Person
		err = rows.Scan(&person.ID, &person.Name, &person.Age, &person.Email)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		elements = append(elements, person)
	}

	// Create pagination object
	pagination := Pagination{
		Page:        page,
		Size:        size,
		TotalPages:  totalPages,
		TotalElements: totalElements,
		Elements:    elements,
	}

	// Encode and respond with JSON
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(pagination); err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/persons", handler)
	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
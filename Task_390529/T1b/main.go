package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Pagination struct {
	Page          int           `json:"page"`
	Size          int           `json:"size"`
	TotalPages    int           `json:"totalPages"`
	TotalElements int           `json:"totalElements"`
	Elements      []interface{} `json:"elements"`
}

type YourStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add other fields as needed
}

// Initialize database connection
func initDB() (*sql.DB, error) {
	// Open a connection to the PostgreSQL database
	// Replace with your actual database credentials
	connStr := "user=postgres dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Get paginated results from the database
func getPaginatedResults(db *sql.DB, page, size int) ([]YourStruct, int, error) {
	// Fetch the paginated data from the database
	query := `SELECT * FROM your_table ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.Query(query, size, (page-1)*size)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var elements []YourStruct
	for rows.Next() {
		var element YourStruct
		err := rows.Scan(&element.ID, &element.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		elements = append(elements, element)
	}

	// Count the total number of elements in the dataset
	var totalElements int
	err = db.QueryRow("SELECT COUNT(*) FROM your_table").Scan(&totalElements)
	if err != nil {
		return nil, 0, err
	}

	return elements, totalElements, nil
}

// Parse pagination query parameters from the request
func parsePaginationParams(r *http.Request) (int, int) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	// Default values if not provided
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10 // Default page size
	}

	return page, size
}

// Handler to serve paginated data
func getPaginatedData(w http.ResponseWriter, r *http.Request) {
	// Initialize the database connection
	db, err := initDB()
	if err != nil {
		http.Error(w, "Error initializing database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Parse pagination parameters
	page, size := parsePaginationParams(r)

	// Get paginated data from the database
	_, totalElements, err := getPaginatedResults(db, page, size)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Prepare pagination response
	pagination := Pagination{
		Page:          page,
		Size:          size,
		TotalPages:    int(math.Ceil(float64(totalElements) / float64(size))),
		TotalElements: totalElements,
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagination)
}

func main() {
	// Define routes
	http.HandleFunc("/api/data", getPaginatedData)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

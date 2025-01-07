package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection string
	connString := "host=localhost dbname=postgres password=root sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Database connection established!")

	// Create a Gin router
	router := gin.Default()

	// Define a route that performs a database query
	router.GET("/data", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM table_name")
		if err != nil {
			if pingErr := db.Ping(); pingErr != nil {
				log.Printf("Error pinging database during query: %v", pingErr)
				c.String(http.StatusInternalServerError, "Database connection lost")
				return
			}

			// For other types of errors, adjust your error handling logic
			c.String(http.StatusInternalServerError, "Failed to fetch data: %v", err)
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			row := make(map[string]interface{})
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to scan row: %v", err)
				return
			}
			row["id"] = id
			row["name"] = name
			results = append(results, row)
		}
		c.JSON(http.StatusOK, results)
	})

	// Start the server
	fmt.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

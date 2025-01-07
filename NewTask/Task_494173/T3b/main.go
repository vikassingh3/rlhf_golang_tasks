
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func connectDB(connString string) (*sql.DB, error) {
	// (Same code as before)
}

func handleUserTransaction(db *sql.DB, c *gin.Context) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Printf("Panic occurred during transaction: %v", p)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	// Perform database operations within the transaction
	userID := c.Param("id")
	var user User
	err = tx.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		tx.Rollback()
		log.Printf("Query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Additional database operations here

	// Commit the transaction if no errors occurred
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func main() {
	connString := "host=localhost dbname=mydatabase user=myuser password=mypassword sslmode=disable"
	db, err := connectDB(connString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/user/:id", func(c *gin.Context) {
		handleUserTransaction(db, c)
	})

	fmt.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
 
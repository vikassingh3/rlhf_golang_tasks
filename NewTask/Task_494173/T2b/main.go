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
    if err = db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }
    fmt.Println("Database connection established!")
    
    router := gin.Default()
    
    // Define a route to retrieve user by ID
    router.GET("/user/:id", func(c *gin.Context) {
        userID := c.Param("id")
        var user User
        // Query the database
        row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userID)
        err := row.Scan(&user.ID, &user.Name, &user.Email)
        if err == sql.ErrNoRows {
            // Handle case where user not found
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        } else if err != nil {
            // Handle other query errors
            log.Printf("Error querying database: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
            return
        }

        // Return the user if found
        c.JSON(http.StatusOK, user)
    })

    // Start the server
    fmt.Println("Starting server on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

// User struct defines the user model
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
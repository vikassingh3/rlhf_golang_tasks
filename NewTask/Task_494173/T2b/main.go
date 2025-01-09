package main  
import (  
    "database/sql"  
    "fmt"  
    _ "github.com/lib/pq" // Import the PostgreSQL driver  
)  

// Define custom error types for different types of database errors
type DatabaseError struct {  
    Message string  
}  

func (e *DatabaseError) Error() string {  
    return fmt.Sprintf("Database Error: %s", e.Message)  
}  

func main() {  
    connString := "host=localhost dbname=postgres password=root sslmode=disable"  
    db, err := sql.Open("postgres", connString)  
    if err != nil {  
        handleError(err)  
        return  
    }  
    defer db.Close()  

    err = performSomeOperation(db)  
    if err != nil {  
        handleError(err)  
        return  
    }  

    fmt.Println("Database operation successful!")  
}  

func performSomeOperation(db *sql.DB) error {  
    // Example query that might fail
    _, err := db.Exec("INSERT INTO users (name) VALUES ($1)", "John Doe")  
    if err != nil {  
        return &DatabaseError{Message: "Error executing query: " + err.Error()}  
    }  
    return nil  
}  

func handleError(err error) {  
    if dbErr, ok := err.(*DatabaseError); ok {  
        fmt.Printf("Error: %s\n", dbErr.Error())  
        // You can handle database errors more specifically here based on the dbErr.Message
    } else {  
        fmt.Printf("Unexpected Error: %v\n", err)  
    }  
}  
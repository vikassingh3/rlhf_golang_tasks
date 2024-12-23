package main  
import (  
   "database/sql"
   "fmt"
   _ "github.com/go-sql-driver/mysql" // Replace this with your desired database driver
)

func insertUser(db *sql.DB, name string, email string) (int64, error) {
   // Begin a transaction
   tx, err := db.Begin()
   if err != nil {
       return 0, err
   }

   // Insert user data into the database
   result, err := tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
   if err != nil {
       // Rollback the transaction if an error occurs
       tx.Rollback()
       return 0, err
   }

   // Get the ID of the inserted row
   userID, err := result.LastInsertId()
   if err != nil {
       // Rollback the transaction if an error occurs
       tx.Rollback()
       return 0, err
   }

   // Commit the transaction if all operations are successful
   if err = tx.Commit(); err != nil {
       return 0, err
   }

   return userID, nil
}

func main() {
   // Open the database connection
   db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
   if err != nil {
       panic(err.Error())
   }
   defer db.Close()

   // Example usage of the insertUser function
   name := "John Doe"
   email := "johndoe@example.com"
   userID, err := insertUser(db, name, email)
   if err != nil {
       fmt.Println("Error inserting user:", err)
   } else {
       fmt.Println("User inserted successfully with ID:", userID)
   }
}
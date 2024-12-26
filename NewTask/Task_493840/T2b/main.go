package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // Assuming you're using SQLite as the database
)

func main() {  
    db, err := sqlx.Open("sqlite3", "./example.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter your age: ")
    ageStr, _ := reader.ReadString('\n')
    ageStr = strings.TrimSpace(ageStr)
    age, _ := strconv.Atoi(ageStr)

    // Use parameterized query to insert user input into the database
    query := "INSERT INTO users (name, age) VALUES (?, ?)"
    _, err = db.Exec(query, name, age)
    if err != nil {
        panic(err)
    }

    fmt.Printf("User data inserted successfully.\n")
}  
  
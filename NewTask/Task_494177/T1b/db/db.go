package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Choose your database dialect
)

func main() {  
    // Open the database connection and set the connection timeout
    db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local&timeout=10s")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Set max idle connections
    db.DB().SetMaxIdleConns(10)

    // Set max open connections
    db.DB().SetMaxOpenConns(100)
    // ...
}
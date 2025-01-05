package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func main() {
	// Create a GORM database instance using SQLite
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Configure connection pooling options (only applies to database/sql connections)
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to configure database connection")
	}
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// AutoMigrate the User schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	// Create a new user
	user := User{Name: "John Doe"}
	db.Create(&user)

	// Read the user from the database
	var readUser User
	db.First(&readUser, "name = ?", user.Name)

	fmt.Println("User:", readUser)
}

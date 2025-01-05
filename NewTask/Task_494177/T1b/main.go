package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	dsn := "host=localhost user=postgres password=root dbname=teastall port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance")
	}

	// Set connection pool configurations
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Example operation
	type User struct {
		ID   uint   `gorm:"primaryKey"`
		Name string `gorm:"not null"`
	}

	// Auto migrate
	db.AutoMigrate(&User{})

	// Create a user
	db.Create(&User{Name: "John Doe"})
}

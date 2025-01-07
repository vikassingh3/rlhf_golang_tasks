package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Initialize the Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Use the default address
		Password: "",             // No password set
		DB:       0,              // Use the default database
	})

	// Ensure client is closed at the end of the program
	defer rdb.Close()

    // CRUD operations on Redis
	
	// Create
	err := rdb.Set("user:1", "Alice", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Read
	val, err := rdb.Get("user:1").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved user:", val)

	// Update
	err = rdb.Incr("user:1:age").Err()
	if err != nil {
		log.Fatal(err)
	}
	newAge, err := rdb.Get("user:1:age").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Alice's new age:", newAge)

	// Delete
	_, err = rdb.Del("user:1", "user:1:age").Result()
	if err != nil {
		log.Fatal(err)
	}
	val, err = rdb.Get("user:1").Result()
	if err == redis.Nil {
		fmt.Println("User not found.")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Unexpected error:", err)
	}
} 
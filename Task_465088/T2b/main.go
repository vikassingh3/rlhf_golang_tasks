package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Initialize the Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Use the default address
		Password: "",               // No password set
		DB:       0,                // Use the default database
	})

	// Create a context for Redis commands
	ctx := context.Background()

	// Ensure client is closed at the end of the program
	defer rdb.Close()

	// CRUD operations on Redis

	// Create
	err := rdb.Set(ctx, "user:1", "Alice", 0).Err() // Pass ctx as the first argument
	if err != nil {
		log.Fatal(err)
	}

	// Read
	val, err := rdb.Get(ctx, "user:1").Result() // Pass ctx as the first argument
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved user:", val)

	// Update
	err = rdb.Incr(ctx, "user:1:age").Err() // Pass ctx as the first argument
	if err != nil {
		log.Fatal(err)
	}
	newAge, err := rdb.Get(ctx, "user:1:age").Result() // Pass ctx as the first argument
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Alice's new age:", newAge)

	// Delete
	_, err = rdb.Del(ctx, "user:1", "user:1:age").Result() // Pass ctx as the first argument
	if err != nil {
		log.Fatal(err)
	}
	val, err = rdb.Get(ctx, "user:1").Result() // Pass ctx as the first argument
	if err == redis.Nil {
		fmt.Println("User not found.")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Unexpected error:", err)
	}
}

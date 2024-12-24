package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

// Initialize Redis client
func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Replace with your Redis server address
	})
	return client
}

// SetToken stores the JWT token in Redis
func setToken(client *redis.Client, token string) error {
	err := client.Set(ctx, "jwtToken", token, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("Error setting token: %v", err)
	}
	return nil
}

// GetToken retrieves the JWT token from Redis
func getToken(client *redis.Client) (string, error) {
	token, err := client.Get(ctx, "jwtToken").Result()
	if err != nil {
		return "", fmt.Errorf("Error getting token: %v", err)
	}
	return token, nil
}

// RemoveToken deletes the JWT token from Redis
func removeToken(client *redis.Client) error {
	err := client.Del(ctx, "jwtToken").Err()
	if err != nil {
		return fmt.Errorf("Error removing token: %v", err)
	}
	return nil
}

// Example usage: Simulate login process
func login(client *redis.Client) {
	// Simulate token retrieval after successful login
	token := "example.jwt.token"

	// Store the token securely
	err := setToken(client, token)
	if err != nil {
		log.Fatalf("Failed to set token: %v", err)
	} else {
		fmt.Println("Token successfully stored in Redis.")
	}
}

func validateToken(client *redis.Client) {
	// Retrieve the token
	token, err := getToken(client)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	} else if token == "" {
		fmt.Println("No token found.")
	} else {
		fmt.Println("Token retrieved:", token)
	}
}

func main() {
	// Initialize Redis client
	client := initRedis()

	// Simulate login
	login(client)

	// Validate the stored token
	validateToken(client)

	// Remove the token
	err := removeToken(client)
	if err != nil {
		log.Fatalf("Failed to remove token: %v", err)
	} else {
		fmt.Println("Token removed from Redis.")
	}
}

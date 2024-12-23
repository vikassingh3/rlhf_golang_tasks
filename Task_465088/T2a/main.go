package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    string `bson:"_id" json:"_id"`
	Name  string `bson:"name" json:"name"`
	Age   int    `bson:"age" json:"age"`
	Email string `bson:"email" json:"email"`
}

func main() {
	// Set up a client to connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Create a database and collection
	db := client.Database("example_db")
	usersCollection := db.Collection("users")

	// Create (Insert) a new user document
	user := User{
		ID:    "user1",
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	_, err = usersCollection.InsertOne(context.TODO(), user) // Ensure context is passed correctly
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	// Read (Retrieve) a user document
	result := usersCollection.FindOne(context.TODO(), bson.M{"_id": "user1"}) // Ensure context is passed correctly
	if result.Err() != nil {
		log.Fatalf("failed to find user: %v", err)
	}
	var retrievedUser User
	if err := result.Decode(&retrievedUser); err != nil {
		log.Fatalf("failed to decode user: %v", err)
	}
	fmt.Println("Retrieved User:", retrievedUser)

	// Update an existing user document
	update := bson.M{"$set": bson.M{"age": 31}}
	_, err = usersCollection.UpdateOne(context.TODO(), bson.M{"_id": "user1"}, update) // Ensure context is passed correctly
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	// Delete a user document
	_, err = usersCollection.DeleteOne(context.TODO(), bson.M{"_id": "user1"}) // Ensure context is passed correctly
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	fmt.Println("CRUD operations completed successfully.")
}

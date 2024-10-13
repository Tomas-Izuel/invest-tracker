package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var client *mongo.Client

// InitializeDB initializes the MongoDB connection with the Stable API
func InitializeDB(uri string, dbName string) {
	// Set Server API options for Stable API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Apply URI and Server API options to client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the server to confirm the connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Set the database for further usage
	db = client.Database(dbName)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

// GetCollection returns a MongoDB collection by name
func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

// DisconnectDB disconnects the client from MongoDB
func DisconnectDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	fmt.Println("Disconnected from MongoDB")
}

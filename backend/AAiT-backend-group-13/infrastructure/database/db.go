package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config holds the MongoDB connection configuration.
type Config struct {
	ConnectString string // MongoDB connection string.
}

var client *mongo.Client
var once sync.Once

// Connect initializes and returns a singleton MongoDB client.
func Connect(config Config) *mongo.Client {
	once.Do(func() {
		var err error
		clientOptions := options.Client().ApplyURI(config.ConnectString)
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Couldn't connect to MongoDB: %v", err)
		}

		if err := client.Ping(context.Background(), nil); err != nil {
			log.Fatalf("Could not ping MongoDB server: %v", err)
		}

		log.Println("MongoDB: Successfully connected!")
	})

	return client
}

// EnsureUserIndexes performs creating index.
func EnsureUserIndexes(client *mongo.Client, dbName string) {
	database := client.Database(dbName)
	usersCollection := database.Collection("users")

	// Define the index model
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"username": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	// Check if the index already exists
	indexNames, err := usersCollection.Indexes().List(context.TODO())
	if err != nil {
		log.Fatalf("Error listing indexes: %v", err)
	}

	var indexExists bool
	for indexNames.Next(context.TODO()) {
		var index bson.M
		if err := indexNames.Decode(&index); err != nil {
			log.Fatalf("Error decoding index: %v", err)
		}

		if name, ok := index["name"].(string); ok && name == "username_1" {
			indexExists = true
			break
		}
	}

	if !indexExists {
		// Create the index
		indexName, err := usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created index %s\n", indexName)
	} else {
		log.Println("Index already exists. No changes made.")
	}
}

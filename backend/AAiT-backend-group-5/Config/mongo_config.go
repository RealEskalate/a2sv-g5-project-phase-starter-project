package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(env *Env) *mongo.Client {
	// mongoURI := env.MONGO_URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client
}

func GetDatabase(client *mongo.Client, env *Env) *mongo.Database {
	dbName := env.DB_NAME
	if dbName == "" {
		log.Fatal("DB_NAME is not set in .env file")
	}
	return client.Database(dbName)
}

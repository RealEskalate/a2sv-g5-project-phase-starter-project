package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var BlogCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB(connectionString string) {

    clientOptions := options.Client().ApplyURI(connectionString)

    client, err := mongo.NewClient(clientOptions)

    if err != nil {
        log.Fatalf("Error creating MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    Client = client
    BlogCollection = client.Database("Starter_blog").Collection("Blogs")
    UserCollection = client.Database("Starter_blog").Collection("users")
}

package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
    mongoClient *mongo.Client
    once        sync.Once
)

// GetMongoClient returns a singleton instance of the MongoDB client.
func GetClient(uri string, dbName string) (*mongo.Database, error) {
    var err error
    once.Do(func() {
        clientOptions := options.Client().ApplyURI(uri)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        mongoClient, err = mongo.Connect(ctx, clientOptions)
        if err != nil {
            log.Fatal(err)
        }


        err = mongoClient.Ping(ctx, nil)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Successfully connected to MongoDB")
    })

    // Replace "your_database" with the name of your database
    return mongoClient.Database(dbName), err
}
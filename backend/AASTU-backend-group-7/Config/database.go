package Config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB successfully!")
	return nil
}

func ConnectDB() *mongo.Client {
	Envinit()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MONGO_CONNECTION_STRING)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = ping(client, ctx)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	return client
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client            *mongo.Client
	UserCollection    *mongo.Collection
	BlogCollection    *mongo.Collection
	CommentCollection *mongo.Collection
}

func NewMongoDBClient(ctx context.Context) (*MongoDBClient, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI not set in environment variables")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("blog_platform_db")
	userCollectionName := os.Getenv("USER_COLLECTION")
	blogCollectionName := os.Getenv("BLOG_COLLECTION")
	commentCollectionName := os.Getenv("COMMENT_COLLECTION")

	if userCollectionName == "" || blogCollectionName == "" || commentCollectionName == "" {
		return nil, fmt.Errorf("collection names must be set in the environment variables")
	}

	return &MongoDBClient{
		Client:            client,
		UserCollection:    db.Collection(userCollectionName),
		BlogCollection:    db.Collection(blogCollectionName),
		CommentCollection: db.Collection(commentCollectionName),
	}, nil
}

func (db *MongoDBClient) CloseDBConnection(ctx context.Context) error {
	err := db.Client.Disconnect(ctx)
	if err != nil {
		log.Printf("Error closing connection to MongoDB: %v", err)
		return err
	}
	log.Println("Connection to MongoDB closed successfully")
	return nil
}

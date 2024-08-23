package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var BlogCollection *mongo.Collection
var UserCollection *mongo.Collection
var LikeCollection *mongo.Collection
var CommentCollection *mongo.Collection

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
    LikeCollection = client.Database("Starter_blog").Collection("likes")
    CommentCollection = client.Database("Starter_blog").Collection("comments")

}

func CreateTextIndex(collection *mongo.Collection) {
    indexModel := mongo.IndexModel{
        Keys: bson.D{
            {Key: "tags", Value: "text"},
            {Key: "autorname", Value: "text"},
        },
        Options: options.Index().SetDefaultLanguage("english"),
    }

    _, err := collection.Indexes().CreateOne(context.Background(), indexModel)
    if err != nil {
        log.Fatal(err)
    }
}
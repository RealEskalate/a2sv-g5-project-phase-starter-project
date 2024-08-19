package db

import (
	"context"
	"log"
	"sync"

	"github.com/group13/blog/infrastructure/migration"
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

// Migrate performs creating index.
func Migrate(client *mongo.Client, dbName string) {
	database := client.Database(dbName)
	usersCollection := database.Collection("users")
	blogsCollection := database.Collection("blogs")

	//creates composite keys for users collection
	migration.CreateIndexWithIDAndFirstName(usersCollection)
	migration.CreateIndexWithIDAndLastName(usersCollection)

	//creates composite keys for blogs collection
	migration.CreateIndexWithIDAndTitle(blogsCollection)
	migration.CreateIndexWithIDAndContent(blogsCollection)
	migration.CreateIndexWithIDAndTags(blogsCollection)
	migration.CreateIndexWithIDAndLikeCount(blogsCollection)
	migration.CreateIndexWithIDAndCommentCount(blogsCollection)
	migration.CreateIndexWithIDAndDisLikeCount(blogsCollection)
	migration.CreateIndexWithIDAndCreatedDate(blogsCollection)
	migration.CreateIndexWithIDAndUpdatedDate(blogsCollection)
	migration.CreateIndexWithIDAndUserID(blogsCollection)

}

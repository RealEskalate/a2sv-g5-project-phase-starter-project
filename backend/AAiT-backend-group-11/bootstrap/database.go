package bootstrap

import (
	"context"
	"log"
	"time"

	"backend-starter-project/mongo"
)

var (
	mongoClient mongo.Client
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := env.DBUri

	clientOptions, err := mongo.NewClient(mongodbURI)
	err = clientOptions.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = clientOptions.Ping(ctx)
	if err != nil {
		log.Println("Failed to connect to MongoDB")
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	mongoClient = clientOptions
	

	if mongoClient == nil {
		return nil
	}

	return &mongoClient
}

func CheckDatabaseConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := mongoClient.Ping(ctx)
	return err
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := mongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

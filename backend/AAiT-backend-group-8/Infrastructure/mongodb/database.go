package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Client {
	uri := "mongodb://localhost:27017"
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CreateCollection(client *mongo.Client, databaseName, clientName string) *mongo.Collection {
	myCollection := client.Database(databaseName).Collection(clientName)
	return myCollection
}

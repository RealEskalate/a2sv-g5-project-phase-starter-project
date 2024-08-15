package infrastructure

import (
	"context"
	"fmt"
	"log"

	// "time"
	"blogs/mongo"
)

func ConnectDB(url string, DBname string) (mongo.Database, mongo.Client, error) {
	// Create a MongoDB client
	// Connect to MongoDB
	client, err := mongo.NewClient(url)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(context.Background())
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connected to MongoDB!")
	DB := client.Database(DBname)
	return DB, client, nil
}

func CreateCollection(database mongo.Database, collectionName string) mongo.Collection {
	collection := database.Collection(collectionName)
	return collection
}

func CloseDB(client mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}

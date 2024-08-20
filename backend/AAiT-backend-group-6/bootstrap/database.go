package bootstrap

import (
	"AAiT-backend-group-6/mongo"
	"context"
	"log"
	// "time"
)

func NewMongoDatabase(env *Env) mongo.Client {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// Create a new instance of your custom mongoClient
	client, err := mongo.NewClient("mongodb://localhost:27017")
	print(client)
	if err != nil {
		log.Fatal(err)
	}

	// Connect the client
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check the connection
	err = client.Ping(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

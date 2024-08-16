package config

import (
	"context"
	"log"
	"time"

	"blog/database"
)

func NewMongoDatabase(env *Env) database.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := env.MONGO_URI

	client, err := database.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client database.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

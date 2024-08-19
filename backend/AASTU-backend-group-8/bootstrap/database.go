package bootstrap

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() (*mongo.Client,error) {

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		return &mongo.Client{}, err
	}

	if err = client.Ping(context.TODO(),nil); err != nil {
		return &mongo.Client{}, err
	}

	return client, nil

}

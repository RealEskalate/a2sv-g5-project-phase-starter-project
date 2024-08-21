package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerConnection struct {
	Client *mongo.Client
}

func (SC *ServerConnection) Connect_could() {

	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env", err.Error())
	}

	url := os.Getenv("DB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	options := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client, connetion_err := mongo.Connect(context.TODO(), options)

	if connetion_err != nil {
		log.Panic("Failed to connect to server", connetion_err.Error())
	}

	if err := client.Database("BlogPost").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Panic("Ping failed", err.Error())
	}

	SC.Client = client
	log.Println("Connected to server")
}

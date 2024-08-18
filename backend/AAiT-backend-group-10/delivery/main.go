package main

import (
	"context"
	"log"
	"os"

	"aait.backend.g10/delivery/router"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	clientOption := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("DB_NAME")
	router.NewRouter(*db)

}

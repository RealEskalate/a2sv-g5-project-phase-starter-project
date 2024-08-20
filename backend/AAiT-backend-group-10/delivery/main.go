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
	err := godotenv.Load("/home/yordanos/Desktop/a2sv-g5-project-phase-starter-project/backend/AAiT-backend-group-10/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	clientOption := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(os.Getenv("DB_NAME"))
	router.NewRouter(db)

}

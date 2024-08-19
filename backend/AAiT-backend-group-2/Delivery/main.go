package main

import (
	"AAiT-backend-group-2/Delivery/routers"
	infrastructure "AAiT-backend-group-2/Infrastructure"

	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoConfig := infrastructure.NewMongoDBConfig(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE"))
	mongoClient, err := mongoConfig.Connect()	
	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(context.Background())

	db := mongoClient.Database(os.Getenv("MONGO_DATABASE"))

	r := gin.Default()

	routers.SetupRouter(db, r, os.Getenv("JWT_SECRET"))

	r.Run(":8080")

}
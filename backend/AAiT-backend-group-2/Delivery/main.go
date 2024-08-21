package main

import (
	"AAiT-backend-group-2/Delivery/routers"
	"AAiT-backend-group-2/Infrastructure"
	"AAiT-backend-group-2/config"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)



func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	mongoConfig := infrastructure.NewMongoDBConfig(configs.DBURI, configs.DbName)
	mongoClient, err := mongoConfig.Connect()	
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	defer mongoClient.Disconnect(context.TODO())

	db := mongoClient.Database(configs.DbName)

	r := gin.Default()

	routers.SetupRouter(db, r, &configs)

	r.Run(":8080")

}
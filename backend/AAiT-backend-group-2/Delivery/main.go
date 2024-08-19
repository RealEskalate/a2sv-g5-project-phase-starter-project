package main

import (
	"AAiT-backend-group-2/Delivery/controllers"
	"AAiT-backend-group-2/Delivery/routers"
	"AAiT-backend-group-2/Infrastructure"
	"AAiT-backend-group-2/Repositories"
	"AAiT-backend-group-2/Usecases"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoConfig := infrastructure.NewMongoDBConfig(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	mongoClient, err := mongoConfig.Connect()

	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	defer mongoClient.Disconnect(context.TODO())

	db := mongoClient.Database(os.Getenv("MONGO_DB"))
	blogRepo := repositories.NewBlogRepository(db)
	blogUseCase := usecases.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	router := routers.SetupRouter(blogController)

	router.Run(":8080")

}

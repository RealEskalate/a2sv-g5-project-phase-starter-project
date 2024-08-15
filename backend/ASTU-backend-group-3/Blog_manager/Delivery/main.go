package main

import (
	"ASTU-backend-group-3/Blog_manager/usecases"
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/repository"
	"ASTU-backend-group-3/Blog_manager/Delivery/routers"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	// connect to MongoDB
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userDatabase := client.Database("Blog management")

	userCollection := userDatabase.Collection("User")
	userRepository := repository.NewUserRepository(userCollection)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	router := routers.SetupRouter(userController)
	log.Fatal(router.Run(":8080"))




}
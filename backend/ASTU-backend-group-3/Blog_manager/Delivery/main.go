package main

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/Delivery/router"
	"ASTU-backend-group-3/Blog_manager/Repository"
	Usecases "ASTU-backend-group-3/Blog_manager/Usecases"

	// "ASTU-backend-group-3/Blog_manager/Delivery/router"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new MongoDB client and connect to the server
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	userDatabase := client.Database("Blog management")

	userCollection := userDatabase.Collection("User")
	userRepository := Repository.NewUserRepository(userCollection)
	userUsecase := Usecases.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	router := router.SetupRouter(userController)
	log.Fatal(router.Run(":8080"))

}

package main

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/Delivery/router"
	"ASTU-backend-group-3/Blog_manager/Repository"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"ASTU-backend-group-3/Blog_manager/infrastructure"

	// "ASTU-backend-group-3/Blog_manager/Delivery/router"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	userDatabase := client.Database("Blog_management")

	userCollection := userDatabase.Collection("User")
	blogCollection := userDatabase.Collection("Blog")

	tokenCollection := userDatabase.Collection("Token")
	userRepository := Repository.NewUserRepository(userCollection, tokenCollection)

	blogRepository := Repository.NewBlogRepository(blogCollection)
	blogUsecase := Usecases.NewBlogUsecase(blogRepository)

	// Initialize the Email Service
	emailService := infrastructure.NewEmailService()

	// Initialize the User Usecase with the User Repository and Email Service
	userUsecase := Usecases.NewUserUsecase(userRepository, emailService)
	blogController := controller.NewBlogController(blogUsecase, userUsecase)
	userController := controller.NewUserController(userUsecase)

	// Start the token cleanup cron job
	userController.StartTokenCleanupJob()

	router := router.SetupRouter(userController, blogController, tokenCollection)
	log.Fatal(router.Run(":8080"))

}

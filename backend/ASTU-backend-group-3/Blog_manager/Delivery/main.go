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

	emailService := infrastructure.NewEmailService()

	// Test sending an email
	if err := emailService.SendEmail("nebiyumusbah378@gmail.com", "Test Email", "<h1>This is a test email</h1>"); err != nil {
		log.Fatalf("Email sending test failed: %v", err)
	} else {
		log.Println("Email sending test passed!")
	}

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
	blogController := controller.NewBlogController(blogUsecase)

	// Initialize the Email Service
	emailService = infrastructure.NewEmailService()

	// Initialize the User Usecase with the User Repository and Email Service
	userUsecase := Usecases.NewUserUsecase(userRepository, emailService)
	userController := controller.NewUserController(userUsecase)

	router := router.SetupRouter(userController, blogController)
	log.Fatal(router.Run(":8080"))

}

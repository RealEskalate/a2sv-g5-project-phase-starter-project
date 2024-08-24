package main

import (
	"blog_project/delivery/controllers"
	"blog_project/delivery/routers"
	"blog_project/infrastructure"
	"blog_project/repositories"
	"blog_project/usecases"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	aiService := infrastructure.NewGeminiService(os.Getenv("GEMINI_API_KEY"))
	emailService := infrastructure.NewEmailService(
		"localhost",
		"1025",
		"",
		"",
	)

	db := client.Database("blog_project")
	userCollection := db.Collection("users")
	blogCollection := db.Collection("blogs")
	tokenCollection := db.Collection("tokens")
	redisCache := infrastructure.NewRedisCache("localhost:6379")
	blogRepo := repositories.NewBlogRepository(blogCollection, redisCache)
	userRepo := repositories.NewUserRepository(userCollection, redisCache)
	tokenRepo := repositories.NewTokenRepository(tokenCollection)
	userUsecase := usecases.NewUserUsecase(userRepo, blogRepo, emailService, tokenRepo)
	blogUsecase := usecases.NewBlogUsecase(aiService, blogRepo, userUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := routers.SetupRouter(blogController, userController)

	r.Run("localhost:8080")
}

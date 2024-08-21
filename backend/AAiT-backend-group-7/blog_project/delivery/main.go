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

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	aiService := infrastructure.NewGeminiService(os.Getenv("GEMINI_API_KEY"))
	db := client.Database("blog_project")
	userCollection := db.Collection("users")
	blogCollection := db.Collection("blogs")
	blogRepo := repositories.NewBlogRepository(blogCollection)
	userRepo := repositories.NewUserRepository(userCollection)
	userUsecase := usecases.NewUserUsecase(userRepo)
	blogUsecase := usecases.NewBlogUsecase(aiService ,blogRepo, userUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := routers.SetupRouter(blogController, userController)
	r.Run("localhost:8080")
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/router"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/repository"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	blogCollection := client.Database("blog_api_test").Collection("blogs")
	
	blogRepository := repository.NewBlogRepository(blogCollection , context.TODO())
	blogUsecase := usecases.NewBlogUseCase(blogRepository)
	blogController := controllers.NewBlogController(blogUsecase)
	r := router.SetupRouter(blogController)
	r.Run("localhost:3000")
}
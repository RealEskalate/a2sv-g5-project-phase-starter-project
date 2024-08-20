package main

import (
	"blog_project/delivery/controllers"
	"blog_project/delivery/routers"
	"blog_project/repositories"
	"blog_project/usecases"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("blog_project")
	userCollection := db.Collection("users")
	blogCollection := db.Collection("blogs")
	blogRepo := repositories.NewBlogRepository(blogCollection)
	userRepo := repositories.NewUserRepository(userCollection)
	userUsecase := usecases.NewUserUsecase(userRepo)
	blogUsecase := usecases.NewBlogUsecase(blogRepo, userUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := routers.SetupRouter(blogController, userController)
	r.Run("localhost:8080")
}

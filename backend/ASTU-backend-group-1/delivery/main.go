package main

import (
	"context"
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/delivery/controllers"
	"astu-backend-g1/delivery/routers"
	"astu-backend-g1/repository"
	"astu-backend-g1/usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("BlogAPI").Collection("Blogs")
	userCollections := client.Database("BlogAPI").Collection("Users")
	// auth := infrastructure.NewAuthMiddleware(userCollections)
	// auther := infrastructure.GeneralAuthorizer(auth)
	authController := infrastructure.NewAuthController(userCollections)
	_ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(blogCollections)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)
	userRepo := repository.NewUserRepository(userCollections)
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	UserController := controllers.NewUserController(userUsecase)
	Router := router.NewMainRouter(*UserController, *blogController, authController)
	Router.GinBlogRouter()

}

package main

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/delivery/controllers"
	router "astu-backend-g1/delivery/routers"
	"astu-backend-g1/repository"
	usecase "astu-backend-g1/usecases"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Legend:AmIHDrlwCqmxgjy2@pycluster.ajv4lb8.mongodb.net/?retryWrites=true&w=majority&appName=pycluster")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("BlogAPI").Collection("Blogs")
	userCollections := client.Database("BlogAPI").Collection("Users")
	// auth := infrastructure.NewAuthMiddleware(userCollections)
	// auther := infrastructure.GeneralAuthorizer(auth)
	authController := infrastructure.NewAuthController()
	_ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(blogCollections)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)
	userRepo := repository.NewUserRepository(userCollections)
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	UserController := controllers.NewUserController(userUsecase, userCollections)
	Router := router.NewMainRouter(*UserController, *blogController, authController)
	Router.GinBlogRouter()

}

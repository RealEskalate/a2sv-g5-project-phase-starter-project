package main

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/config"
	"astu-backend-g1/delivery/controllers"
	router "astu-backend-g1/delivery/routers"
	"astu-backend-g1/repository"
	usecase "astu-backend-g1/usecases"
	"context"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config,err := config.LoadConfig()
	clientOptions := options.Client().ApplyURI(config.Database.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("BlogAPI").Collection("Blogs")
	userCollections := client.Database("BlogAPI").Collection("Users")

	authController := infrastructure.NewAuthController()
	_ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(mongoifc.WrapCollection(blogCollections))
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)

	userRepo := repository.NewUserRepository(mongoifc.WrapCollection	(userCollections))
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	UserController := controllers.NewUserController(userUsecase)
	Router := router.NewMainRouter(*UserController, *blogController, authController)
	Router.GinBlogRouter()

}

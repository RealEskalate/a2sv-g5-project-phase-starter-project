package main

import (
	"AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	"AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Repository"
	usecase "AAiT-backend-group-8/Usecase"

	"context"
)

var SECRET_KEY = "123456abcd"

func main() {
	mongoClient := infrastructure.InitMongoDB("mongodb://localhost:27017")
	userCollection := mongoClient.Database("starterproject").Collection("users")
	tokenCollection := mongoClient.Database("starterproject").Collection("token")
	blogCollection := mongoClient.Database("starterproject").Collection("blogs")
	commentCollection := mongoClient.Database("starterproject").Collection("comments")

	userRepo := repository.NewUserRepository(userCollection, context.TODO())
	ts := infrastructure.NewTokenService(SECRET_KEY)
	ps := infrastructure.NewPasswordService()
	tr := repository.NewTokenRepository(tokenCollection, context.TODO())
	ms := infrastructure.NewMailService()

	blogRepo := repository.NewBlogRepository(blogCollection)
	blogUseCase := usecase.NewBlogUseCase(blogRepo)

	commentRepo := repository.NewCommentRepository(commentCollection, context.TODO())
	infra := infrastructure.NewInfrastructure()

	commentUseCase := usecase.NewCommentUseCase(commentRepo, infra)
	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)
	controller := Controller.NewController(blogUseCase, commentUseCase, userUseCase)

	//userHandler := Controller.NewUserHandler(userUseCase)

	r := Router.InitRouter(controller)

	r.Run(":8080")
}

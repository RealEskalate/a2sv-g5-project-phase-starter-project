package main

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
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
	likeCollection := mongoClient.Database("starterproject").Collection("likes")

	userRepo := repository.NewUserRepository(userCollection, context.TODO())
	ts := infrastructure.NewTokenService(SECRET_KEY)
	ps := infrastructure.NewPasswordService()
	tr := repository.NewTokenRepository(tokenCollection, context.TODO())
	ms := infrastructure.NewMailService()
	//	ts := infrastructure.NewTokenService(SECRET_KEY)
	infrastructure := infrastructure.NewInfrastructure()

	blogRepo := repository.NewBlogRepository(blogCollection)
	blogUseCase := usecase.NewBlogUseCase(blogRepo)

	commentRepo := repository.NewCommentRepository(commentCollection, context.TODO())

	commentUseCase := usecase.NewCommentUseCase(commentRepo, *infrastructure, ts)
	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)
	likeRepo := repository.NewLikeRepository(likeCollection, context.TODO())
	likeUseCase := usecase.NewLikeUseCase(*likeRepo, *infrastructure)

	controller := controller.NewController(commentUseCase, userUseCase, likeUseCase, blogUseCase)

	r := Router.InitRouter(controller)

	r.Run(":8080")
}

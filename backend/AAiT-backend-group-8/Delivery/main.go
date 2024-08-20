package main

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Repository"
	usecase "AAiT-backend-group-8/Usecase"

	"context"
)

var SECRET_KEY = "123456abcd"

func main() {
	mongoClient := infrastructure.InitMongoDB("mongodb://localhost:27017")

	user_collection := mongoClient.Database("starterproject").Collection("users")
	token_collection := mongoClient.Database("starterproject").Collection("token")
	comment_collection := mongoClient.Database("starterproject").Collection("comments")
	//like_collection := mongoClient.Database("starterproject").Collection("likes")

	userRepo := repository.NewUserRepository(user_collection, context.TODO())
	ts := infrastructure.NewTokenService(SECRET_KEY)
	ps := infrastructure.NewPasswordService()
	tr := repository.NewTokenRepository(token_collection, context.TODO())
	ms := infrastructure.NewMailService()
//	ts := infrastructure.NewTokenService(SECRET_KEY)

	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)

	commentRepo := repository.NewCommentRepository(*comment_collection, context.TODO())
	infrastructure := infrastructure.NewInfrastructure()
	commentUseCase := usecase.NewCommentUseCase(*commentRepo, *infrastructure , ts)
	controller := controller.NewController(*commentUseCase, userUseCase)

	r := Router.InitRouter(*&controller)

	r.Run(":8080")
}

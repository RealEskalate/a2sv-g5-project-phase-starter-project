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
	user_collection := mongoClient.Database("starterproject").Collection("users")
	token_collection := mongoClient.Database("starterproject").Collection("token")
	comment_collection := mongoClient.Database("starterproject").Collection("comments")
	commnetRepo := repository.NewCommentRepository(*comment_collection, context.TODO())
	infra := infrastructure.NewInfrastructure()
	commentUseCase := usecase.NewCommentUseCase(
		*commnetRepo,
		*infra,
	)

	controller := controllers.NewController(*commentUseCase)

	userRepo := repository.NewUserRepository(user_collection, context.TODO())
	ts := infrastructure.NewTokenService(SECRET_KEY)
	ps := infrastructure.NewPasswordService()
	tr := repository.NewTokenRepository(token_collection, context.TODO())
	ms := infrastructure.NewMailService()

	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)

	userHandler := controllers.NewUserHandler(userUseCase)
	r := Router.InitRouter(userHandler, *controller)

	r.Run(":8080")
}

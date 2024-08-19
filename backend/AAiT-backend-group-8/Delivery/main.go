package main

import (
	"AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	repository "AAiT-backend-group-8/Repository"
	"AAiT-backend-group-8/Usecase"
	"AAiT-backend-group-8/infrastructure"
	"context"
)

var SECRET_KEY = "123456abcd"

func main() {
	mongoClient := infrastructure.InitMongoDB("mongodb://localhost:27017")
	user_collection := mongoClient.Database("starterproject").Collection("users")
	token_collection := mongoClient.Database("starterproject").Collection("token")

	userRepo := repository.NewUserRepository(user_collection, context.TODO())
	ts := infrastructure.NewTokenService(SECRET_KEY)
	ps := infrastructure.NewPasswordService()
	tr := repository.NewTokenRepository(token_collection, context.TODO())
	ms:= infrastructure.NewMailService()

	userUseCase := Usecase.NewUserUseCase(userRepo, ts, ps, tr,ms)

	userHandler := Controller.NewUserHandler(userUseCase)
	r := Router.InitRouter(userHandler)

	r.Run(":8080")
}

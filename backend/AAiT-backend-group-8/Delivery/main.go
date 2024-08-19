package main

import (
	
	"AAiT-backend-group-8/Delivery/Controller"
	"AAiT-backend-group-8/infrastructure"
	"AAiT-backend-group-8/Repository"	
	"AAiT-backend-group-8/Usecase"
	"AAiT-backend-group-8/Delivery/Routes"
)

func main() {
	mongoClient := infrastructure.InitMongoDB("mongodb://localhost:27017")
	collection := mongoClient.Database("starterproject").Collection("users")

	userRepo := Repository.NewUserRepository(collection)
	userUseCase := Usecases.NewUserUseCase(userRepo)

	userHandler := Controller.NewUserHandler(userUseCase)
	r := Router.InitRouter(userHandler)

	r.Run(":8080")
}

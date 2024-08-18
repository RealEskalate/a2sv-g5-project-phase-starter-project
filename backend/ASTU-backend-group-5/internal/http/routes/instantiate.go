package routes

import (
	"blogApp/internal/http/handlers"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase"
	"blogApp/internal/usecase/user"

	"go.mongodb.org/mongo-driver/mongo"
)

func InstantaiteUserHandler(collection *mongo.Collection) *handlers.UserHandler {
	userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)
	return userHandler
}

func InstantaiteTokenUsecase(collection *mongo.Collection) *usecase.TokenUsecase {
	tokenRepo := mongodb.NewMongoTokenRepository(collection)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)
	return tokenUsecase
}

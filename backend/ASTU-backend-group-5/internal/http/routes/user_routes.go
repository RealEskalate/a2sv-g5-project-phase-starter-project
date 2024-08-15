package routes

import (
	"blogApp/internal/http/handlers"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRouters(collection *mongo.Collection, router *gin.Engine) {
	// userCollection := mongo.GetCollection("users")

	userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	loginUseCase := user.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(loginUseCase)

	router.POST("/auth/login", userHandler.Login)
	router.POST("/auth/register", userHandler.Register)

}

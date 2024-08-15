package routes

import (
	"blogApp/internal/http/handlers"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterVerificationRoutes(collection *mongo.Collection, router *gin.Engine) {

	userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	router.POST("/auth/verify/request", userHandler.RequestVerifyEmail)
	router.GET("/auth/confirm", userHandler.VerifyEmail) //I used this naming to make things clear

	router.POST("/auth/reset-password/request", userHandler.ResetPasswordRequest)
	router.POST("/auth/reset-password/confirm", userHandler.ResetPassword)
}

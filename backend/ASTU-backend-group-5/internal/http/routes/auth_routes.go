package routes

import (


	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterVerificationRoutes(collection *mongo.Collection, router *gin.Engine) {

	// userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	// userUsecase := user.NewUserUsecase(userRepo)
	userHandler := InstantaiteUserHandler(collection)
	authRoutes := router.Group("/api/v1/auth")

	{
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.POST("/verify/request", userHandler.RequestVerifyEmail)
		authRoutes.GET("/confirm", userHandler.VerifyEmail) //I used this naming to make things clear
		authRoutes.POST("/reset-password/request", userHandler.ResetPasswordRequest)
		authRoutes.POST("/reset-password/confirm", userHandler.ResetPassword)

	}
}

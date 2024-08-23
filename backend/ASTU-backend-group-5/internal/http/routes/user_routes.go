package routes

import (
	"blogApp/internal/http/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(collection *mongo.Collection, router *gin.Engine) {
	userHandler := InstantaiteUserHandler(collection)
	userRoute := router.Group("/api/v1/accounts", middleware.AuthMiddleware())
	{
		userRoute.GET("/me", userHandler.GetUser)
		userRoute.DELETE("/me", userHandler.DeleteUser)
		userRoute.PUT("/me", userHandler.UpdateUser)
		userRoute.GET("/any/:id", userHandler.GetAnyUser)
		userRoute.POST("upload", userHandler.UploadProfilePic)
	}

}

package routes

import (
	"blogApp/internal/http/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAdminUserRoutes(collection *mongo.Collection, router *gin.Engine) {
	userHandler := InstantaiteUserHandler(collection)

	adminRoute := router.Group("/api/v1/admin")
	adminRoute.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoute.PUT("/promote/:id",middleware.OwnerMiddleware(), userHandler.PromoteToAdmin)
		adminRoute.PUT("/demote/:id", userHandler.DemoteUser)
		adminRoute.GET("/users", userHandler.GetAllUsers)
		adminRoute.DELETE("/delete/:id", userHandler.AdminRemoveUser)
		adminRoute.GET("/users/:id", userHandler.GetUser)
		adminRoute.GET("/users/filtered", userHandler.FilterUsers)
	}

}
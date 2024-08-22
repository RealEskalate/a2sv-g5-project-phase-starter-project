package routes

import (
	"blogApp/internal/http/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAdminUserRoutes(userCollection *mongo.Collection, router *gin.Engine) {
	userHandler := InstantaiteUserHandler(userCollection)

	// tokenRepo := mongodb.NewMongoTokenRepository(tokenCollection)
	// tokenUseCase := usecase.NewTokenUsecase(tokenRepo)

	adminRoute := router.Group("/api/v1/admin")
	adminRoute.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoute.PUT("/user/promote/:id", middleware.OwnerMiddleware(), userHandler.PromoteToAdmin)
		adminRoute.PUT("/user/demote/:id", middleware.OwnerMiddleware(), userHandler.DemoteUser)
		adminRoute.GET("/users", userHandler.GetAllUsers)
		adminRoute.DELETE("/user/:id", userHandler.AdminRemoveUser)
		adminRoute.GET("/user/:id", userHandler.GetUser)
		adminRoute.GET("/users/filtered", userHandler.FilterUsers)
	}
}

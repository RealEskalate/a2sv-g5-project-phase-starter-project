package routers

import (
	"AAiT-backend-group-2/Delivery/controllers"
	domain "AAiT-backend-group-2/Domain"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	"AAiT-backend-group-2/Infrastructure/services"
	repositories "AAiT-backend-group-2/Repositories"
	usecases "AAiT-backend-group-2/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAIRouter(db *mongo.Database, group *gin.RouterGroup, configs *domain.Config) {
	aiService := services.NewAIService(configs.GeminiApiKey)
	aiRepo := repositories.NewAIRepository(db)
	aiUsecase := usecases.NewAIUsecase(aiRepo, aiService, 10*time.Second)
	aiController := controllers.NewAIController(aiUsecase)

	protectedRoute := group.Group("")
	protectedRoute.Use(infrastructure.AuthMiddleWare(configs.SecretKey))

	protectedRoute.POST("/chat", aiController.CreateChat)
	protectedRoute.GET("/chat/:id", aiController.GetChat)
}
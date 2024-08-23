package routers

import (
	"AAiT-backend-group-2/Delivery/controllers"
	domain "AAiT-backend-group-2/Domain"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	"AAiT-backend-group-2/Infrastructure/services"
	"AAiT-backend-group-2/Repositories/user_repository"
	"AAiT-backend-group-2/Usecases/user_usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(db *mongo.Database, group *gin.RouterGroup, configs *domain.Config) {
	jwtService := services.NewJWTService([]byte(configs.SecretKey))
	userRepo := user_repository.NewUserRepository(db)
	emailService := services.NewEmailService(configs.EmailHost, configs.EmailPort, configs.SenderEmail, configs.SenderPassword)
	imageService := services.NewImageService(configs.CloudinaryUrl)

	userUsecase := user_usecase.NewUserUsecase(userRepo, jwtService, emailService, imageService, 10*time.Second)

	authController := controllers.NewAuthController(userUsecase)

	group.POST("/register", authController.CreateUser)
	group.POST("/login", authController.Login)

	userRoutes := group.Group("")
	userRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey))
	userRoutes.POST("/refresh-token", authController.RefreshToken)
}
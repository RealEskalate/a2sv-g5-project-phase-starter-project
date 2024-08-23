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


func NewUserRouter(db *mongo.Database, group *gin.RouterGroup, configs *domain.Config) {
	
	jwtService := services.NewJWTService([]byte(configs.SecretKey))
	emailService := services.NewEmailService(configs.EmailHost, configs.EmailPort, configs.SenderEmail, configs.SenderPassword)
	imageService := services.NewImageService(configs.CloudinaryUrl)
	userRepo := user_repository.NewUserRepository(db)

	userUsecase := user_usecase.NewUserUsecase(userRepo, jwtService, emailService, imageService, 10*time.Second)

	userController := controllers.NewUserController(userUsecase)

	userRoutes := group.Group("")
	userRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey))
	{
		userRoutes.GET("users/:id", userController.GetUserByID)
		userRoutes.PUT("/users/:id", userController.UpdateUser)
		userRoutes.DELETE("/users/:id", userController.DeleteUser)
		userRoutes.POST("/users/forgot-password", userController.ForgotPassword)
		userRoutes.POST("/users/reset-password", userController.ResetPassword)
		userRoutes.PUT("/users/change-password", userController.ChangePassword)
		userRoutes.PUT("/users/update-profile", userController.UpdateProfile)
	}

	adminRoutes := group.Group("/admin")
	adminRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey), infrastructure.RoleMiddleware())

	{
		adminRoutes.GET("/users", userController.GetAllUsers)
		adminRoutes.PUT("/users/:id/promote", userController.PromoteUser)
		adminRoutes.PUT("/users/:id/demote", userController.DemoteAdmin)
	}
	
}
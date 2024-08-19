package routers

import (
	"AAiT-backend-group-2/Delivery/controllers"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	repositories "AAiT-backend-group-2/Repositories"
	usecases "AAiT-backend-group-2/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func NewUserRouter(db *mongo.Database, group *gin.RouterGroup, jwtSecret string) {
	
	jwtService := infrastructure.NewJWTService([]byte(jwtSecret))
	validatorService := infrastructure.NewValidatorService()
	userRepo := repositories.NewUserRepository(db, jwtService)

	userUsecase := usecases.NewUserUsecase(userRepo, 10*time.Second, validatorService)

	userController := controllers.NewUserController(userUsecase)

	group.GET("/users", userController.GetAllUsers)
	group.GET("users/:id", userController.GetUserByID)
	group.POST("/register", userController.CreateUser)
	group.PUT("/users/:id", userController.UpdateUser)
	group.POST("/login", userController.Login)
	group.DELETE("/users/:id", userController.DeleteUser)

	adminRoutes := group.Group("/admin")
	adminRoutes.Use(infrastructure.AuthMiddleWare(jwtSecret), infrastructure.RoleMiddleware())
	adminRoutes.PUT("/users/:id/promote", userController.PromoteUser)
	adminRoutes.PUT("/users/:id/demote", userController.DemoteAdmin)
}
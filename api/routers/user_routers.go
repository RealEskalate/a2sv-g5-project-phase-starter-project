package routers

import (
	"loan-management/api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sv-tools/mongoifc"
)

func AddUserRoutes(r *gin.Engine, db mongoifc.Database) {
	userController := controllers.NewUserController(db)
	userRouteGroup := r.Group("/users")
	{
		userRouteGroup.POST("/register", userController.SignUp)
		userRouteGroup.POST("/verify-email", userController.VerifyEmail)
		userRouteGroup.POST("/login", userController.Login)
	}
}

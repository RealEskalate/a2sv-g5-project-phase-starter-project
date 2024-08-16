package router

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/infrastructure"

	// "ASTU-backend-group-3/Blog_manager/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	
	
	usersRoute := router.Group("/")
	usersRoute.Use(infrastructure.AuthMiddleware()) // make sure to add Auth_User in the middleware
	usersRoute.PUT("/update/:username", userController.UpdateUser)
	usersRoute.POST("/logout", userController.Logout)

	protected := usersRoute.Group("/")
	protected.Use(infrastructure.AdminMiddleware()) // make sure to add Auth_User in the middleware
	protected.DELETE("/delete/:username", userController.DeleteUser)

	

	return router
}

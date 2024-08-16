package router

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	utils "ASTU-backend-group-3/Blog_manager/Infrastructure"

	// "ASTU-backend-group-3/Blog_manager/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)

	usersRoute := router.Group("/user")
	usersRoute.Use(utils.Auth_User()) // make sure to add Auth_User in the middleware
	usersRoute.PUT("/update/:username", userController.UpdateUser)

	protected := router.Group("/admin")
	protected.Use(utils.Auth_Admin()) // make sure to add Auth_User in the middleware
	protected.DELETE("/delete/:username", userController.DeleteUser)

	

	return router
}

package routers


import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	// "ASTU-backend-group-3/Blog_manager/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)


	// router.POST("/login", userController.Login)
	// router.GET("/users", userController.GetUsers)
	// router.GET("/users/:username", userController.GetUser)
	// router.PUT("/users/:username", userController.UpdateUser)
	// router.DELETE("/users/:username", userController.DeleteUser)

	return router
}
package router

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/infrastructure"

	// "ASTU-backend-group-3/Blog_manager/infrastructure"

	// "ASTU-backend-group-3/Blog_manager/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController, blogController *controller.BlogController) *gin.Engine {
	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/refresh", userController.RefreshToken)
	router.POST("/forgot-password", userController.ForgotPassword)
	router.GET("/reset/:token", userController.ResetPassword)

	usersRoute := router.Group("/")
	usersRoute.Use(infrastructure.AuthMiddleware()) // make sure to add Auth_User in the middleware
	usersRoute.PUT("/update/:username", userController.UpdateUser)
	usersRoute.PUT("/change_password", userController.ChangePassword)
	usersRoute.POST("/logout", userController.Logout)
	// usersRoute.POST("/logout", userController.Logout)

	usersRoute.POST("/blogs", blogController.CreateBlog)
	usersRoute.GET("/blogs", blogController.RetrieveBlogs)
	usersRoute.DELETE("/blogs/:id", blogController.DeleteBlogByID)
	usersRoute.GET("/blogs/search", blogController.SearchBlogs)
	usersRoute.PUT("/blogs/update/:id", blogController.UpdateBlog)
	// usersRoute := router.Group("/user")
	// usersRoute.Use(infrastructure.AuthMiddleware()) // make sure to add Auth_User in the middleware
	// usersRoute.PUT("/update/:username", userController.UpdateUser)
	protected := usersRoute.Group("/")
	// protected.Use(infrastructure.AdminMiddleware()) // make sure to add Auth_User in the middleware
	protected.DELETE("/delete/:username", userController.DeleteUser)

	// protected := usersRoute.Group("/")
	// protected.Use(infrastructure.RoleMiddleware("admin")) // make sure to add Auth_User in the middleware
	// protected.DELETE("/delete/:username", userController.DeleteUser)

	return router
}

package router

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(userController *controller.UserController, blogController *controller.BlogController, tokenCollection *mongo.Collection) *gin.Engine {
	router := gin.Default()

	// Public routes (no authentication required)
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/refresh", userController.RefreshToken)
	router.POST("/forgot-password", userController.ForgotPassword)
	router.GET("/reset/:token", userController.ResetPassword)
	router.GET("/verify/:token", userController.Verify)
	router.POST("/chat", controller.Chat)

	// Public blog routes
	router.GET("/blogs", blogController.RetrieveBlogs)

	// Authenticated user routes
	usersRoute := router.Group("/")
	usersRoute.Use(infrastructure.AuthMiddleware(tokenCollection)) // Apply authentication middleware

	// User management routes
	usersRoute.PUT("/update/:username", userController.UpdateUser)
	usersRoute.PUT("/change_password", userController.ChangePassword)
	usersRoute.POST("/logout", userController.Logout)

	// Blog management routes for authenticated users
	blogsRoute := usersRoute.Group("/blogs")
	blogsRoute.POST("/", blogController.CreateBlog)
	blogsRoute.GET("/search", blogController.SearchBlogs)
	blogsRoute.PUT("/update/:id", blogController.UpdateBlog)
	blogsRoute.DELETE("/delete/:id", blogController.DeleteBlogByID)
	blogsRoute.POST("/:id/like", blogController.ToggleLike)
	blogsRoute.POST("/:id/dislike", blogController.ToggleDislike)
	blogsRoute.POST("/:id/comment", blogController.AddComment)

	// Increment view count route
	blogsRoute.POST("/:id/view", blogController.IncrementViewCount)

	// Admin routes (requires admin role)
	adminRoute := usersRoute.Group("/")
	adminRoute.Use(infrastructure.RoleMiddleware("admin")) // Apply admin role middleware

	adminRoute.DELETE("/delete/:username", userController.DeleteUser)
	adminRoute.PUT("/promote/:username", userController.PromoteToAdmin)

	return router
}

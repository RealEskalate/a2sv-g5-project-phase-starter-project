package routers

import (
	"meleket/delivery/controllers"
	"meleket/infrastructure"
	"meleket/usecases"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, blogUsecase *usecases.BlogUsecase, userUsecase *usecases.UserUsecase, jwtService *infrastructure.JwtService) {
	// Initialize controllers
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Admin middleware
	adminMiddleware := infrastructure.AdminMiddleware(jwtService)

	// Public routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/forgot-password", userController.ForgotPassword)
	r.POST("/refresh-token", userController.RefreshToken)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{
		// User profile routes
		auth.GET("/profile", userController.GetProfile)
		auth.PUT("/profile", userController.UpdateProfile)
		auth.POST("/logout", userController.Logout)

		// Blog routes
		auth.POST("/blogs", blogController.CreateBlogPost)
		auth.GET("/blogs", blogController.GetAllBlogPosts)
		auth.GET("/blogs/:id", blogController.GetBlogByID)
		auth.PUT("/blogs/:id", blogController.UpdateBlogPost)
		auth.POST("/blogsearch", blogController.SearchBlogPost)
		auth.DELETE("/blogs/:id", blogController.DeleteBlogPost)

		// Admin-specific routes
		auth.POST("/getallusers", adminMiddleware, userController.GetAllUsers)
		auth.PUT("/deleteusers/:id", adminMiddleware, userController.DeleteUser)
	}
}

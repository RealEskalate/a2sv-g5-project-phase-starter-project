package router

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController domain.UserController, blogController domain.BlogController, blogAssistantController domain.BlogAssistantController, authMiddleware domain.MiddlewareService) *gin.Engine {
	r := gin.Default()

	// user related routes
	r.POST("/promote", authMiddleware.Authorize("admin"), userController.PromoteUser)
	r.POST("/demote", authMiddleware.Authorize("admin"), userController.DemoteUser)

	// Blog routes
	blogRoutes := r.Group("/blogs")
	{
		blogRoutes.Use(authMiddleware.Authenticate())

		blogRoutes.POST("/", blogController.CreateBlog)
		blogRoutes.GET("/:id", blogController.GetBlog)
		blogRoutes.GET("/", blogController.GetBlogs)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
		blogRoutes.GET("/search/title", blogController.SearchBlogsByTitle)
		blogRoutes.GET("/search/author", blogController.SearchBlogsByAuthor)
		blogRoutes.GET("/filter", blogController.FilterBlogs)
		blogRoutes.POST("/:id/like", blogController.LikeBlog)
		blogRoutes.POST("/:id/dislike", blogController.DislikeBlog)
		blogRoutes.POST("/:id/comments", blogController.AddComment)
		blogRoutes.DELETE("/:id/comments/:comment_id", blogController.DeleteComment)
		blogRoutes.PUT("/:id/comments/:comment_id", blogController.EditComment)

		// blog assistant related routes
		blogRoutes.POST("/generate", blogAssistantController.GenerateBlog)
		blogRoutes.POST("/enhance", blogAssistantController.EnhanceBlog)
		blogRoutes.GET("/suggest", blogAssistantController.SuggestBlog)
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.GET("/verify/:token", userController.VerifyEmail)
		userRoutes.POST("/login", userController.Login)
		userRoutes.POST("/forgot_password", userController.ForgotPassword)
		userRoutes.POST("/reset/:token", userController.ResetPassword)

		userRoutes.POST("/logout", userController.Logout).Use(authMiddleware.Authenticate())
		userRoutes.POST("/update/:id", userController.UpdateProfile).Use(authMiddleware.Authenticate())
		userRoutes.POST("/upload_profile_picture", userController.ImageUpload)
	}

		userRoutes.POST("/promote", userController.PromoteUser).Use(authMiddleware.Authenticate()).Use(authMiddleware.Authorize("admin"))
		userRoutes.POST("/demote", userController.DemoteUser).Use(authMiddleware.Authenticate()).Use(authMiddleware.Authorize("admin"))
	}

	return r
}

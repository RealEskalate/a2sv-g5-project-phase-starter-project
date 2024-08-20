package router

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/gin-gonic/gin"

)

func SetupRouter(userController domain.UserController, blogController domain.BlogController, blogAssistantController domain.BlogAssistantController, jwtService domain.JwtService) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/forgot-password", userController.ForgotPassword)
	r.POST("/logout", userController.Logout)
	r.PUT("/update-profile", userController.UpdateProfile)

	// Protected routes
	authMiddleware := infrastructure.NewMiddlewareService(jwtService)
	r.Use(authMiddleware.Authenticate())

	// user related routes
	r.POST("/promote", authMiddleware.Authorize("admin"), userController.PromoteUser)
	r.POST("/demote", authMiddleware.Authorize("admin"), userController.DemoteUser)

	// blog related routes
	r.POST("/post-blog", blogController.CreateBlog)
	r.GET("/get-blog/:id", blogController.GetBlog)
	r.GET("/get-blogs", blogController.GetBlogs)
	r.PUT("/edit-blog/:id", blogController.UpdateBlog)
	r.DELETE("/delete-blog/:id", blogController.DeleteBlog)
	r.GET("/search-blogs", blogController.SearchBlogs)
	r.GET("/filter-blogs", blogController.FilterBlogs)
	r.POST("/like-blog/:id", blogController.LikeBlog)
	r.POST("/add-comment", blogController.AddComment)
	r.DELETE("/delete-comment/:id", blogController.DeleteComment)
	r.PUT("/edit-comment/:id", blogController.EditComment)

	// blog assistant related routes
	r.POST("/generate-blog", blogAssistantController.GenerateBlog)
	r.POST("/enhance-blog", blogAssistantController.EnhanceBlog)
	r.GET("/suggest-blog", blogAssistantController.SuggestBlog)

	return r
}
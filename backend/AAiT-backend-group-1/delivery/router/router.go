package router

import (
	"log"
	"os"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter(userController domain.UserController, blogController domain.BlogController, blogAssistantController domain.BlogAssistantController, jwtService domain.JwtService) *gin.Engine {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUri := os.Getenv("DATABASE_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseService := infrastructure.NewDatabaseService(databaseUri, databaseName)
	
	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")
	jwtService = infrastructure.NewJWTTokenService(accessSecret, refreshSecret, databaseService.GetCollection("tokens"))
	
	cacheDbUri := os.Getenv("CACHE_DB_URI")
	cacheDbPassword := os.Getenv("CACHE_DB_PASSWORD")
	cacheService := infrastructure.NewCacheService(cacheDbUri, cacheDbPassword, 0)
	
	// Protected routes
	authMiddleware := infrastructure.NewMiddlewareService(jwtService, cacheService)
	r.Use(authMiddleware.Authenticate())

	// user related routes
	r.POST("/promote", authMiddleware.Authorize("admin"), userController.PromoteUser)
	r.POST("/demote", authMiddleware.Authorize("admin"), userController.DemoteUser)

	// Blog routes
	blogRoutes := r.Group("/blogs")
	{
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
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/", userController.Register)
		userRoutes.GET("/verify/:token", userController.VerifyEmail)
		userRoutes.POST("/login", userController.Login)
		userRoutes.POST("/forgot_password", userController.ForgotPassword)
		userRoutes.POST("/reset/:token", userController.ResetPassword)
		userRoutes.POST("/logout", userController.Logout).Use(authMiddleware.Authenticate())
		userRoutes.POST("/promote", userController.PromoteUser)
		userRoutes.POST("/demote", userController.DemoteUser)
		userRoutes.POST("/update/:id", userController.UpdateProfile).Use(authMiddleware.Authenticate())
		userRoutes.POST("/upload_profile_picture", userController.ImageUpload)
	}

	// blog assistant related routes
	r.POST("/generate-blog", blogAssistantController.GenerateBlog)
	r.POST("/enhance-blog", blogAssistantController.EnhanceBlog)
	r.GET("/suggest-blog", blogAssistantController.SuggestBlog)

	return r
}

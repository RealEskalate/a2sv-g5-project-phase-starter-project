package router

import (
	"context"
	"os"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database, redisClient *redis.Client) {
	router := gin.Default()

	//load email configuration from .env
	email := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	username := os.Getenv("EMAIL_USERNAME")
	host := os.Getenv("EMAIL_HOST")

	jwtService := infrastructures.JwtService{JwtSecret: os.Getenv("JWT_SECRET")}

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))
	cacheRepo := infrastructures.NewCacheRepo(redisClient, context.Background())

	pwdService := infrastructures.PwdService{}
	emailService := infrastructures.EmailService{
		AppEmail:    email,
		AppPass:     password,
		AppUsername: username,
		AppHost:     host,
	}
	aiService := infrastructures.NewAIService(os.Getenv("GEMINI_API_KEY"))

	blogRepo := repositories.NewBlogRepository(db, os.Getenv("BLOG_COLLECTION"))
	blogUseCase := usecases.NewBlogUseCase(blogRepo, userRepo, aiService)
	blogController := controllers.NewBlogController(blogUseCase)

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo, userRepo, cacheRepo),
	}

	likeRepo := repositories.NewLikeRepository(db, os.Getenv("LIKE_COLLECTION_NAME"))
	likeController := controllers.LikeController{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo, cacheRepo),
	}

	authUsecases := usecases.NewAuthUsecase(userRepo, jwtService, pwdService, emailService)
	authController := controllers.NewAuthController(authUsecases, controllers.GoogleOAuthConfig)

	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	router.POST("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.CreateBlog)
	router.GET("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.GetAllBlogs)
	router.GET("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.GetBlogByID)
	router.PUT("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.UpdateBlog)
	router.DELETE("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", infrastructures.AuthMiddleware(&jwtService), blogController.AddView)
	router.GET("/blogs/search", infrastructures.AuthMiddleware(&jwtService), blogController.SearchBlogs)
	router.POST("blogs/generate", infrastructures.AuthMiddleware(&jwtService), blogController.GenerateBlogContent)
	router.POST("blogs/suggest", infrastructures.AuthMiddleware(&jwtService), blogController.SuggestImprovements)

	router.PATCH("/users/promote", infrastructures.AuthMiddleware(&jwtService), infrastructures.AdminMiddleWare(), userController.PromoteUser)
	router.PUT("/users/:id", infrastructures.AuthMiddleware(&jwtService), userController.UpdateProfile)

	router.GET("/comment/:blog_id", infrastructures.AuthMiddleware(&jwtService), commentController.GetComments)
	router.POST("/comment", infrastructures.AuthMiddleware(&jwtService), commentController.AddComment)
	router.PUT("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.UpdateComment)
	router.DELETE("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.DeleteComment)

	router.PUT("/like", infrastructures.AuthMiddleware(&jwtService), likeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(&jwtService), likeController.DeleteLike)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/refresh-token", authController.RefreshToken)
	router.POST("/forgot-password", authController.ForgotPassword)
	router.POST("/reset-password", authController.ResetPassword)
	router.GET("/auth/google", authController.HandleGoogleLogin)
	router.GET("/auth/google/callback", authController.HandleGoogleCallback)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}

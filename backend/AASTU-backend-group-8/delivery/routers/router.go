package routers

import (
	"meleket/delivery/controllers"
	"meleket/infrastructure"
	"meleket/usecases"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, blogUsecase *usecases.BlogUsecase, userUsecase *usecases.UserUsecase, likeUsecase *usecases.LikeUsecase, commentUsecase *usecases.CommentUsecase, tokenUsecase *usecases.TokenUsecase, otpUsecase *usecases.OTPUsecase, jwtService infrastructure.JWTService) {
	// func InitRoutes(r *gin.Engine, blogUsecase *usecases.BlogUsecase, userUsecase *usecases.UserUsecase, tokenUsecase *usecases.TokenUsecase, otpUsecase *usecases.OTPUsecase) {

	// Initialize controllers
	signupController := controllers.NewSignupController(userUsecase, otpUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)
	// refreshTokenController := controllers.NewRefreshTokenController(userUsecase)
	commentController := controllers.NewCommentController(commentUsecase)
	likeController := controllers.NewLikeController(likeUsecase)

	// Admin middleware
	// adminMiddleware := infrastructure.AdminMiddleware(jwtService)
	adminMiddleware := infrastructure.AdminMiddleware(jwtService)

	// Public routes
	r.POST("/signup", signupController.Signup)
	r.POST("/login", userController.Login)
	r.POST("/verify", signupController.VerifyOTP)
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
		// auth.POST("/blogsearch", blogController.SearchBlogPost)
		auth.DELETE("/blogs/:id", blogController.DeleteBlogPost)

		// Comment routes
		auth.POST("/blogs/:id/comments", commentController.AddComment)
		auth.GET("/blogs/:id/comments", commentController.GetCommentsByBlogID)
		auth.PUT("/comments/:id", commentController.UpdateComment)
		auth.DELETE("/comments/:id", commentController.DeleteComment)

		// Like routes
		auth.POST("/blogs/:id/likes", likeController.AddLike)
		auth.GET("/blogs/:id/likes", likeController.GetLikesByBlogID)
		auth.DELETE("/likes/:id", likeController.RemoveLike)

		// Admin-specific routes
		auth.POST("/getallusers", adminMiddleware, userController.GetAllUsers)
		auth.PUT("/deleteusers/:id", adminMiddleware, userController.DeleteUser)

	}
}

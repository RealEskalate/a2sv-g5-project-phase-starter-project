package routers

import (
	"meleket/delivery/controllers"
	"meleket/infrastructure"
	"meleket/usecases"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, blogUsecase *usecases.BlogUsecase, userUsecase *usecases.UserUsecase, refreshTokenUsecase *usecases.TokenUsecase, otpUsecase *usecases.OTPUsecase, jwtService infrastructure.JWTService) {

	// Initialize controllers
	signupController := controllers.NewSignupController(userUsecase, otpUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)
	refreshTokenController := controllers.NewRefreshTokenController(userUsecase, refreshTokenUsecase)
	forgotPasswordController := controllers.NewForgotPasswordController()

	// Admin middleware
	// adminMiddleware := infrastructure.AdminMiddleware(jwtService)
	adminMiddleware := infrastructure.AdminMiddleware(jwtService)

	// Public routes
	r.POST("/signup", signupController.Signup)
	r.POST("/verify", signupController.VerifyOTP)
	r.POST("/login", userController.Login)
	r.POST("/refreshtoken", refreshTokenController.RefreshToken)
	r.POST("/forgotpassword", forgotPasswordController.ForgotPassword)
	r.POST("/verfiyforgotpassword", forgotPasswordController.VerifyForgotOTP)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{

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

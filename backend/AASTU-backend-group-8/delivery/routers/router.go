package routers

import (
	"meleket/delivery/controllers"
	"meleket/delivery/external"
	"meleket/infrastructure"
	"meleket/usecases"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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

	goth.UseProviders(
		google.New(os.Getenv("OAUTH_CLIENT_ID"), os.Getenv("OAUTH_CLIENT_SECRET"), os.Getenv("OAUTH_CALLBACK_URL")),
	)

	oauthHandler := external.NewOauthHandler(userUsecase)

	gothic.Store = sessions.NewCookieStore([]byte("secret"))
	r.GET("/auth/:provider", oauthHandler.SignInWithProvider)
	r.GET("/auth/:provider/callback", oauthHandler.CallbackHandler)
}

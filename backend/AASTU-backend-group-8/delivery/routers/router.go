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

func InitRoutes(r *gin.Engine, blogUsecase *usecases.BlogUsecase, userUsecase *usecases.UserUsecase, refreshTokenUsecase *usecases.TokenUsecase, jwtService infrastructure.JWTService, commentUsecase *usecases.CommentUsecase, tokenUsecase *usecases.TokenUsecase, otpUsecase *usecases.OTPUsecase) {
	r.MaxMultipartMemory = 8 << 20 // 8 MB
	r.Static("/public", "./uploads")

	// Initialize controllers
	signupController := controllers.NewSignupController(userUsecase, otpUsecase)
	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	refreshTokenController := controllers.NewRefreshTokenController(userUsecase, refreshTokenUsecase, jwtService)
	forgotPasswordController := controllers.NewForgotPasswordController(userUsecase, otpUsecase)
	logoutController := controllers.NewLogoutController(refreshTokenUsecase)

	commentController := controllers.NewCommentController(commentUsecase)
	// likeController := controllers.NewLikeController(likeUsecase)

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

	r.GET("/blogs/:id", blogController.GetBlogByID)
	r.GET("/blogs", blogController.GetAllBlogPosts)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(infrastructure.AuthMiddleware(jwtService))
	{
		// Blog routes
		auth.POST("/blogs", blogController.CreateBlogPost)
		auth.PUT("/blogs/:id", blogController.UpdateBlogPost)

		auth.POST("/logout", logoutController.Logout)

		// auth.POST("/blogsearch", blogController.SearchBlogPost)
		auth.DELETE("/blogs/:id", blogController.DeleteBlogPost)
		auth.POST("/blogs/:id/dislike", blogController.DislikeBlogPost)
		auth.POST("/blogs/:id/comment", blogController.AddComment)
		auth.POST("/blogs/:id/like", blogController.LikeBlogPost)
		// auth.GET("/blogs/:id/likes", likeController.GetLikesByBlogID)

		// Admin-specific routes
		auth.GET("/getallusers", adminMiddleware, userController.GetAllUsers)
		auth.DELETE("/deleteuser/:id", adminMiddleware, userController.DeleteUser)
	}

	goth.UseProviders(
		google.New(os.Getenv("OAUTH_CLIENT_ID"), os.Getenv("OAUTH_CLIENT_SECRET"), os.Getenv("OAUTH_CALLBACK_URL")),
	)

	oauthHandler := external.NewOauthHandler(userUsecase)

	gothic.Store = sessions.NewCookieStore([]byte("secret"))
	r.GET("/auth/:provider", oauthHandler.SignInWithProvider)
	r.GET("/auth/:provider/callback", oauthHandler.CallbackHandler)

	// Comment routes
	// auth.POST("/blogs/:id/comments", commentController.AddComment)
	auth.GET("/blogs/:id/comments", commentController.GetCommentsByBlogID)
	auth.PUT("/comments/:id", commentController.UpdateComment)
	auth.DELETE("/comments/:id", commentController.DeleteComment)

	// Like routes

	// auth.DELETE("/likes/:id", likeController.RemoveLike)

	// Admin-specific routes
	// auth.POST("/getallusers", adminMiddleware, userController.GetAllUsers)
	// auth.PUT("/deleteusers/:id", adminMiddleware, userController.DeleteUser)

	// }
}

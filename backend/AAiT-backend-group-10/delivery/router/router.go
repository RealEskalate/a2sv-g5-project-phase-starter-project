package router

import (
	"aait.backend.g10/delivery/config"
	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/infrastructures"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// Struct to accept all the RouterControllers
type RouterControllers struct {
	CommentController *controllers.CommentController
	LikeController    *controllers.LikeController
	AuthController    *controllers.AuthController
	UserController    *controllers.UserController
	BlogController    *controllers.BlogController
}

type RouterServices struct {
	JwtService *infrastructures.JwtService
}

func NewRouter(db *mongo.Database, redisClient *redis.Client, routerControllers RouterControllers, routerServices RouterServices) {
	router := gin.Default()

	jwtService := routerServices.JwtService

	router.POST("/blogs", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.CreateBlog)
	router.GET("/blogs", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.GetAllBlogs)
	router.GET("/blogs/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.GetBlogByID)
	router.PUT("/blogs/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.UpdateBlog)
	router.DELETE("/blogs/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.AddView)
	router.GET("/blogs/search", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.SearchBlogs)
	router.POST("blogs/generate", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.GenerateBlogContent)
	router.POST("blogs/suggest", infrastructures.AuthMiddleware(jwtService), routerControllers.BlogController.SuggestImprovements)

	router.GET("/users/:id", infrastructures.AuthMiddleware(jwtService), infrastructures.AdminMiddleWare(), routerControllers.UserController.GetUserByID)
	router.PATCH("/users/promote", infrastructures.AuthMiddleware(jwtService), infrastructures.AdminMiddleWare(), routerControllers.UserController.PromoteUser)
	router.PUT("/users/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.UserController.UpdateProfile)

	router.GET("/comment/:blog_id", infrastructures.AuthMiddleware(jwtService), routerControllers.CommentController.GetComments)
	router.POST("/comment", infrastructures.AuthMiddleware(jwtService), routerControllers.CommentController.AddComment)
	router.PUT("/comment/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.CommentController.UpdateComment)
	router.DELETE("/comment/:id", infrastructures.AuthMiddleware(jwtService), routerControllers.CommentController.DeleteComment)
	jwtService := infrastructures.Jwt{JwtSecret: os.Getenv("JWT_SECRET")}

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))

	pwdService := infrastructures.PwdService{}
	emailService := infrastructures.EmailService{}

	blogRepo := repositories.NewBlogRepository(db, os.Getenv("BLOG_COLLECTION"))
	blogUseCase := usecases.NewBlogUseCase(blogRepo, userRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo),
	}

	likeRepo := repositories.NewLikeRepository(db, os.Getenv("LIKE_COLLECTION_NAME"))
	likeController := controllers.LikeController{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo),
	}

	authUsecases := usecases.NewAuthUsecase(userRepo, jwtService, pwdService, emailService)
	authController := controllers.NewAuthController(authUsecases, controllers.GoogleOAuthConfig)

	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	router.POST("/blogs", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.CreateBlog)
	router.GET("/blogs", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.GetAllBlogs)
	router.GET("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.GetBlogByID)
	router.PUT("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.UpdateBlog)
	router.DELETE("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.AddView)
	router.GET("/blogs/search", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, blogController.SearchBlogs)

	router.PATCH("/users/promote", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, infrastructures.AdminMiddleWare(), userController.PromoteUser)

	router.GET("/comment/:blog_id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, commentController.GetComments)
	router.GET("/comment_count/:blog_id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, commentController.GetCommentsCount)
	router.POST("/comment", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, commentController.AddComment)
	router.PUT("/comment/:id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, commentController.UpdateComment)
	router.DELETE("/comment/:id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, commentController.DelelteComment)

	router.PUT("/like", infrastructures.AuthMiddleware(jwtService), routerControllers.LikeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(jwtService), routerControllers.LikeController.DeleteLike)
	router.PUT("/like", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, likeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, likeController.DeleteLike)
	router.GET("/like/:blog_id", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, likeController.BlogLikeCount)

	router.POST("/register", routerControllers.AuthController.Register)
	router.POST("/login", routerControllers.AuthController.Login)
	router.POST("/refresh-token", routerControllers.AuthController.RefreshToken)
	router.POST("/forgot-password", routerControllers.AuthController.ForgotPassword)
	router.POST("/reset-password", routerControllers.AuthController.ResetPassword)
	router.GET("/auth/google", routerControllers.AuthController.HandleGoogleLogin)
	router.GET("/auth/google/callback", routerControllers.AuthController.HandleGoogleCallback)
	router.POST("/upload-image", infrastructures.AuthMiddleware(&jwtService), authController.VerifyUserAccessToken, userController.UploadProfilePic)

	router.Run(":" + config.ENV.PORT)
}

package router

import (
	"os"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) {
	router := gin.Default()

	jwtService := infrastructures.JwtService{JwtSecret: os.Getenv("JWT_SECRET")}

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))

	jwtService := infrastructures.Jwt{JwtSecret: os.Getenv("jwtSecret")}
	pwdService := infrastructures.PwdService{}
	emailService := infrastructures.EmailService{}

	UserRepo := repositories.NewUserRepositoryMongo(db, os.Getenv("USER_COLLECTION"))

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

	authUsecases := usecases.NewAuthUsecase(UserRepo, jwtService, pwdService, emailService)
	authController := controllers.NewAuthController(authUsecases)

	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	router.POST("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.CreateBlog)
	router.GET("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.GetAllBlogs)
	router.GET("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.GetBlogByID)
	router.PUT("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.UpdateBlog)
	router.DELETE("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", infrastructures.AuthMiddleware(&jwtService), blogController.AddView)
	router.GET("/blogs/search", infrastructures.AuthMiddleware(&jwtService), blogController.SearchBlogs)

	router.PATCH("/users/promote", infrastructures.AuthMiddleware(&jwtService), infrastructures.AdminMiddleWare(), userController.PromoteUser)

	router.GET("/comment/:blog_id", infrastructures.AuthMiddleware(&jwtService), commentController.GetComments)
	router.GET("/comment_count/:blog_id", infrastructures.AuthMiddleware(&jwtService), commentController.GetCommentsCount)
	router.POST("/comment", infrastructures.AuthMiddleware(&jwtService), commentController.AddComment)
	router.PUT("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.UpdateComment)
	router.DELETE("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.DelelteComment)

	router.PUT("/like", infrastructures.AuthMiddleware(&jwtService), likeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(&jwtService), likeController.DeleteLike)
	router.GET("/like/:blog_id", infrastructures.AuthMiddleware(&jwtService), likeController.BlogLikeCount)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/refresh", authController.RefreshToken)
	router.POST("/refresh-token", authController.RefreshToken)
	router.POST("/forgot-password", authController.ForgotPassword)
	router.POST("/reset-password", authController.ResetPassword)
	
	port := os.Getenv("PORT")
	router.Run(":" + port)
}

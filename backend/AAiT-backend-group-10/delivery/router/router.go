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

	router.PUT("/like", infrastructures.AuthMiddleware(jwtService), routerControllers.LikeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(jwtService), routerControllers.LikeController.DeleteLike)

	router.POST("/register", routerControllers.AuthController.Register)
	router.POST("/login", routerControllers.AuthController.Login)
	router.POST("/refresh-token", routerControllers.AuthController.RefreshToken)
	router.POST("/forgot-password", routerControllers.AuthController.ForgotPassword)
	router.POST("/reset-password", routerControllers.AuthController.ResetPassword)
	router.GET("/auth/google", routerControllers.AuthController.HandleGoogleLogin)
	router.GET("/auth/google/callback", routerControllers.AuthController.HandleGoogleCallback)
	router.GET("/auth/activate/:token", routerControllers.AuthController.ActivateUser)
	router.POST("/logout", infrastructures.AuthMiddleware(jwtService), routerControllers.AuthController.LogoutUser)

	router.Run(":" + config.ENV.PORT)
}

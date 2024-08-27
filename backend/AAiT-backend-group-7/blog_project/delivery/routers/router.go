package routers

import (
	"blog_project/domain"
	"blog_project/infrastructure"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(blogController domain.IBlogController, userController domain.IUserController) *gin.Engine {
	r := gin.Default()

	// Blog routes
	blogs := r.Group("/blogs")
	blogs.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		blogs.GET("/", blogController.GetAllBlogs)
		blogs.POST("/", blogController.CreateBlog)
		blogs.PUT("/:id", blogController.UpdateBlog)
		blogs.DELETE("/:id", blogController.DeleteBlog)

		blogs.POST("/comment/:id", blogController.AddComment)
		blogs.POST("/like/:id", blogController.LikeBlog)
		blogs.POST("/dislike/:id", blogController.DislikeBlog)
		blogs.POST("/search", blogController.Search)
		blogs.POST("/generate-content", blogController.AiRecommendation)
	}

	// User routes
	users := r.Group("/users")
	users.POST("/", userController.CreateUser)
	users.POST("/login", userController.Login)
	users.POST("/forget-password/", userController.ForgetPassword)
	users.POST("/reset-password/:username/:password", userController.ResetPassword)
	users.POST("/logout", userController.Logout)
	users.POST("/refresh-token", userController.RefreshToken)

	users.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		users.GET("/", infrastructure.AdminMiddleware(), userController.GetAllUsers)
		users.GET("/:id", infrastructure.AdminMiddleware(), userController.GetUserByID)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
		users.POST("/promote/:id", infrastructure.AdminMiddleware(), userController.PromoteUser)
		users.POST("/demote/:id", infrastructure.AdminMiddleware(), userController.DemoteUser)
	}

	// Upload routes
	uploads := r.Group("/uploads")
	uploads.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		uploads.Static("/", "./uploads")
	}

	return r
}

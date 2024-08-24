package routers

import (
	"blog_project/domain"
	"os"

	"blog_project/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(blogController domain.IBlogController, userController domain.IUserController) *gin.Engine {

	r := gin.Default()

	blogs := r.Group("/blogs")

	blogs.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		// http://localhost:8080/blogs?sort=DESC&page=1&limit=10
		blogs.GET("/", blogController.GetAllBlogs)
		blogs.POST("/", blogController.CreateBlog)
		blogs.PUT("/:id", blogController.UpdateBlog)
		blogs.DELETE("/:id", blogController.DeleteBlog)

		blogs.POST("/:blog_id/:author_id/comment", blogController.AddComment)
		blogs.POST("/:blog_id/:author_id/like", blogController.LikeBlog)
		blogs.POST("/:blog_id/:author_id/dislike", blogController.DislikeBlog)
		blogs.POST("/search", blogController.Search)
		blogs.POST("/GenerateContent", blogController.AiRecommendation)
	}

	users := r.Group("/users")
	users.POST("/", userController.CreateUser)
	users.POST("/login", userController.Login)
	users.POST("/forget-password/:email", userController.ForgetPassword)
	users.POST("/reset-password/:username/:password", userController.ResetPassword)
	users.POST("/logout", userController.Logout)
	users.POST("/refresh-token", userController.RefreshToken)

	users.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		users.GET("/", infrastructure.AdminMiddleware(), userController.GetAllUsers)
		users.GET("/:id", infrastructure.AdminMiddleware(), userController.GetUserByID)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
		// users.POST("/:userID/blog", userController.AddBlog)
		users.POST("/promote/:id", infrastructure.AdminMiddleware(), userController.PromoteUser)
		users.POST("/demote/:id", infrastructure.AdminMiddleware(), userController.DemoteUser)
	}

	uploads := r.Group("/uploads")
	uploads.Use(infrastructure.JwtAuthMiddleware(os.Getenv("jwt_secret")))
	{
		uploads.Static("/", "./uploads")
	}

	return r

}

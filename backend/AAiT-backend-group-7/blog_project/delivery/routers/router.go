package routers

import (
	"blog_project/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(blogController domain.IBlogController, userController domain.IUserController) *gin.Engine {

	r := gin.Default()

	blogs := r.Group("/blogs")
	blogs.GET("/", blogController.GetAllBlogs)
	blogs.POST("/", blogController.CreateBlog)
	blogs.PUT("/:id", blogController.UpdateBlog)
	blogs.DELETE("/:id", blogController.DeleteBlog)

	blogs.POST("/:blog_id/:author_id/comment", blogController.AddComment)
	blogs.POST("/:blog_id/:author_id/like", blogController.LikeBlog)
	blogs.POST("/:blog_id/:author_id/dislike", blogController.DislikeBlog)
	blogs.POST("/search", blogController.Search)

	users := r.Group("/users")

	users.GET("/", userController.GetAllUsers)
	users.GET("/:id", userController.GetUserByID)
	users.POST("/", userController.CreateUser)
	users.PUT("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)
	// users.POST("/:userID/blog", userController.AddBlog)
	users.POST("/login", userController.Login)
	users.POST("/forget-password/:email", userController.ForgetPassword)
	users.POST("/reset-password/:username/:password", userController.ResetPassword)
	users.POST("/promote/:id", userController.PromoteUser)
	users.POST("/demote/:id", userController.DemoteUser)
	users.POST("/refresh-token", userController.RefreshToken)

	return r

}

package router

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouter(router *gin.Engine, com *controllers.CommentController, c *controllers.BlogController, cu *controllers.UserController, client *mongo.Client) {
	router.POST("/blog", c.CreateBlog)
	router.GET("/blog", c.RetrieveBlog)
	router.PUT("/blog/:id", c.UpdateBlog)
	router.DELETE("/blog/:id", c.DeleteBlog)
	router.GET("/blog/search", c.SearchBlog)
	router.GET("/blog/filter", c.FilterBlog)

	router.POST("/blog/:blog_id/comment", com.CreateComment)
	router.GET("/blog/:blog_id/comment", com.GetComment)
	router.PUT("/blog/:blog_id/comment/:id", com.UpdateComment)
	router.DELETE("/blog/:blog_id/comment/:id", com.DeleteComment)

	router.POST("/user/register", cu.RegisterUser)
	router.POST("/user/login", cu.LoginUser)

	router.GET("/logout", middleware.AuthMiddleware(client), cu.LogoutUser)

	router.POST("/forget-password", cu.ForgotPassword)
	router.POST("/reset-password", cu.ResetPassword)

}

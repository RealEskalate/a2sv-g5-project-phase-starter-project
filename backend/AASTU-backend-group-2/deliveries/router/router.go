package router

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetRouter(router *gin.Engine, c *controllers.BlogController, u *controllers.UserController, client *mongo.Client) {
	router.POST("/blog", c.CreateBlog)
	router.GET("/blog", c.RetrieveBlog)
	router.PUT("/blog/:id", c.UpdateBlog)
	router.DELETE("/blog/:id", c.DeleteBlog)
	router.GET("/blog/search", c.SearchBlog)
	router.GET("/blog/filter", c.FilterBlog)

	router.POST("/user/register", u.RegisterUser)
	router.POST("/user/login", u.LoginUser)

	router.GET("/logout", middleware.AuthMiddleware(client), u.LogoutUser)

}

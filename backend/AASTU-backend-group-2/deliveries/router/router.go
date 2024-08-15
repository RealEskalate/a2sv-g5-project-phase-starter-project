package router

import (
	"blog_g2/deliveries/controllers"

	"github.com/gin-gonic/gin"
)

func SetRouter(c *controllers.BlogController) {

	router := gin.Default()
	router.POST("/blog", c.CreateBlog)
	router.GET("/blog", c.RetrieveBlog)
	router.PUT("/blog/:id", c.UpdateBlog)
	router.DELETE("/blog/:id", c.DeleteBlog)
	router.GET("/blog/search", c.SearchBlog)
	router.GET("/blog/filter", c.FilterBlog)

	router.Run("8080")
}

package routers

import (
	"AAiT-backend-group-2/Delivery/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(blogController *controllers.BlogController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// router.Use(CORSMiddleware())

	router.POST("/blogs", blogController.CreateBlog)
	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.POST("/blogs", blogController.Search)
	return router
}

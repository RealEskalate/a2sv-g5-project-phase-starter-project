package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"github.com/gin-gonic/gin"
)

func CreateBlogRouter(router *gin.Engine, blogController *controllers.BlogController){
	router.POST("/blogs", blogController.CreateBlog)
	router.PATCH("/blogs/id", blogController.EditBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.GET("/blogs", blogController.GetBlogs)
}
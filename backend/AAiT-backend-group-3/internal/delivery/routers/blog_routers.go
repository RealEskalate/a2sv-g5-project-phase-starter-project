package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateBlogRouter(router *gin.Engine, blogController controllers.BlogControllerInterface, authMiddleware middlewares.IAuthMiddleware) {
	router.POST("/blogs", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.CreateBlog)
	
	router.PATCH("/blogs/:id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.UpdateBlog)
	
	router.DELETE("/blogs/:id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.DeleteBlog)
	
	router.GET("/blogs/:id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogByID)
	
	router.GET("/blogs", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogs)
	
	router.GET("/blogs/author/:author_id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogsByAuthorID)
	
	router.GET("/blogs/popularity", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogsByPopularity)
	
	router.GET("/blogs/tags", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogsByTags)
	router.POST("/blogs/:id/like", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.LikeBlog)
	router.POST("/blogs/:id/view", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.ViewBlog)
}


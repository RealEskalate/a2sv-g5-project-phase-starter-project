package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateBlogRouter(router *gin.Engine, blogController controllers.BlogControllerInterface, authMiddleware middlewares.IAuthMiddleware) {
	router.POST("/blogs", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN"), blogController.CreateBlog)
	router.PATCH("/blogs/id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.UpdateBlog)
	router.DELETE("/blogs/:id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.DeleteBlog)
	router.GET("/blogs/:id", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogByID)
	router.GET("/blogs", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), blogController.GetBlogs)
}

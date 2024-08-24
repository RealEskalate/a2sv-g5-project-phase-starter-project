package routers

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"
	"github.com/gin-gonic/gin"
)

func CreateCommentRouter(router *gin.Engine, commentController controllers.CommentControllerInterface, authMiddleware middlewares.IAuthMiddleware) {

	router.POST("/blog/:blogID/comments", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.CreateComment)
	router.GET("/comments/:commentID", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.GetCommentByID)
	router.PATCH("/comments/:commentID", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.EditComment)
	router.DELETE("/comments/:commentID", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.DeleteComment)
	router.POST("/comments/list", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.GetCommentsByIDList)
	router.GET("/comments/author/:authorID", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.GetCommentByAuthorID)
	router.GET("/comments/blog/:blogID", authMiddleware.Authentication(), authMiddleware.RoleAuth("ADMIN", "USER"), commentController.GetCommentByBlogID)
}

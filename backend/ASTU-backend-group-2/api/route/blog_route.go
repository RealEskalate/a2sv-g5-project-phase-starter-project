package route

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/gin-gonic/gin"
)

func BlogHandlers(r *gin.Engine, ctrl controller.BlogController) {

	// anyone can access
	r.GET("/blogs", ctrl.GetBlogs())
	r.GET("/blogs/:id", ctrl.GetBlog())

	// only authenticated users can access
	r.Use(middleware.UserMiddleware())
	r.POST("/blogs", ctrl.CreateBlog())
	r.PUT("/blogs/:id", ctrl.UpdateBlog())
	r.PATCH("/blogs/:id", ctrl.UpdateBlog())
	r.DELETE("/blogs/:id", ctrl.DeleteBlog())

	// comments
	r.GET("/blogs/:id/comments", ctrl.GetComments())
	r.POST("/blogs/:id/comments", ctrl.CreateComment())

	// only authenticated users can access
	r.GET("/blogs/:id/comments/:comment_id", ctrl.GetComment())
	r.PUT("/blogs/:id/comments/:comment_id", ctrl.UpdateComment())
	r.PATCH("/blogs/:id/comments/:comment_id", ctrl.UpdateComment())
	r.DELETE("/blogs/:id/comments/:comment_id", ctrl.DeleteComment())

	r.POST("/blogs/:id/like", ctrl.CreateLike())

}

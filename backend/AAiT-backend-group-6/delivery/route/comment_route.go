package route

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/delivery/middleware"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"

	"github.com/gin-gonic/gin"
)

func NewCommentRouter(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	tr := repository.NewCommentRepository(db, domain.CollectionComments)
	br := repository.NewBlogRepository(db, domain.CollectionBlogs)
	ur := repository.NewUserRepository(db, domain.UserCollection)
	cu := usecase.NewCommentUseCase(tr, br, ur)
	cc := controller.CommentController{
		CommentUseCase: cu,
	}
	protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	protectedRoute.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	publicRoute.GET("/comments/:comment_id", cc.GetComment)
	protectedRoute.POST("/comments/:blog_id", cc.CreateComment)
	publicRoute.PUT("/comments/:comment_id", cc.UpdateComment)
	publicRoute.DELETE("/comments/:comment_id", cc.DeleteComment)

}

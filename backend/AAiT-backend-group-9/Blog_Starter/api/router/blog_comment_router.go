package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogCommentRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	cr := repository.NewCommentRepository(db, domain.CollectionComment)
	br := repository.NewBlogRepository(db, domain.CollectionBlog)
	bcu := usecase.NewCommentUseCase(cr, br, 100*time.Second)
	cc := controller.NewBlogCommentController(bcu, timeout)

	group.POST("/comment", cc.CreateComment)
	group.PUT("/comment", cc.UpdateComment)
	group.DELETE("/comment/:id", cc.DeleteCommment)
}
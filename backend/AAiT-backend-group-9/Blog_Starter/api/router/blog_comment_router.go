package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogCommentRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cr := repository.NewCommentRepository(db, domain.CollectionComment, &ctx)
	br := repository.NewBlogRepository(db, domain.CollectionBlog, &ctx)
	bcu := usecase.NewCommentUseCase(cr, br)
	cc := controller.NewBlogCommentController(bcu, ctx)

	group.POST("/comment", cc.CreateComment)
	group.PUT("/comment", cc.UpdateComment)
	group.DELETE("/comment/:id", cc.DeleteCommment)
}
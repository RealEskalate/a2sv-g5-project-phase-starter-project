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

func NewBlogRatingRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	bra := repository.NewBlogRatingRepository(db, domain.CollectionRating, &ctx)
	br := repository.NewBlogRepository(db, domain.CollectionBlog, &ctx)
	bru := usecase.NewBlogRatingUseCase(bra, br)
	brc := controller.NewBlogRatingController(bru, ctx)

	group.POST("/rating", brc.InserttAndUpdateRating)
	group.DELETE("/rating", brc.DeleteRating)
}
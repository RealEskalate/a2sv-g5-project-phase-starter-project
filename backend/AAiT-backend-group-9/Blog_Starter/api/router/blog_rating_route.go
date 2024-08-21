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

func NewBlogRatingRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	bra := repository.NewBlogRatingRepository(db, domain.CollectionRating)
	br := repository.NewBlogRepository(db, domain.CollectionBlog)
	bru := usecase.NewBlogRatingUseCase(bra, br, 100*time.Second)
	brc := controller.NewBlogRatingController(bru, timeout)

	group.POST("/rating/:blog_id", brc.InserttAndUpdateRating)
	group.DELETE("/rating/:rating_id", brc.DeleteRating)
}
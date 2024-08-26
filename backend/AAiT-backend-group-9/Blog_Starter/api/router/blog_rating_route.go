package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"Blog_Starter/config"
)

func NewBlogRatingRouter(env *config.Env , timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
	database := db.Database(env.DBName) // Replace with your actual database name

	bra := repository.NewBlogRatingRepository(database, domain.CollectionRating)
	br := repository.NewBlogRepository(database, domain.CollectionBlog)
	bru := usecase.NewBlogRatingUseCase(bra, br, timeout)
	brc := controller.NewBlogRatingController(bru, timeout)

	group.POST("/rating", brc.InsertRating)
	group.PATCH("/rating/:id", brc.UpdateRating)
	group.DELETE("/rating/:id", brc.DeleteRating)
}
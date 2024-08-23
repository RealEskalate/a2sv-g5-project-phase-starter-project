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

func NewBlogLikeRouter(env *config.Env , timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
	database := db.Database(env.DBName) // Replace with your actual database name

	lr := repository.NewLikeRepository(database, domain.CollectionLike)
	br := repository.NewBlogRepository(database, domain.CollectionBlog)
	lu := usecase.NewLikeUseCase(lr, br, timeout)
	lc := controller.NewLikeController(lu, timeout)

	group.POST("/like", lc.LikeBlog)
	group.DELETE("/like/:id", lc.UnlikeBlog)
	group.GET("/like/:id", lc.GetByID)
}
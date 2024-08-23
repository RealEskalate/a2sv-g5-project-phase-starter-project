package router

import (
	"Blog_Starter/api/controller"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"time"
	"Blog_Starter/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogCommentRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {

	database := db.Database(env.DBName) // Replace with your actual database name

	cr := repository.NewCommentRepository(database, domain.CollectionComment)
	br := repository.NewBlogRepository(database, domain.CollectionBlog)
	bcu := usecase.NewCommentUseCase(cr, br, timeout)
	cc := controller.NewBlogCommentController(bcu, timeout)

	group.POST("/comment", cc.CreateComment)
	group.PATCH("/comment/:id", cc.UpdateComment)
	group.DELETE("/comment/:id", cc.DeleteCommment)
}
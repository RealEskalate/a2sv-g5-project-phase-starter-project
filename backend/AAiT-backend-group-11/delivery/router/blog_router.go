package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup)  {
	br := repository.NewBlogRepository(db.Collection("blogs"), context.TODO())
	bs := service.NewBlogService(br)

	ac := controller.NewBlogController(bs)

	group.POST("/blogs", ac.CreateBlogPost)
	group.GET("/blogs", ac.GetBlogPosts)
	group.GET("/blogs/:id", ac.GetBlogPost)
	group.PUT("/blogs/:id", ac.UpdateBlogPost)
	group.DELETE("/blogs/:id", ac.DeleteBlogPost)
}

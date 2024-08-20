package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommmentRouter(db *mongo.Database,  group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db.Collection("users"))
	br := repository.NewBlogRepository(db.Collection("blogs"),context.TODO())
	cr := repository.NewCommentRepository(db.Collection("comments"), context.TODO())
	cs := service.NewCommentService(cr,br,ur)
	cc := controller.NewCommentController(cs)
	

	
		group.POST("/:blogId",cc.AddComment)
		group.GET("/:blogId", cc.GetCommentsByBlogPostId)
		group.PUT("/:id", cc.UpdateComment)
		group.DELETE("/:id",  cc.DeleteComment)

}
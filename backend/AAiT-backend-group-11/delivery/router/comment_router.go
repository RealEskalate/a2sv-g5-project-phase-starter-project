package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/mongo"
	"context"

	"github.com/gin-gonic/gin"
)

func NewCommmentRouter(db *mongo.Database,  group *gin.RouterGroup) {
	blogcollection := (*db).Collection("blogs")
	usercollection := (*db).Collection("users")
	commentcollection := (*db).Collection("comments")

	ur := repository.NewUserRepository(usercollection)
	br := repository.NewBlogRepository(&blogcollection,context.TODO())
	cr := repository.NewCommentRepository(&commentcollection, context.TODO())
	cs := service.NewCommentService(cr,br,ur)
	cc := controller.NewCommentController(cs)
	

	
		group.POST(":blogId",cc.AddComment)
		group.GET(":blogId", cc.GetCommentsByBlogPostId)
		group.PUT(":id", cc.UpdateComment)
		group.DELETE(":id",  cc.DeleteComment)

}
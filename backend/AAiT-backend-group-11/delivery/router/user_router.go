package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(db *mongo.Database,group *gin.RouterGroup){
	ur := repository.NewUserRepository(db.Collection("users"))
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)

	group.POST("promote/:id", uc.PromoteUser)
	group.POST("demote/:id", uc.DemoteUser)
}
package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/mongo"

	"github.com/gin-gonic/gin"

)

func NewUserRouter(db *mongo.Database,group *gin.RouterGroup){
	userCollection := (*db).Collection("users")
	ur := repository.NewUserRepository(userCollection)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)

	group.POST("promote/:id", uc.PromoteUser)
	group.POST("demote/:id", uc.DemoteUser)
}
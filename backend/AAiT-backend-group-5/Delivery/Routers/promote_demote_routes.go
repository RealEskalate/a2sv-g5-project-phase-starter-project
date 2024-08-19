package routers

import (
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPromoteDemoteRouter(database mongo.Database, group *gin.RouterGroup) {

	user_repo := repository.NewUserRepository(&database)

	// instantiate PromoteDemote controller
	PromteDemoteController := &controllers.PromoteDemoteController{
		PromoteDemoteUC: usecases.NewUserUsecase(user_repo),
	}

	group.POST("/promoteUser/:id", PromteDemoteController.PromoteUser)
	group.POST("/demoteUser/:id", PromteDemoteController.DemoteUser)
}

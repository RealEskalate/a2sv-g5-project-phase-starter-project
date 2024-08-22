package route

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPromoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	pu := usecase.NewPromoteUsecase(ur, timeout)
	promoteUserController := controller.NewPromoteController(ur, pu, env)

	group.PUT("/promote-user/:id", promoteUserController.PromoteUser)
	group.PUT("/demote-user/:id", promoteUserController.DemoteUser)
}
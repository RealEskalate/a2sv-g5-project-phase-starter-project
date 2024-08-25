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


func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	ru := usecase.NewRefreshTokenUsecase(ur, timeout)
	uu := usecase.NewUserUsecase(ur, timeout)
	rc := controller.NewRefreshTokenController(uu, ru, env)

	group.POST("/refresh_token", rc.RefreshTokenRequest)
}
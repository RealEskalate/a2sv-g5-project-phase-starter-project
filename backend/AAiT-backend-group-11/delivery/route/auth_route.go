package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup)  {
	ur := repository.NewUserRepository(db.Collection("users"))
	usr := service.NewUserService(ur)
	tokenRepo := repository.NewRefreshTokenRepository(db.Collection("refresh_tokens"))

	ac := controller.AuthController{
		AuthService : service.NewAuthService(usr, tokenRepo),
		TokenRepo : repository.NewRefreshTokenRepository(db.Collection("refresh_tokens")),
		Env: env,
	}

	group.POST("/register", ac.RegisterUser)

}

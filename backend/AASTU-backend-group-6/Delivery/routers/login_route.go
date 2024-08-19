package routers

import (
	"blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *infrastructure.Config, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, env.UserCollection)
	aur := repositories.NewActiveUserRepository(db, env.ActiveUserCollection)
	lc := &controllers.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(ur, aur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

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

func NewLogoutRouter(env *infrastructure.Config, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	aur := repositories.NewActiveUserRepository(db, env.ActiveUserCollection)
	lc := &controllers.LogoutController{
		LogoutUsecase: usecases.NewLogoutUsecase(aur, timeout),
		Env:           env,
	}
	group.GET("/logout", lc.Logout)
}

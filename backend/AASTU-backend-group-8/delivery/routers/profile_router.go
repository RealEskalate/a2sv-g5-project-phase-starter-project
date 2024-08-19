package routers

import (
	"github.com/gin-gonic/gin"
	"meleket/domain"
)

type ProfileRouter struct {
	profileController domain.ProfileHandler
	engine            *gin.Engine
}

func NewProfileRouter(p domain.ProfileHandler, engine *gin.Engine) domain.ProfileRouter {
	return &ProfileRouter{
		profileController: p,
		engine:            engine,
	}
}

func (p *ProfileRouter) InitProfileRoutes(auth *gin.RouterGroup) {
	// User profile routes

	auth.GET("profile/:user_id", p.profileController.FindProfile)
	auth.PUT("profile/", p.profileController.UpdateProfile)
	auth.POST("profile", p.profileController.SaveProfile)
}

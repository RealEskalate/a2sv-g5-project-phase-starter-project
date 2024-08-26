package routers

import (
	"meleket/delivery/controllers"
	"meleket/domain"
	"meleket/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewProfileRoutes(r *gin.Engine, profileUsecase domain.ProfileUsecase, jwtService infrastructure.JWTService) {
	p := controllers.NewProfileHandler(profileUsecase)
	// User profile routes
	r.Static("/public", "./uploads")
	r.GET("profile/:user_id", p.FindProfile)
	auth := r.Group("/api")
	auth.Use(infrastructure.AdminMiddleware(jwtService))
	{
		auth.PUT("profile/", p.UpdateProfile)
		auth.POST("profile", p.SaveProfile)
	}
}

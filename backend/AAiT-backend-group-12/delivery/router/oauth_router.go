package router

import (
	"blog_api/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func NewOAuthRouter(routerGroup *gin.RouterGroup) {
	controller := controllers.NewOAuthController()
	routerGroup.GET("/auth/google/start", controller.GoogleAuthInit)
	routerGroup.GET("/auth/google/callback", controller.OAuthCallback)
}

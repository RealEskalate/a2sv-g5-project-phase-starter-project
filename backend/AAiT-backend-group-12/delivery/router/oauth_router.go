package router

import (
	"blog_api/delivery/controllers"

	"github.com/gin-gonic/gin"
)

// NewOAuthRouter sets up the OAuth routes
func NewOAuthRouter(routerGroup *gin.RouterGroup) {
	controller := controllers.NewOAuthController()
	routerGroup.GET("/auth/google/start", controller.GoogleAuthInit)
	routerGroup.GET("/auth/google/callback", controller.OAuthCallback)
}

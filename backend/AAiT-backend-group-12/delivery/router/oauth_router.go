package router

import (
	"blog_api/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func NewOAuthRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/auth/google/start", controllers.BeginGoogleAuth)
	routerGroup.GET("auth/google/callback", controllers.OAuthCallback)
}

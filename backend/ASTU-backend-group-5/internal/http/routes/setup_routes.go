package routes

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoute(router *gin.Engine) {
	RegisterUserRoutes(router)
	RegisterVerificationRoutes(router)
	RegisterAdminUserRoutes(router)
	RegisterBlogRoutes(router)

}

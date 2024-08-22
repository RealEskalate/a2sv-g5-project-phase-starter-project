package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type GeneralAuthorizationController interface {
	AuthenticationMiddleware() gin.HandlerFunc
	ADMINMiddleware() gin.HandlerFunc
	USERMiddleware() gin.HandlerFunc
	OWNERMiddleware() gin.HandlerFunc
}

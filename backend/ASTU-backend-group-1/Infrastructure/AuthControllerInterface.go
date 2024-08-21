package infrastructure

import "github.com/gin-gonic/gin"

type GeneralAuthorizationController interface {
	AuthMiddlewareGIn() gin.HandlerFunc
	AdminMiddlewareGin() gin.HandlerFunc
	UserMiddlewareGin() gin.HandlerFunc
	AuthorMiddlewareGin() gin.HandlerFunc
}

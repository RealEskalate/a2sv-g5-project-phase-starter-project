package controllers

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth infrastructure.GeneralAuthorizer
}

func NewAuthController(auth infrastructure.GeneralAuthorizer) GeneralAuthorizationController {
	return &AuthController{
		auth: auth,
	}
}
func (ac *AuthController) AuthMiddlewareGIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = os.Getenv("JWT_KEY")
		tokenString := c.GetHeader("Authorization")
		claims := ac.auth.AUTH(tokenString, secretKey)
		if claims != nil {
			c.Set("claims", claims)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort() 
	}

}

func (ac *AuthController) AdminMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an admin to do such task"})
				c.Abort()
			}
		}()
		claims := c.MustGet("claims")
		if !ac.auth.AdminAuth(claims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an admin to do such task"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (ac *AuthController) UserMiddlewareGin() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an user to do such task"})
				c.Abort()
			}
		}()
		claims := c.MustGet("claims")
		if !ac.auth.UserAuth(claims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "you must log in first"})
			c.Abort()
			return
		}
		c.Next()
	}
}

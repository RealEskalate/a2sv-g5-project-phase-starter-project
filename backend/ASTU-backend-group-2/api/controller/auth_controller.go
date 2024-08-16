package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlers interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	Signup() gin.HandlerFunc
	ForgotPassword() gin.HandlerFunc
	ResetPassword() gin.HandlerFunc
	VerifyEmail() gin.HandlerFunc
}

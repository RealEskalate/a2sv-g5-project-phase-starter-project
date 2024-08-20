package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Signup() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	ForgotPassword() gin.HandlerFunc
	ResetPassword() gin.HandlerFunc
	VerifyEmail() gin.HandlerFunc
}

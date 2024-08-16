package handlers

import (
    "github.com/gin-gonic/gin"
)

// Handler defines the interface for all HTTP handlers.
type UserInterfece interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	RequestVerifyEmail(c *gin.Context)
	VerifyEmail(c *gin.Context)
	ResetPasswordRequest(c *gin.Context)
	ResetPassword(c *gin.Context)
	


}
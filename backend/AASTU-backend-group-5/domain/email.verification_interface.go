package domain

import "github.com/gin-gonic/gin"

type VerifyEmail_Repository_interface interface {
	VerifyUser(id string) error
}

type VerifyEmail_Usecase_interface interface {
	SendVerifyEmail(id string, vuser VerifyEmail) error
	VerifyUser(token string) error
	SendForgretPasswordEmail(id string , vuser VerifyEmail) error
	ValidateForgetPassword(id string , token string) error
}

type VerifyEmail_Controller_interface interface {
	SendVerificationEmail() gin.HandlerFunc
	VerifyEmail() gin.HandlerFunc
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type EmailVControler struct{
	user_usecase domain.User_Usecase_interface
	email_uc domain.VerifyEmail_Usecase_interface
}

func NewEmailVController(email_usecase domain.VerifyEmail_Usecase_interface , user_usecase domain.User_Usecase_interface) *EmailVControler {
	return &EmailVControler{
		email_uc: email_usecase,
		user_usecase: user_usecase,
	}
}

func (ctrl *EmailVControler) SendVerificationEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model domain.VerifyEmail
		if err := ctx.BindJSON(&model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}
		id := ctx.Param("id")
		if err := ctrl.email_uc.SendVerifyEmail(id , model); err  != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"errorss" : err.Error()})
			return 
		}

		ctx.IndentedJSON(http.StatusAccepted , gin.H{"message" :fmt.Sprintf("email sent to: %s",model.Email)})
	}
}

func (ctrl *EmailVControler) VerifyEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Param("token")
		err := ctrl.email_uc.VerifyUser(token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted , gin.H{"message":"Verified"})
	}
}


func (ctrl *EmailVControler) ForgetPasswordValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		token := ctx.Query("token")

		err := ctrl.email_uc.ValidateForgetPassword(id , token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		update_password := domain.UpdatePassword{
			Password: "12345678",
			ConfirmPassword: "12345678",
		}
		user,err := ctrl.user_usecase.UpdatePassword(id , update_password)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted , gin.H{"user" : user,"message":"your password is reset to 12345678, you can change it anytime you want"})
	}
}

func (ctrl *EmailVControler) SendForgetPasswordEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model domain.VerifyEmail
		if err := ctx.BindJSON(&model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}
		id := ctx.Param("id")
		if err := ctrl.email_uc.SendForgretPasswordEmail(id , model); err  != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"errorss" : err.Error()})
			return 
		}

		ctx.IndentedJSON(http.StatusAccepted , gin.H{"message" :fmt.Sprintf("email sent to: %s",model.Email)})
	}
}

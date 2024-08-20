package controller

import (
	"fmt"
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

type EmailVControler struct{
	email_uc usecase.EmailVUsecase
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
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		ctx.IndentedJSON(http.StatusAccepted , gin.H{"message" :fmt.Sprintf("email sent to: %s",model.Email)})
	}
}

func (ctrl *EmailVControler) VerifyEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Param("token")
		ctrl.email_uc.VerifyUser(token)
	}
}
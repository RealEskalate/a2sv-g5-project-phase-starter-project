package controller

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ForgetPWController struct {
	Userusecase domain.UserUsecase
	ForgetPWUsecase domain.ForgetPWUsecase
	Env 		*bootstrap.Env
}

func NewForgetPWController(forgetPWUsecase domain.ForgetPWUsecase, userUsecase domain.UserUsecase, env *bootstrap.Env) *ForgetPWController {
	return &ForgetPWController{
		Userusecase: userUsecase,
		ForgetPWUsecase: forgetPWUsecase,
		Env: env,
	}
}

func (fpc *ForgetPWController) ForgetPW(c *gin.Context) {
	var request domain.ForgetPWRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:err.Error()})
		return
	}

	user, err := fpc.Userusecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,domain.ErrorResponse{Message:err.Error()})
		return
	}

	err = fpc.ForgetPWUsecase.ForgetPW(c, user.Email, fpc.Env.ServerAddress)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: "reset password request accepted"})
}

func (fpc *ForgetPWController) ResetPW(c *gin.Context) {
	var request domain.ResetPWRequest

	username, ok := c.GetQuery("user")
	if !ok{
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:"Invalid password recovery token"})
		return
	}
	recovery_token, ok := c.GetQuery("token")
	if !ok{
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:"Invalid password recovery token"})
		return
	}

	err := fpc.ForgetPWUsecase.VerifyForgetPWRequest(c, username, recovery_token)

	if err != nil{
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:err.Error()})
		return
	}

	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:err.Error()})
		return
	}

	user, err := fpc.Userusecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,domain.ErrorResponse{Message:err.Error()})
		return
	}

	if user.Username != username{
		c.JSON(http.StatusUnauthorized,domain.ErrorResponse{Message:"Username does not match"})
		return
	}

	password := infrastructure.HashPassword(request.Password)

	updatedUser := &domain.User{
		ID: user.ID,
		Password: password,
	}

	err = fpc.Userusecase.UpdateUser(c, updatedUser)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: "password reset successful"})
}
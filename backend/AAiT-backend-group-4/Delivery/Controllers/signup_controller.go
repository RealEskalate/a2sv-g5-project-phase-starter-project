package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SingupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) SignUp(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	otpResponse, err := sc.SingupUsecase.Signup(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, otpResponse)

}

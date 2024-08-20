package controllers

import (
	"net/http"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func NewUserController(userUC domain.UserUseCase) domain.UserController {
	return UserController{
		UserUseCase: userUC,
	}
}
func (userController *UserController) Register(cxt *gin.Context) {
	var registeringUser domain.User
	errUnmarshal := cxt.ShouldBind(&registeringUser)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}
	errCreate := userController.UserUseCase.Register(cxt, &registeringUser)
	if errCreate != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errCreate.Error()})
		return
	}
	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User Successfully Registered"})
}

func (userController *UserController) Login(c *gin.Context) {
  var loginInfo struct{username string `json:"username" binding="required"`, password string`json:"username" binding="required"`,}
  errUnmarshal := cxt.ShouldBind(&loginInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}
  errLogin
}

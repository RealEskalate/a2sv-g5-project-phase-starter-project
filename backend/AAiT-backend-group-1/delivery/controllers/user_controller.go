package controller

import (
	"fmt"
	"net/http"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type userController struct {
// 	UserUseCase domain.UserUseCase
// }

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	NewPasswor      string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController *UserController) Register(cxt *gin.Context) {
	var registeringUser domain.User
	errUnmarshal := cxt.ShouldBind(&registeringUser)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}
	errCreate := userController.UserUseCase.RegisterStart(cxt, &registeringUser)
	if errCreate != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errCreate.Error()})
		return
	}
	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User verification email sent"})
}

func (userController *UserController) VerifyEmail(cxt *gin.Context) {
	var token string
	errUnmarshal := cxt.ShouldBind(&token)
func (userController *UserController) Login(c *gin.Context) {
  var loginInfo struct{username string `json:"username" binding="required"`, password string`json:"username" binding="required"`,}
  errUnmarshal := cxt.ShouldBind(&loginInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errRegister := userController.UserUseCase.RegisterEnd(cxt, token)
	if errRegister != nil {
		cxt.JSON(errRegister.StatusCode(), gin.H{"Error": errRegister.Error()})
		return
	}

	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User email verified successfully "})
}

func (userController *UserController) Login(cxt *gin.Context) {
	var loginInfo LoginRequest
	errUnmarshal := cxt.ShouldBind(&loginInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	loginResult, errLogin := userController.UserUseCase.Login(cxt, loginInfo.Username, loginInfo.Password)
	if errLogin != nil {
		cxt.JSON(errLogin.StatusCode(), gin.H{"Error": errLogin.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"data": loginResult})
}

func (userController *UserController) ForgotPassword(cxt *gin.Context) {
	var email string
	errUnmarshal := cxt.ShouldBind(&email)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errForgot := userController.UserUseCase.ForgotPassword(cxt, email)
	if errForgot != nil {
		cxt.JSON(errForgot.StatusCode(), gin.H{"Error": errForgot.Error()})
	}
	cxt.JSON(http.StatusOK, gin.H{"Message": fmt.Sprintf("Reset link have been sent to the email %v", email)})
}

func (userController *UserController) ResetPassword(cxt *gin.Context) {
	var resetInfo ResetPasswordRequest
	errUnmarshal := cxt.ShouldBind(&resetInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errReset := userController.UserUseCase.ResetPassword(cxt, resetInfo.Token)
	if errReset != nil {
		cxt.JSON(errReset.StatusCode(), gin.H{"Error": errReset.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "Password reset successfully"})

}

func (userController *UserController) Logout(cxt *gin.Context) {
	var logoutInfo LogoutRequest
	errUnmarshal := cxt.ShouldBind(&logoutInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errLogout := userController.UserUseCase.Logout(cxt, map[string]string{
		"access_token":  logoutInfo.AccessToken,
		"refresh_token": logoutInfo.RefreshToken,
	})

	if errLogout != nil {
		cxt.JSON(errLogout.StatusCode(), gin.H{"Error": errLogout.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User logged out successfully"})

}

func (userController *UserController) PromoteUser(cxt *gin.Context) {
	var updateID string
	errUnmarshal := cxt.ShouldBind(&updateID)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errPromote := userController.UserUseCase.PromoteUser(cxt, updateID)
	if errPromote != nil {
		cxt.JSON(errPromote.StatusCode(), gin.H{"Error": errPromote.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User promoted successfully"})
}

func (userController *UserController) DemoteUser(cxt *gin.Context) {
	var updateID string
	errUnmarshal := cxt.ShouldBind(&updateID)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errDemote := userController.UserUseCase.DemoteUser(cxt, updateID)
	if errDemote != nil {
		cxt.JSON(errDemote.StatusCode(), gin.H{"Error": errDemote.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User demoted successfully"})
}

func (userController *UserController) UpdateProfile(cxt *gin.Context) {
	var updateInfo domain.User
	errUnmarshal := cxt.ShouldBind(&updateInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	userID := cxt.Param("id")

	updateInfo.ID = primitive.NewObjectID()

	errUpdate := userController.UserUseCase.UpdateProfile(cxt, userID, &updateInfo)
	if errUpdate != nil {
		cxt.JSON(errUpdate.StatusCode(), gin.H{"Error": errUpdate.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User profile updated successfully"})

}

package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	NewPasswor      string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string
	ResetToken      int `json:"reset_token"`
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type userController struct {
	UserUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) domain.UserController {
	return &userController{
		UserUseCase: userUseCase,
	}
}

func (userController *userController) Register(cxt *gin.Context) {
	var registeringUser domain.User
	errUnmarshal := cxt.ShouldBind(&registeringUser)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Edrror": errUnmarshal.Error()})
		return
	}
	errCreate := userController.UserUseCase.RegisterStart(cxt, &registeringUser)
	if errCreate != nil {
		cxt.JSON(errCreate.StatusCode(), gin.H{"Errojgr": errCreate.Error()})
		return
	}
	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User verification email sent"})
}

func (userController *userController) VerifyEmail(cxt *gin.Context) {
	token := cxt.Param("token")
	errRegister := userController.UserUseCase.RegisterEnd(cxt, token)
	if errRegister != nil {
		cxt.JSON(errRegister.StatusCode(), gin.H{"Error": errRegister.Error()})
		return
	}

	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User email verified successfully "})
}

func (userController *userController) Login(cxt *gin.Context) {
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

func (userController *userController) RefreshToken(cxt *gin.Context) {
	var refreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	errUnmarshal := cxt.ShouldBind(&refreshToken)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	refreshResult, errRefresh := userController.UserUseCase.RefreshToken(cxt, refreshToken.RefreshToken)
	if errRefresh != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"Error": errRefresh.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"data": refreshResult})
}

func (userController *userController) ForgotPassword(cxt *gin.Context) {
	var email struct {
		Email string `json:"email"`
	}
	errUnmarshal := cxt.ShouldBind(&email)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errForgot := userController.UserUseCase.ForgotPassword(cxt, email.Email)
	if errForgot != nil {
		cxt.JSON(errForgot.StatusCode(), gin.H{"Error": errForgot.Error()})
		return
	}
	cxt.JSON(http.StatusOK, gin.H{"Message": fmt.Sprintf("Reset link have been sent to the email %v", email.Email)})
}

func (userController *userController) ResetPassword(cxt *gin.Context) {
	var resetInfo ResetPasswordRequest
	errUnmarshal := cxt.ShouldBind(&resetInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	resetToken := cxt.Param("token")
	resetInfo.Token = resetToken

	errReset := userController.UserUseCase.ResetPassword(cxt, resetInfo.NewPasswor, resetInfo.ConfirmPassword, resetInfo.Token, resetInfo.ResetToken)
	if errReset != nil {
		cxt.JSON(errReset.StatusCode(), gin.H{"Error": errReset.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "Password reset successfully"})

}

func (userController *userController) Logout(cxt *gin.Context) {
	var logoutInfo LogoutRequest
	errUnmarshal := cxt.ShouldBind(&logoutInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	accessToken := cxt.GetHeader("Authorization")

	errLogout := userController.UserUseCase.Logout(cxt, map[string]string{
		"access_token":  strings.TrimPrefix(accessToken, "Bearer "),
		"refresh_token": logoutInfo.RefreshToken,
	})

	if errLogout != nil {
		cxt.JSON(errLogout.StatusCode(), gin.H{"Error": errLogout.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User logged out successfully"})

}

func (userController *userController) PromoteUser(cxt *gin.Context) {
	var updateID map[string]string
	errUnmarshal := cxt.ShouldBind(&updateID)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errPromote := userController.UserUseCase.PromoteUser(cxt, updateID["id"])
	if errPromote != nil {
		cxt.JSON(errPromote.StatusCode(), gin.H{"Error": errPromote.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User promoted successfully"})
}

func (userController *userController) DemoteUser(cxt *gin.Context) {
	var updateID map[string]string
	errUnmarshal := cxt.ShouldBind(&updateID)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	errDemote := userController.UserUseCase.DemoteUser(cxt, updateID["id"])
	if errDemote != nil {
		cxt.JSON(errDemote.StatusCode(), gin.H{"Error": errDemote.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User demoted successfully"})
}

func (userController *userController) UpdateProfile(cxt *gin.Context) {
	var updateInfo map[string]interface{}
	errUnmarshal := cxt.ShouldBind(&updateInfo)
	if errUnmarshal != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
		return
	}

	userID := cxt.Param("id")

	errUpdate := userController.UserUseCase.UpdateProfile(cxt, userID, updateInfo)
	if errUpdate != nil {
		cxt.JSON(errUpdate.StatusCode(), gin.H{"Error": errUpdate.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"Message": "User profile updated successfully"})
}

func (userController *userController) ImageUpload(cxt *gin.Context) {
	file, header, errFile := cxt.Request.FormFile("profile_picture")
	if errFile != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errFile.Error()})
		return
	}
	defer file.Close()

	userID := cxt.Param("id")

	errUpload := userController.UserUseCase.ImageUpload(cxt, &file, header, userID)
	if errUpload != nil {
		cxt.JSON(errUpload.StatusCode(), gin.H{"Error": errUpload.Error()})
		return
	}
	cxt.JSON(http.StatusOK, gin.H{"Message": "Image uploaded successfully"})
}

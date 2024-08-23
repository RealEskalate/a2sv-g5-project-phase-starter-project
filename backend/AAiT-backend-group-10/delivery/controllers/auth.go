package controllers

import (
	"context"
	"log"
	"net/http"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUsecase  usecases.IAuthUsecase
	googleConfig *oauth2.Config
}

func NewAuthController(uc usecases.IAuthUsecase, googleConfig *oauth2.Config) *AuthController {
	return &AuthController{
		userUsecase:  uc,
		googleConfig: googleConfig,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var userDTO dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userUsecase.RegisterUser(&userDTO)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *AuthController) Login(c *gin.Context) {
	var loginDTO dto.LoginUserDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := uc.userUsecase.LoginUser(&loginDTO)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (uc *AuthController) RefreshToken(c *gin.Context) {
	var refreshToken dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := uc.userUsecase.RefreshTokens(refreshToken.RefreshToken)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (uc *AuthController) ForgotPassword(c *gin.Context) {
	var forgotPasswordDTO dto.ForgotPasswordRequestDTO
	if err := c.ShouldBindJSON(&forgotPasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.ForgotPassword(&forgotPasswordDTO)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email"})
}

func (uc *AuthController) ResetPassword(c *gin.Context) {
	var resetPasswordDTO dto.ResetPasswordRequestDTO
	if err := c.ShouldBindJSON(&resetPasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.ResetPassword(&resetPasswordDTO)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func (uc *AuthController) HandleGoogleCallback(ctx *gin.Context) {
	code := ctx.Query("code")

	// Exchange the code for a token
	token, err := uc.googleConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Println("No id_token found")
		ctx.IndentedJSON(400, gin.H{
			"message": "no id_token found",
		})
		return
	}

	// Verify the ID token and extract the user's information
	payload, err := idtoken.Validate(context.Background(), idToken, uc.googleConfig.ClientID)
	if err != nil {
		log.Printf("Failed to validate ID token: %s\n", err.Error())
		ctx.IndentedJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	userInfo := &domain.User{
		FullName:     payload.Claims["name"].(string),
		Email:        payload.Claims["email"].(string),
		GoogleSignIn: true,
		IsAdmin:      false,
	}
	picture, ok := payload.Claims["picture"].(string)
	if ok {
		userInfo.ImageURL = picture
	}
	accesstoken, refreshtoken, errs := uc.userUsecase.HandleGoogleCallback(userInfo)
	if errs != nil {
		ctx.IndentedJSON(400, gin.H{
			"message": errs.Error(),
		})
		return
	}
	ctx.IndentedJSON(400, &dto.TokenResponseDTO{
		RefreshToken: refreshtoken,
		AccessToken:  accesstoken,
	})
}

func (uc *AuthController) HandleGoogleLogin(c *gin.Context) {
	// Redirect to Google login page
	url := uc.googleConfig.AuthCodeURL("state-token")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

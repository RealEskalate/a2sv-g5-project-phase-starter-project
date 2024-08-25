package controller

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	UserUsecase   domain.UserUsecase
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func NewSignupController(uu domain.UserUsecase,su domain.SignupUsecase, env *bootstrap.Env) *SignupController {
	return &SignupController{
		UserUsecase: uu,
		SignupUsecase: su,
		Env: env,
	}
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.UserUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	_, err = sc.UserUsecase.GetUserByUsername(c, request.Username)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given username"})
		return
	}

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user.Token = accessToken
	user.Refresh_token = refreshToken

	validationErr := infrastructure.ValidateUser(&user)
	if validationErr != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: validationErr.Error()})
		return
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Success: true, Message: "user created successfuly", Data: signupResponse})
}

func (sc *SignupController) VerifyEmail(c *gin.Context){
	var request domain.VerifyEmailRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := sc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = sc.SignupUsecase.VerifyEmail(c, request.Email, request.Verification_code)

	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updatedUser := &domain.User{
		ID: user.ID,
		Is_active: true,
	}

	err = sc.UserUsecase.UpdateUser(c, updatedUser)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: "Email verified successfully"})
}
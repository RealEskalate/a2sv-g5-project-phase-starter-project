package controller

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


type RefreshTokenController struct {
	UserUsecase domain.UserUsecase
	refreshTokenUsecase domain.RefreshTokenUsecase
	Env 		*bootstrap.Env
}

func NewRefreshTokenController(uu domain.UserUsecase, ru domain.RefreshTokenUsecase, env *bootstrap.Env) *RefreshTokenController {
	return &RefreshTokenController{
		UserUsecase: uu,
		refreshTokenUsecase: ru,
		Env: env,
	}
}

func (rc *RefreshTokenController) RefreshTokenRequest(c *gin.Context) {
	var request domain.RefreshTokenRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message:err.Error()})
		return
	}

	claims, err := utils.ValidateToken(request.RefreshToken, rc.Env.RefreshTokenSecret)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JWT"})
		return
	}

	user, err := rc.UserUsecase.GetUserByID(c, claims.User_id)
	if err != nil {
		c.JSON(http.StatusNotFound,domain.ErrorResponse{Message:err.Error()})
	}

	accessToken, err := rc.refreshTokenUsecase.CreateAccessToken(user, rc.Env.AccessTokenSecret, rc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	
	updatedUser := &domain.User{
		ID: user.ID,
		Token: accessToken,
	}

	err = rc.UserUsecase.UpdateUser(c, updatedUser)

	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: "access token refresh successful", Data: refreshTokenResponse})
}
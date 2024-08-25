package controller

import (
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase entities.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request entities.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
			return
		}
		middleware.CustomErrorResponse(c, err)
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		return
	}

	if !user.Active {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotActive))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrCredentialsNotValid))
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	err = lc.LoginUsecase.UpdateRefreshToken(c.Request.Context(), user.ID.Hex(), refreshToken)

	loginResponse := entities.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

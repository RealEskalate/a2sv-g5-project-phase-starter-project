package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type OpenAuthController interface {
	GoogleAuthCallback() gin.HandlerFunc
}

// ProfileController is a struct to hold the usecase and env
type OAuthController struct {
	UserUsecase  entities.UserUsecase
	LoginUsecase entities.LoginUsecase
	Env          *bootstrap.Env
}

func (oc *OAuthController) OAuthCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		// provider := c.Param("provider")
		// c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))

		provider := c.Param("provider")
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

		fmt.Println("user from callback", user.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Check if user exists
		dbUser, err := oc.UserUsecase.GetUserByEmail(context.TODO(), user.Email)

		if err != nil {
			// User does not exist, create user

			// Check if user is owner, if user to be created is the first user, set as owner
			isOwner, err := oc.UserUsecase.IsOwner(context.TODO())

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			insertUser := entities.User{
				Email:      user.Email,
				FirstName:  user.FirstName,
				LastName:   user.LastName,
				Active:     true,
				ProfileImg: user.AvatarURL,
				IsOwner:    isOwner,
				Password:   "",
				Role:       "user",
			}

			dbUser, err = oc.UserUsecase.CreateUser(context.TODO(), &insertUser)

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

		}

		accessToken, err := oc.LoginUsecase.CreateAccessToken(dbUser, oc.Env.AccessTokenSecret, oc.Env.AccessTokenExpiryHour)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		refreshToken, err := oc.LoginUsecase.CreateRefreshToken(dbUser, oc.Env.RefreshTokenSecret, oc.Env.RefreshTokenExpiryHour)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = oc.LoginUsecase.UpdateRefreshToken(c.Request.Context(), dbUser.ID.Hex(), refreshToken)

		loginResponse := entities.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		c.JSON(http.StatusOK, loginResponse)

	}
}

func (pc *OAuthController) OAuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := c.Param("provider")
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (pc *OAuthController) OAuthLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		gothic.Logout(c.Writer, c.Request)

		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}

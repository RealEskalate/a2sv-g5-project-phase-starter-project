package domain

import "github.com/gin-gonic/gin"

type AuthController interface {
	SignUp() gin.HandlerFunc
	OauthSignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	OauthLogIn() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

type AuthUsecase interface {
	RegisterUser(input RegisterUser) (User, error)
	OAuthSignUp(provider, token string) (User, string, string, error) 
	LoginUser(email, password string) (User, string, string, error)
	OAuthLogin(provider, token string) (User, string, string, error)
	RefreshTokens(refreshToken string) (string, string, error) 
}

type AuthRepository interface{
	SaveUser(user *User) error
	FindUserByEmail(email string) (*User, error)
	FindUserByOAuthID(provider, oauthID string) (*User, error)
}
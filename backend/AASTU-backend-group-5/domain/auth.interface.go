package domain

import "github.com/gin-gonic/gin"

type AuthController interface {
	SignUp() gin.HandlerFunc
	OauthSignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	OauthLogIn() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	VerifyEmail() gin.HandlerFunc
	Refresh() gin.HandlerFunc
	ForgotPassword() gin.HandlerFunc
}

type AuthUsecase interface {
	RegisterUser(input RegisterUser) (User, error)
	OAuthSignUp(provider, token string) (User, string, string, error) 
	LoginUser(email, password string) (User, string, string, error)
	OAuthLogin(provider, token string) (User, string, string, error)
	VerifyEmail(verificationToken string) error
	RefreshTokens(refreshToken string) (string, string, error) 
	Logout(userID string) error
	ForgotPassword(email string) error                           
	ResetPassword(resetToken, newPassword string) error     
}

type AuthRepository interface{
	SaveUser(user *User) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id string) (*User, error)
	UpdateUser(user *User) error
	FindUserByOAuthID(provider, oauthID string) (*User, error)
	SavePasswordResetToken(email, resetToken string) error  
	FindUserByResetToken(resetToken string) (*User, error)  
}
package domain

import "github.com/gin-gonic/gin"

type AuthController interface {
	SignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	GoogleLogIn() gin.HandlerFunc
	GoogleCallBack() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

type AuthUsecase interface {
	RegisterUser(RegisterUser) (User, error)
	LoginUser(string, string) (User, string, string, error)
	GoogleLogin() (string, error)
	GoogleCallBack(string, string)(*User, string, string, error)
	RefreshTokens(string) (string, string, error) 
}

type AuthRepository interface{
	SaveUser(*User) error
	FindUserByEmail(string) (*User, error)
}

type StateRepository interface{
	InsertState(State) error
	GetState(string) (*State, error)
	DeleteState(string) error
}
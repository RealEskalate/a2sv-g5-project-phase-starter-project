package interfaces

import  "backend-starter-project/domain/entities"

type AuthenticationService interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	Login(emailOrUsername, password string) (*entities.RefreshToken,string, error)
	Logout(userId string) error
	RefreshAccessToken(token string) (string,error)
	VerifyEmail(email string, code string) error
	ResendOtp(request entities.ResendOTPRequest) error 
}


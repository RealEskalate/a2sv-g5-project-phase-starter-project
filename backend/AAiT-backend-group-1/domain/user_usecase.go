package domain

type UserUseCase interface {
	Register(user *User) CustomError
	Login(email, password string) (string, string, CustomError)
	Authenticate(token string) (string, CustomError)
	ForgotPassword(email string) CustomError
	Logout(token string) CustomError
	PromoteUser(userID string) CustomError
	DemoteUser(userID string) CustomError
	UpdateProfile(userID string, user *User) CustomError
}
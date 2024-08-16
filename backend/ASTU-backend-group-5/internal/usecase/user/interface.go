package user

import (
	"blogApp/internal/domain"
)

type UserUseCaseInterface interface {
	RegisterUser(user *domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, *domain.Token, error)

	RequestEmailVerification(user domain.User) error
	RequestPasswordResetUsecase(userEmail string) error
	ResetPassword(token string, password string, email string) error
	VerifyEmail(token string, email string) error

	UpdateUser(user *domain.User) error
	DeleteUser(id string) error

	PromoteToAdmin(UserId string) error
	GetAllUsers() ([]*domain.User, error)
	FilterUsers(filter map[string]interface{}) ([]*domain.User, error)
}

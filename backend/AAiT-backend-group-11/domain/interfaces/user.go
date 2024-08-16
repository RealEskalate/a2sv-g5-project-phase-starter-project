package interfaces

import   "backend-starter-project/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserById(userId string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(userId string) error
	PromoteUserToAdmin(userId string) error
	DemoteUserToRegular(userId string) error
}

type UserService interface {
	CreateUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserById(userId string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(userId string) error
	PromoteUserToAdmin(userId string) error
	DemoteUserToRegular(userId string) error
}
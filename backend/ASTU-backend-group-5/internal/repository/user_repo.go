package repository

import(
	"blogApp/internal/domain"
)

type UserRepo interface{
	CreateUser(data *domain.User) (*domain.User, error)
}


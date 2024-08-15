package Repositories

import (
	"blogapp/Domain"
)

type authRepository struct {
	collection Domain.Collection
}

func NewAuthRepository(_collection Domain.Collection) *authRepository {
	return &authRepository{

		collection: _collection,
	}

}

// login
func (a *authRepository) Login() (string, error, int) {
	// return error
	return "", nil, 0
}

// register
func (a *authRepository) Register() (*Domain.OmitedUser, error, int) {
	// return error
	return nil, nil, 0
}

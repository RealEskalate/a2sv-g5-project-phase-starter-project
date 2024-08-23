package domain

import "context"

type UnverifiedUser struct {
	Email     string
	UserToken string
	OTP       string
}
type UnverifiedUserRepository interface {
	StoreUnverifiedUser(ctx context.Context, uv UnverifiedUser) error
	FindUnverifiedUser(ctx context.Context, email string) (UnverifiedUser, error)
	DeleteUnverifiedUser(ctx context.Context, email string) error
}

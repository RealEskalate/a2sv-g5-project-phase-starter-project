package repository

import (
	"blogApp/internal/domain"
	"context"
)

// UserRepository is an interface that defines the methods a user repository should implement
// You can add more methods if needed
// FindByUserName(ctx context.Context, username string) (*domain.User, error)
// Update(ctx context.Context, user *domain.User) error
// Delete(ctx context.Context, id string) error
// CreateUser(ctx context.Context, user *domain.User) error
type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// FindUserByUserName(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdatePassword(ctx context.Context, email string, password string) error
	UpdateUser(ctx context.Context, user *domain.User) error
	// UpdateUser(ctx context.Context, user *domain.User) error
	// DeleteUser(ctx context.Context, id string) error
	// UpdateUserPassword(ctx context.Context, id string, password string) error
	// UpdateUserRole(ctx context.Context, id string, role string) error
	// UpdateUserEmail(ctx context.Context, id string, email string) error
}

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
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	
}

package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) *models.ErrorResponse

	GetUserByEmailOrUsername(ctx context.Context, username string, email string) (*models.User, *models.ErrorResponse)
	GetUserByID(ctx context.Context, id string) (*models.User, *models.ErrorResponse)

	UpdateUser(ctx context.Context, user *models.User, id string) *models.ErrorResponse
	DeleteUser(ctx context.Context, userID string) *models.ErrorResponse

	PromoteUser(ctx context.Context, userID string) *models.ErrorResponse
	DemoteUser(ctx context.Context, userID string) *models.ErrorResponse

	StoreAccessToken(ctx context.Context, userID string, token string) *models.ErrorResponse
	StoreRefreshToken(ctx context.Context, userID string, token string) *models.ErrorResponse

	DeleteTokensFromDB(ctx context.Context, userID string) *models.ErrorResponse
}

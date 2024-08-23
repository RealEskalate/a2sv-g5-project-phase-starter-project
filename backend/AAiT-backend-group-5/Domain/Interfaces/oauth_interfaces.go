package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type OAuthController interface {
	LoginHandlerController(*gin.Context)
	OAuthHanderController(*gin.Context)
	OAuthCallbackHandler(*gin.Context)
}

type OAuthUseCase interface {
	LoginHandlerUseCase(ctx context.Context, user dtos.OAuthRequest) *models.ErrorResponse
	SaveSession(ctx context.Context, user dtos.OAuthRequest) *models.ErrorResponse
}

type OAuthService interface {
	OAuthTokenValidator(token string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse)
	RefreshTokenValidator(refreshToken string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse)
	GenerateAccessToken(ctx context.Context, refreshToken string) (string, *models.ErrorResponse)
}

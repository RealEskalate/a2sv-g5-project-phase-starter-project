package usecases

import (
	"context"
	"aait-backend-group4/Domain"
)

type logoutUsecase struct {
	TokenService domain.TokenInfrastructure
}

func NewLogoutUsecase(tokenService domain.TokenInfrastructure) domain.LogoutUsecase {
	return &logoutUsecase{
		TokenService: tokenService,
	}
}

func (u *logoutUsecase) Logout(ctx context.Context, token string) (domain.LogoutResponse, error) {
	// Extract user ID from the token using the TokenService
	userID, err := u.TokenService.ExtractUserIDFromToken(token)
	if err != nil {
		return domain.LogoutResponse{}, err
	}

	// Remove tokens for the user
	err = u.TokenService.RemoveTokens(userID)
	if err != nil {
		return domain.LogoutResponse{}, err
	}

	// Return a response indicating successful logout
	return domain.LogoutResponse{
		Message: "Logout successful",
	}, nil
}

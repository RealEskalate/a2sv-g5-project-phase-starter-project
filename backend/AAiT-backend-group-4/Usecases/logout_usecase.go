package usecases

import (
	"context"
	"aait-backend-group4/Domain"
)

type logoutUsecase struct {
	TokenService domain.TokenInfrastructure
}

// NewLogoutUsecase initializes a new instance of logoutUsecase
// This use case handles the logout process by interacting with the TokenService.
// It requires a token service (infrastructure) that manages the tokens.
func NewLogoutUsecase(tokenService domain.TokenInfrastructure) domain.LogoutUsecase {
	return &logoutUsecase{
		TokenService: tokenService,
	}
}

// Logout processes the logout request for a user
// It first extracts the user ID from the provided token using the TokenService.
// Then, it attempts to remove all tokens associated with that user ID to complete the logout.
// If successful, it returns a response indicating a successful logout; otherwise, it returns an error.
func (u *logoutUsecase) Logout(ctx context.Context, token string) (domain.LogoutResponse, error) {
	

	userID, err := u.TokenService.ExtractUserIDFromToken(token)
	if err != nil {
		return domain.LogoutResponse{}, err
	}

	err = u.TokenService.RemoveTokens(userID)
	if err != nil {
		return domain.LogoutResponse{}, err
	}

	return domain.LogoutResponse{
		Message: "Logout successful",
	}, nil
}

package usecase

import (
	"time"

	"group3-blogApi/domain"
)

func (uc *UserUsecase) OAuthLogin(oauthUserInfo domain.OAuthUserInfo, deviceID string) (domain.LogInResponse, error) {
	// Find or create the user based on OAuth information
	user, err := uc.UserRepo.FindOrCreateUserByGoogleID(oauthUserInfo, deviceID)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	// Generate access token
	accessToken, err := uc.TokenGen.GenerateToken(*user)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	// Generate refresh token
	refreshToken, err := uc.TokenGen.GenerateRefreshToken(*user)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	// Add the new refresh token to the user's list of refresh tokens
	user.RefreshTokens = append(user.RefreshTokens, domain.RefreshToken{
		Token:     refreshToken,
		DeviceID:  deviceID,
		CreatedAt: time.Now(),
	})

	// Update the user in the repository
	err = uc.UserRepo.UpdateUser(user)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	// Return the login response with the generated tokens
	return domain.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

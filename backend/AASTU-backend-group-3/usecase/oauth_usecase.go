package usecase

import (
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"time"
)

func (uc *UserUsecase) OAuthLogin(oauthUserInfo domain.OAuthUserInfo, deviceID string) (domain.LogInResponse, error) {

	user, err := uc.UserRepo.FindOrCreateUserByGoogleID(oauthUserInfo, deviceID)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	// Generate access and refresh tokens
	accessToken, err := infrastracture.GenerateToken(*user)
	if err != nil {
		return domain.LogInResponse{}, err
	}
	
	refreshToken, err := infrastracture.GenerateRefreshToken(user)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	user.RefreshTokens = append(user.RefreshTokens, domain.RefreshToken{
		Token:     refreshToken,
		DeviceID:  deviceID,
		CreatedAt: time.Now(),
	})

	err = uc.UserRepo.UpdateUser(user)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	return domain.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

	
}
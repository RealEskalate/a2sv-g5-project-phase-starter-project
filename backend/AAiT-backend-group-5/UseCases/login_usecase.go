package usecases

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type loginUsecase struct {
	jwtService      interfaces.JwtService
	passwordService interfaces.PasswordService
	repository      interfaces.UserRepository
}

func NewLoginUsecase(jwtService interfaces.JwtService, passwordService interfaces.PasswordService, repository interfaces.UserRepository) interfaces.LoginUsecase {
	return &loginUsecase{
		jwtService:      jwtService,
		passwordService: passwordService,
		repository:      repository,
	}
}

func (uc *loginUsecase) LoginUser(ctx context.Context, userReqest dtos.LoginRequest) (*dtos.LoginResponse, *models.ErrorResponse) {

	// check the user exists
	user, err := uc.repository.GetUserByEmailOrUsername(ctx, userReqest.UsernameOrEmail, userReqest.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	// validate password
	if validPassword := uc.passwordService.ValidatePassword(userReqest.Password, user.Password); !validPassword {
		return nil, models.Unauthorized("Invalid creaditional")
	}

	// generate access token
	accessToken, aErr := uc.GenerateAccessToken(user, 15)
	refresheToken, rErr := uc.GenerateRefreshToken(user, 15)

	if aErr != nil || rErr != nil {
		return nil, models.InternalServerError("Something went wrong")
	}

	// save the refresh token
	err = uc.repository.StoreAccessToken(ctx, user.ID, accessToken)
	tErr := uc.repository.StoreRefreshToken(ctx, user.ID, refresheToken)
	if err != nil {
		return nil, err
	}

	if tErr != nil {
		return nil, tErr
	}

	return &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refresheToken,
	}, nil

}
func (uc *loginUsecase) GenerateAccessToken(user *models.User, expiry int) (string, *models.ErrorResponse) {
	token, err := uc.jwtService.CreateAccessToken(*user, expiry)

	if err != nil {
		return "", models.InternalServerError("Error occured while generating the access token")
	}

	return token, nil
}

func (uc *loginUsecase) GenerateRefreshToken(user *models.User, expiry int) (string, *models.ErrorResponse) {
	token, err := uc.jwtService.CreateRefreshToken(*user, expiry)

	if err != nil {
		return "", models.InternalServerError("Error occured while creating the refresh token")
	}

	return token, nil
}

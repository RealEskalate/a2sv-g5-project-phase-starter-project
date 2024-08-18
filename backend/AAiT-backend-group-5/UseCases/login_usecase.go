package usecases

import (
	"context"

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

func (uc *loginUsecase) LoginUser(ctx context.Context, emailOrUsername string, password string) (*models.User, *models.ErrorResponse) {

	// check the user exists
	user, err := uc.repository.GetUserByEmailOrUsername(ctx, emailOrUsername, emailOrUsername)
	if err != nil {
		return nil, err
	}

	// validate password
	if validPassword := uc.passwordService.ValidatePassword(password, user.Password); !validPassword {
		return nil, models.Unauthorized("Invalid creaditional")
	}

	return user, nil

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

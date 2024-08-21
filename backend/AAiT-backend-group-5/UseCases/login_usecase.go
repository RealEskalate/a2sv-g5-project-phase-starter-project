package usecases

import (
	"context"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type loginUsecase struct {
	jwtService      interfaces.JwtService
	passwordService interfaces.PasswordService
	repository      interfaces.UserRepository
	session         interfaces.SessionRepository
	env             config.Env
}

func NewLoginUsecase(jwtService interfaces.JwtService, passwordService interfaces.PasswordService, repository interfaces.UserRepository, session interfaces.SessionRepository, env config.Env) interfaces.LoginUsecase {
	return &loginUsecase{
		jwtService:      jwtService,
		passwordService: passwordService,
		repository:      repository,
		session:         session,
		env:             env,
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
	accessToken, aErr := uc.GenerateAccessToken(user, uc.env.ACCESS_TOKEN_EXPIRY_HOUR)
	refresheToken, rErr := uc.GenerateRefreshToken(user, uc.env.REFRESH_TOKEN_EXPIRY_HOUR)

	if aErr != nil || rErr != nil {
		return nil, models.InternalServerError("Something went wrong")
	}

	// save the refresh token
	session := models.Session{
		UserID:       user.ID,
		RefreshToken: refresheToken,
	}

	userToken, _ := uc.session.GetToken(ctx, user.ID)

	if userToken != nil {
		if tErr := uc.session.UpdateToken(ctx, &session); tErr != nil {
			return nil, tErr
		}
	} else {
		if tErr := uc.session.SaveToken(ctx, &session); tErr != nil {
			return nil, tErr
		}
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

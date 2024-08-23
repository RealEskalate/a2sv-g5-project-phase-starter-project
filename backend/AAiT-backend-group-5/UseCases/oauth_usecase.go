package usecases

import (
	"context"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type OAuthUseCase struct {
	userRepo interfaces.UserRepository
	session  interfaces.SessionRepository
}

func NewOAuthUseCase(
	userRepo interfaces.UserRepository,
	session interfaces.SessionRepository,
) interfaces.OAuthUseCase {
	return &OAuthUseCase{
		userRepo: userRepo,
		session:  session,
	}
}

func (uc *OAuthUseCase) LoginHandlerUseCase(ctx context.Context, requestUser dtos.OAuthRequest) *models.ErrorResponse {
	newUser := &models.User{
		Username: requestUser.Email,
		Email:    requestUser.Email,
		Name:     requestUser.Name,
	}

	existUser, err := uc.userRepo.GetUserByEmailOrUsername(ctx, newUser.Username, newUser.Email)
	if err != nil && err.Code != http.StatusNotFound {
		return err
	}

	if existUser != nil {
		if err := uc.SaveSession(ctx, requestUser); err != nil {
			return err
		}
		return nil
	}

	if err.Code == http.StatusNotFound {
		if err := uc.userRepo.CreateUser(ctx, newUser); err != nil {
			return err
		}

		if err := uc.SaveSession(ctx, requestUser); err != nil {
			return err
		}
	}

	return nil
}

func (uc *OAuthUseCase) SaveSession(ctx context.Context, requestUser dtos.OAuthRequest) *models.ErrorResponse {
	user, err := uc.userRepo.GetUserByEmailOrUsername(ctx, requestUser.Email, requestUser.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return models.NotFound("User not found")
	}

	newSession := &models.Session{
		UserID:       user.ID,
		RefreshToken: requestUser.RefreshToken,
		AccessToken:  requestUser.AccessToken,
	}

	err = uc.session.SaveToken(ctx, newSession)
	if err != nil {
		return err
	}

	return nil
}

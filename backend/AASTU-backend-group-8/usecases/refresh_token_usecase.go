package usecases

import (
	"errors"
	// "fmt"
	"meleket/domain"
	"meleket/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenUsecase struct {
	tokenRepo domain.TokenRepositoryInterface
	jwtSvc    infrastructure.JWTService
}

func NewTokenUsecase(tr domain.TokenRepositoryInterface, js infrastructure.JWTService) *TokenUsecase {
	return &TokenUsecase{
		tokenRepo: tr,
		jwtSvc:    js,
	}
}

// RefreshToken refreshes a user's JWT token
func (u *TokenUsecase) RefreshToken(userID primitive.ObjectID, role string) (string, error) {
	_, err := u.tokenRepo.FindRefreshToken(userID)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// Generate a new JWT token
	// fmt.Println("yazew",userID, role)
	newToken, err := u.jwtSvc.GenerateToken(userID, role)
	if err != nil {
		return "", errors.New("could not generate new JWT token")
	}

	return newToken, nil
}

func (u *TokenUsecase) DeleteRefreshToken(userID primitive.ObjectID) error {
	err := u.tokenRepo.DeleteRefreshTokenByUserID(userID)
	return err
}

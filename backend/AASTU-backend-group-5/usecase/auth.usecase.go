package usecase

import (
	// "errors"
	"fmt"
	// "time"

	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUsecase struct {
	AuthRepo      domain.AuthRepository
	PasswordSrv   domain.PasswordService
	TokenSrv      domain.TokenService
	OAuthSrv      domain.OAuthService
}

func NewAuthUsecase(authRepo domain.AuthRepository, passwordSrv domain.PasswordService, tokenSrv domain.TokenService, oauthSrv domain.OAuthService) *AuthUsecase {
	return &AuthUsecase{
		AuthRepo:    authRepo,
		PasswordSrv: passwordSrv,
		TokenSrv:    tokenSrv,
		OAuthSrv:    oauthSrv,
	}
}

func (u *AuthUsecase) RegisterUser(input domain.RegisterUser) (domain.User, error) {
    var user domain.User

    // Hash the user's password
    hashedPassword, err := u.PasswordSrv.HashPassword(input.Password)
    if err != nil {
        return user, err
    }

    // Create the user model
    user = domain.User{
        ID:                primitive.NewObjectID(), 
        UserName:          input.UserName,
        Email:             input.Email,
        Password:          hashedPassword,
        Is_Admin:          false, 
        IsVerified:        false, 
        OAuthProvider:     "",
        OAuthID:           "",
        VerificationToken: "", 
    }

    err = u.AuthRepo.SaveUser(&user)
    if err != nil {
        return user, err
    }

    return user, nil
}

func (u *AuthUsecase) LoginUser(email, password string) (domain.User, string, string, error) {
    var user domain.User

    foundUser, err := u.AuthRepo.FindUserByEmail(email)
    if err != nil {
        return user, "", "", err
    }

    if foundUser == nil {
        return user, "", "", fmt.Errorf("user not found")
    }

    isMatch, err := u.PasswordSrv.ComparePassword(foundUser.Password, password)
    if err != nil {
        return user, "", "", err
    }

    if !isMatch {
        return user, "", "", fmt.Errorf("invalid password")
    }

    accessToken, err := u.TokenSrv.GenerateAccessToken(*foundUser)
    if err != nil {
        return user, "", "", err
    }

    refreshToken, err := u.TokenSrv.GenerateRefreshToken(*foundUser)
    if err != nil {
        return user, "", "", err
    }

    return *foundUser, accessToken, refreshToken, nil
}

func (u *AuthUsecase) RefreshTokens(refreshToken string) (string, string, error) {
    user, err := u.TokenSrv.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", "", err
    }

    newAccessToken, err := u.TokenSrv.GenerateAccessToken(*user)
    if err != nil {
        return "", "", err
    }

    newRefreshToken, err := u.TokenSrv.GenerateRefreshToken(*user)
    if err != nil {
        return "", "", err
    }

    return newAccessToken, newRefreshToken, nil
}

// func (u *AuthUsecase) OAuthSignUp(provider, token string) (domain.User, string, string, error) {
	
// }

// func (u *AuthUsecase) OAuthLogin(provider, token string) (domain.User, string, string, error) {
	
// }
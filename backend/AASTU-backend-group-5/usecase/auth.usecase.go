package usecase

// import (
// 	"errors"
// 	"time"

// 	"github.com/RealEskalate/blogpost/domain"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type AuthUsecase struct {
// 	AuthRepo      domain.AuthRepository
// 	PasswordSrv   domain.PasswordService
// 	TokenSrv      domain.TokenService
// 	OAuthSrv      domain.OAuthService
// }

// func NewAuthUsecase(authRepo domain.AuthRepository, passwordSrv domain.PasswordService, tokenSrv domain.TokenService, oauthSrv domain.OAuthService) *AuthUsecase {
// 	return &AuthUsecase{
// 		AuthRepo:    authRepo,
// 		PasswordSrv: passwordSrv,
// 		TokenSrv:    tokenSrv,
// 		OAuthSrv:    oauthSrv,
// 	}
// }

// func (u *AuthUsecase) RegisterUser(input domain.RegisterUser) (domain.User, string, string, error) {
	
// }

// func (u *AuthUsecase) OAuthSignUp(provider, token string) (domain.User, string, string, error) {
	
// }

// func (u *AuthUsecase) LoginUser(email, password string) (domain.User, string, string, error) {
	
// }

// func (u *AuthUsecase) OAuthLogin(provider, token string) (domain.User, string, string, error) {
	
// }

// func (u *AuthUsecase) RefreshTokens(refreshToken string) (string, string, error) {
	
// }

// func (u *AuthUsecase) Logout(userID string) error {
	
// }


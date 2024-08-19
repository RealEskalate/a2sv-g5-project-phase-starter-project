package usecases

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"time"

	emailverifier "github.com/AfterShip/email-verifier"
)

var Verifier = emailverifier.NewVerifier()

type loginUsecase struct {
	userRepository domain.UserRepository
	tokenService   domain.TokenInfrastructure
	contextTimeout time.Duration
	Env            *bootstrap.Env
}

// NewLoginUsecase creates a new instance of the LoginUsecase struct.
// It takes a userRepository of type domain.UserRepository, a tokenService of type domain.TokenInfrastructure,
// a timeout of type time.Duration, and an env pointer of type *bootstrap.Env.
// It returns a domain.LoginUsecase interface.
// The userRepository is responsible for handling user data storage and retrieval.
// The tokenService is responsible for generating and validating authentication tokens.
// The timeout specifies the maximum duration for a login operation.
// The env pointer provides access to the application's environment variables.
func NewLoginUsecase(userRepository domain.UserRepository, tokenService domain.TokenInfrastructure, timeout time.Duration, env *bootstrap.Env) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		tokenService:   tokenService,
		contextTimeout: timeout,
		Env:            env,
	}
}

// LoginWithIdentifier authenticates a user with the given identifier.
// It first verifies the identifier and then retrieves the user from the repository based on the identifier.
// If the identifier is a username, it searches for the user by username.
// If the identifier is an email, it searches for the user by email.
// It then creates access and refresh tokens for the user using the provided secrets and expiry times.
// The access token and refresh token are returned along with any error that occurred during the process.
func (lu *loginUsecase) LoginWithIdentifier(c context.Context, identifier string) (accessToken string, refreshToken string, err error) {
	_, err = Verifier.Verify(identifier)
	var user domain.User
	if err != nil {
		user, err = lu.userRepository.GetByUsername(c, identifier)
		if err != nil {
			return "", "", fmt.Errorf("user with this username not found")
		}
	} else {
		user, err = lu.userRepository.GetByEmail(c, identifier)
		if err != nil {
			return "", "", fmt.Errorf("user with this email not found")
		}
	}

	if !user.Verified {
		return "", "", fmt.Errorf("account not verified")
	}

	accessToken, refreshToken, err = lu.CreateAllTokens(c, &user, lu.Env.AccessTokenSecret, lu.Env.RefreshTokenSecret,
		lu.Env.AccessTokenExpiryMinute, lu.Env.RefreshTokenExpiryHour)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

// CreateAllTokens generates access and refresh tokens for the given user and saves them in the database.
// It returns the access token, refresh token, and any error encountered during the process.
func (lu *loginUsecase) CreateAllTokens(c context.Context, user *domain.User, accessSecret string, refreshSecret string,
	accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error) {
	accessToken, refreshToken, err = lu.tokenService.CreateAllTokens(user, accessSecret, refreshSecret, accessExpiry, refreshExpiry)

	if err != nil {
		return "", "", err
	}

	newUserUpdate := domain.UserUpdate{
		Access_Token:  &accessToken,
		Refresh_Token: &refreshToken,
	}

	_, err = lu.userRepository.UpdateUser(c, user.ID.Hex(), newUserUpdate)
	if err != nil {
		return "", "", fmt.Errorf("unable to save access and refresh tokens in databse")
	}

	return accessToken, refreshToken, nil

}

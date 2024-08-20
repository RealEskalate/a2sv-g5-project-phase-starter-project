package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type tokenService struct {
	accessTokenSecret, refreshTokenSecret string
	tokenRepository                       interfaces.RefreshTokenRepository
	userRepo                              interfaces.UserRepository
}

func NewTokenService(accessTokenSecret, refreshTokenSecret string, token_repo interfaces.RefreshTokenRepository, user_repo interfaces.UserRepository) interfaces.TokenService {
	return &tokenService{
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
		tokenRepository:    token_repo,
		userRepo:           user_repo,
	}
}

func (service *tokenService) InvalidateAccessToken(token string) (string, error) {

	//extract the claims from the token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	claims["exp"] = 0
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, err := newToken.SignedString([]byte(service.accessTokenSecret))
	if err != nil {
		return "", err
	}
	return expiredToken, nil
}

func (service *tokenService) InvalidateRefreshToken(token string) (string, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	claims["exp"] = 0
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, err := newToken.SignedString([]byte(service.refreshTokenSecret))
	if err != nil {
		return "", err
	}
	return expiredToken, nil
}

func (service *tokenService) GenerateAccessToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userName": user.Username,
			"userId": user.ID,
			"email":  user.Email,
			"role":   user.Role,
			"exp":    time.Now().Add(time.Hour).Unix(),
		})
	accessToken, err := token.SignedString(service.accessTokenSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *tokenService) GenerateRefreshToken(user *entities.User) (*entities.RefreshToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userName": user.Username,
			"userId": user.ID,
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	refreshToken, err := token.SignedString(service.refreshTokenSecret)
	if err != nil {
		return &entities.RefreshToken{}, err
	}

	// this may be moved to other layers
	refresh_token := entities.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	service.tokenRepository.CreateRefreshToken(&refresh_token)
	return &refresh_token, nil
}

func (service *tokenService) VerifyAccessToken(refresh, token string) (string, error) {

	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.accessTokenSecret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			err := service.VerifyRefreshToken(refresh)
			if err != nil {
				return "", err
			}
			userId, ok := claims["userId"].(string)
			if !ok {
				return "", fmt.Errorf("invalid user ID in claims")
			}
			user, err := service.userRepo.FindUserById(userId)
			if err != nil {
				return "", err
			}
			newToken, _ := service.GenerateAccessToken(user)
			return newToken, nil
		}
	}
	return "", nil
}

func (service *tokenService) VerifyRefreshToken(token string) error {
	refreshToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.refreshTokenSecret), nil
	})

	if err != nil {
		return err
	}
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			userId, ok := claims["userId"].(string)
			if !ok {
				return errors.New("invalid user ID in claims")
			}
			service.tokenRepository.DeleteRefreshTokenByUserId(userId)
			return errors.New("token is expired login again")
		}
		return nil
	}
	return nil
}

func (service *tokenService) GetClaimsFromToken(token string) map[string]string {
	Token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.refreshTokenSecret), nil
	})

	if err != nil {
		return map[string]string{}
	}

	if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
		resp := make(map[string]string)
		for key, value := range claims {
			resp[key] = fmt.Sprintf("%v", value)
		}
		return resp
	}
	return map[string]string{}
}

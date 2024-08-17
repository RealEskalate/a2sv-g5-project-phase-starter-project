package service

import "github.com/dgrijalva/jwt-go/v4"

type tokenService struct {
	accessTokenSecret, refreshTokenSecret string
}

func NewTokenService(accessTokenSecret string, refreshTokenSecret string ) *tokenService {
	return &tokenService{
		accessTokenSecret: accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
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
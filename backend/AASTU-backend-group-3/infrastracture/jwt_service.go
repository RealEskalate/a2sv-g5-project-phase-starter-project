package infrastracture

import (
	"group3-blogApi/config"
	"group3-blogApi/domain"
	"time"
	"github.com/golang-jwt/jwt"
)


func GenerateToken(user domain.User) (string, error) {
	var acessTokenSecret = []byte(config.EnvConfigs.JwtSecret)
	var AccessTokenExpiryHour = config.EnvConfigs.AccessTokenExpiryHour


	claims := domain.JwtCustomClaims{
		Authorized: true,
		UserID:     user.ID.Hex(),
		Role:       user.Role,
		Username:   user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(AccessTokenExpiryHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(acessTokenSecret)
	if err != nil {
		return "", err
	}
	return t, err
}


func RefreshToken(token string) (string, error) {
	var secret = []byte(config.EnvConfigs.JwtRefreshSecret)
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", err
	}
	return claims["user_id"].(string), nil
}
func GenerateRefreshToken(user *domain.User) (string, error) {
	var secret = []byte(config.EnvConfigs.JwtRefreshSecret)
	var expiry = config.EnvConfigs.RefreshTokenExpiryHour

	claims := domain.JwtCustomClaims{
		Authorized: true,
		UserID:     user.ID.Hex(),
		Role:       user.Role,
		Username:   user.Username,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
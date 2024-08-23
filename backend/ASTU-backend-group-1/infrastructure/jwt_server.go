package infrastructure

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user *domain.User, pwd string) (string, string, error) {
	configjwt, err := config.LoadConfig()
	if err != nil {
		return "", "", err
	}
	jwtSecret := []byte(configjwt.Jwt.JwtKey)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
		return "", "", errors.New("invalid username or password")
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &domain.Claims{
		ID:       user.ID,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsActive: user.IsActive,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	expirationTime = time.Now().Add(1 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

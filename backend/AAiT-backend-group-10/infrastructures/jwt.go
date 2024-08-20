package infrastructures

import (
	"errors"
	"time"

	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
)

func (uc *Infranstructure) GenerateToken(user *domain.User) (string, string, error) {
	// Define JWT claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration
		"admin": user.IsAdmin,
	}

	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		return "", "", err
	}

	// Create refresh token (valid for 7 days)
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func (uc *Infranstructure) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(uc.JWTSecret), nil
	})
}

func (uc *Infranstructure) GenerateResetToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Token valid for 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

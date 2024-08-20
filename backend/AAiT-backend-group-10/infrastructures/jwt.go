package infrastructures

import (
	"errors"
	"time"

	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	JwtSecret string
}

func (s *Jwt) GenerateToken(user *domain.User) (string, string, error) {
	// Define JWT claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration
		"admin": user.IsAdmin,
	}

	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}

	// Create refresh token (valid for 7 days)
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func (s *Jwt) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.JwtSecret), nil
	})
}

func (s *Jwt) GenerateResetToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Token valid for 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

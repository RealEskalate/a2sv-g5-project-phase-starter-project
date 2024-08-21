package infrastructures

import (
	"time"

	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	JwtSecret string
}

func (s *Jwt) GenerateToken(user *domain.User) (string, string, *domain.CustomError) {
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
		return "", "", domain.ErrTokenGenerationFailed
	}

	// Create refresh token (valid for 7 days)
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", domain.ErrRefreshTokenGenerationFailed
	}

	return tokenString, refreshTokenString, nil
}

func (s *Jwt) ValidateToken(token string) (*jwt.Token, *domain.CustomError) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(s.JwtSecret), nil
	})

	if err != nil {
		return nil, domain.ErrTokenParsingFailed
	}

	return parsedToken, nil
}

func (s *Jwt) GenerateResetToken(email string, code int64) (string, *domain.CustomError) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
		"code":  code,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", domain.ErrResetTokenGenerationFailed
	}

	return tokenString, nil
}

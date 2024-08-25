package infrastructures

import (
	"time"

	"aait.backend.g10/domain"
	"github.com/golang-jwt/jwt/v4"
)

type JwtService struct {
	JwtSecret string
}

func (s *JwtService) GenerateToken(user *domain.User) (string, string, *domain.CustomError) {
	// Define JWT claims
	if !user.Activated {
		return "", "", domain.ErrTokenGenerationFailed
	}
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration
		"is_admin": user.IsAdmin,
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

func (s *JwtService) ValidateToken(token string) (*jwt.Token, *domain.CustomError) {
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

func (s *JwtService) GenerateResetToken(email string, code int64) (string, *domain.CustomError) {
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
func (s *JwtService) GenerateActivationToken(email string) (string, *domain.CustomError) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", domain.ErrResetTokenGenerationFailed
	}

	return tokenString, nil
}

func (s *JwtService) CheckToken(authPart string) (*jwt.Token, *domain.CustomError) {
	token, err := jwt.Parse(authPart, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(s.JwtSecret), nil
	})

	if err != nil {
		return nil, domain.ErrTokenParsingFailed
	}

	return token, nil
}

func (s *JwtService) FindClaim(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

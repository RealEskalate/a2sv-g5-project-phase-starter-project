package infrastructures

import (
	"time"

	"aait.backend.g10/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	JwtSecret string
}	


func (s *JwtService) GenerateToken(user domain.User) (string, *domain.CustomError) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	token.Claims = jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
	}

	jwt_token, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", domain.ErrTokenGenerationFailed
	}
	return jwt_token, nil
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

func FindClaim(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

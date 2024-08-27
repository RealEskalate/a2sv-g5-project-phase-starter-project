package infrastructures

import (
	"fmt"
	"loan-management/internal/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID  string
	Email   string
	IsAdmin bool
}

func GenerateJWTToken(user domain.User, t time.Duration) (string, error) {
	claims := UserClaims{
		UserID:         user.ID,
		Email:          user.Email,
		IsAdmin:        user.IsAdmin,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(t).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secrete"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWTToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secrete"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

package infrastructure

import (
	domain "aait-backend-group4/Domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type services struct{}

func NewServices() domain.TokenInfrastructure {
	return &services{}
}

func (s *services) CreateAllTokens(user *domain.User, accessSecret string,
	refreshSecret string, accessExpriy int, refreshExpiry int) (accessToken string, refreshToken string, err error) {
	claims := domain.JwtCustomClaims{
		UserID:   user.ID.Hex(),
		Email:    user.Email,
		Username: user.Username,
		Role:     user.User_Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(accessExpriy)).Unix(),
		},
	}

	refreshClaims := domain.JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(refreshExpiry)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(accessSecret))
	if err != nil {
		return "", "", err
	}

	refresh, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(refreshSecret))
	if err != nil {
		return "", "", err
	}

	return token, refresh, nil
}

func (s *services) ValidateToken(tokenString string, secret string) (claims *domain.JwtCustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&domain.JwtCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, fmt.Errorf("the token is expired")
	}

	return claims, nil
}

func (s *services) ExtractRoleFromToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["Role"].(string), nil
}

func (s *services) CheckTokenExpiry(tokenString string, secret string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return false, nil
		}
	}

	return true, nil
}

func (s *services) ExtractClaims(tokenString string, secret string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims := make(map[string]interface{})
	claims["UserID"] = token.Claims.(jwt.MapClaims)["user_id"].(string)
	claims["UserName"] = token.Claims.(jwt.MapClaims)["user_name"].(string)
	claims["Role"] = token.Claims.(jwt.MapClaims)["role"].(string)
	claims["exp"] = token.Claims.(jwt.MapClaims)["exp"].(float64)

	return claims, nil
}

package infrastructure

import (
	"fmt"
	"strings"
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Env *config.Env
}

func NewJwtService(env *config.Env) interfaces.JwtService {
	return &JwtService{
		Env: env,
	}
}

func (j *JwtService) CreateAccessToken(user models.User, expTim int) (accessToken string, err error) {
	expTime := time.Now().Add(time.Minute * time.Duration(expTim)).Unix()
	secret := []byte(j.Env.JWT_SECRET)

	claims := &models.JWTCustome{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return t, err
}

func (j *JwtService) CreateRefreshToken(user models.User, expTim int) (refreshToken string, err error) {
	expTime := time.Now().Add(time.Hour * time.Duration(expTim)).Unix()
	secret := []byte(j.Env.JWT_SECRET)

	claims := &models.JWTCustome{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (j *JwtService) ValidateToken(tokenStr string) (*models.JWTCustome, error) {
	jwtSecret := []byte(j.Env.JWT_SECRET)

	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTCustome{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*models.JWTCustome)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}

func (j *JwtService) ValidateAuthHeader(authHeader string) ([]string, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	return authParts, nil
}

func (j *JwtService) CreateURLToken(user models.User, expTim int) (accessToken string, err error) {
	expTime := time.Now().Add(time.Minute * time.Duration(expTim)).Unix()
	secret := []byte(j.Env.JWT_SECRET)

	claims := &models.URLTokenCustom{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return t, err
}

func (j *JwtService) ValidateURLToken(tokenStr string) (*models.URLTokenCustom, error) {
	jwtSecret := []byte(j.Env.JWT_SECRET)

	token, err := jwt.ParseWithClaims(tokenStr, &models.URLTokenCustom{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*models.URLTokenCustom)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}

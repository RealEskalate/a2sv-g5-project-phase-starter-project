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
	// Set the expiration time for the refresh token, e.g., 7 days
	expTime := time.Now().Add(time.Hour * time.Duration(expTim)).Unix()
	secret := []byte(j.Env.JWT_SECRET)

	// Define the claims for the refresh token
	claims := &models.JWTCustome{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	// Create the token with the specified signing method and claims
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
<<<<<<< HEAD
=======


>>>>>>> origin/aait.backend.g5.bisrat.setup-db-and-user-repo

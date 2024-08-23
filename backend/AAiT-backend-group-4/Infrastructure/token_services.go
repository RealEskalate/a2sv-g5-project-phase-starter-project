package infrastructure

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenService struct {
	UserRepository domain.UserRepository
	Env            *bootstrap.Env
}

func NewTokenService(userRepository domain.UserRepository, env *bootstrap.Env) domain.TokenInfrastructure {
	return &tokenService{
		UserRepository: userRepository,
		Env:            env,
	}
}

// CreateAllTokens generates access and refresh tokens for a user.
// It takes the user object, access and refresh secrets, access and refresh expiry durations as input.
// It returns the access token, refresh token, and any error encountered during token generation.
func (s *tokenService) CreateAllTokens(user *domain.User, accessSecret string,
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

// ValidateToken validates the given token string using the provided secret key.
// It returns the claims extracted from the token if the token is valid and not expired.
// Otherwise, it returns an error indicating the reason for the validation failure.
func (s *tokenService) ValidateToken(tokenString string, secret string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwt.StandardClaims{})
	if err != nil {
		return false, fmt.Errorf("invalid token: %v", err)
	}

	// Verify the signing method
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return false, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return true, nil
}

// ExtractRoleFromToken extracts the role from a JWT token.
// It takes the token string and the secret key as input parameters.
// It returns the role as a string and an error if the token is invalid.
func (s *tokenService) ExtractRoleFromToken(tokenString string, secret string) (string, error) {
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

func (s *tokenService) ExtractUserIDFromToken(tokenString string) (string, error) {

	type JwtCustomClaims struct {
		UserID string `json:"user_id"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {

		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("invalid token")
}

// CheckTokenExpiry checks the expiry of a given token.
// It takes a token string and a secret as input parameters.
// It returns a boolean value indicating whether the token has expired or not, and an error if any.
func (s *tokenService) CheckTokenExpiry(tokenString string, secret string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	if token == nil {
		return false, errors.New("token is nil")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, errors.New("invalid exp claim")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

// ExtractClaims extracts the claims from a JWT token.
// It takes the token string and the secret key as input parameters.
// It returns a map[string]interface{} containing the extracted claims and an error if any.
func (s *tokenService) ExtractClaims(tokenString string, secret string) (map[string]interface{}, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Printf("Token parsing error: %v", err)
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		extractedClaims := make(map[string]interface{})

		log.Printf("claims: %v", claims)

		if userID, ok := claims["user_id"].(string); ok {
			extractedClaims["UserID"] = userID
		} else {
			return nil, fmt.Errorf("user_id claim is missing or not a string")
		}

		if userName, ok := claims["username"].(string); ok {
			extractedClaims["UserName"] = userName
		} else {
			return nil, fmt.Errorf("user_name claim is missing or not a string")
		}

		if role, ok := claims["role"].(string); ok {
			extractedClaims["Role"] = role
		} else {
			return nil, fmt.Errorf("role claim is missing or not a string")
		}

		if exp, ok := claims["exp"].(float64); ok {
			extractedClaims["exp"] = exp
		} else {
			return nil, fmt.Errorf("exp claim is missing or not a float64")
		}

		return extractedClaims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func (s *tokenService) UpdateTokens(id string) (accessToken string, refreshToken string, err error) {
	userFound, err := s.UserRepository.GetByID(context.Background(), id)
	if err != nil {
		return "", "", err
	}

	user := domain.User{
		ID:        userFound.ID,
		Email:     userFound.Email,
		Username:  userFound.Username,
		User_Role: userFound.User_Role,
		Password:  userFound.Password,
	}

	accessToken, refreshToken, err = s.CreateAllTokens(&user, s.Env.AccessTokenSecret,
		s.Env.RefreshTokenSecret, s.Env.AccessTokenExpiryMinute, s.Env.RefreshTokenExpiryHour)

	if err != nil {
		return "", "", err
	}

	_, err = s.UserRepository.UpdateTokens(context.Background(), userFound.ID.Hex(), accessToken, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

func (s *tokenService) RemoveTokens(id string) error {
	userFound, err := s.UserRepository.GetByID(context.Background(), id)
	if err != nil {
		return err
	}

	_, err = s.UserRepository.UpdateTokens(context.Background(), userFound.ID.Hex(), "", "")
	if err != nil {
		return err
	}

	return nil
}

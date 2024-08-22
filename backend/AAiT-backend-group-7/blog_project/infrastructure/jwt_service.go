package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"blog_project/domain"
)

func GenerateJWTAccessToken(user *domain.User, accessTokenSecret string, accessTokenExpiryHour int) (string, error) {
	// Set token expiry time
	expirationTime := time.Now().Add(time.Hour * time.Duration(accessTokenExpiryHour)).Unix()

	// Create JWT claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expirationTime,
		"iat":      time.Now().Unix(), // Issued at claim
	}

	// Create token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	accessToken, err := token.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateJWTRefreshToken(user *domain.User, refreshTokenSecret string, refreshTokenExpiryHours int) (string, error) {
	// Set token expiry time
	expirationTime := time.Now().Add(time.Hour * time.Duration(refreshTokenExpiryHours)).Unix()

	// Create JWT claims with minimal information
	claimsRefresh := jwt.MapClaims{
		"id":  user.ID,
		"exp": expirationTime,
		"iat": time.Now().Unix(), // Issued at claim
	}

	// Create the refresh token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	// Sign the token with the secret
	refreshToken, err := token.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
func IsAuthorized(requestToken string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

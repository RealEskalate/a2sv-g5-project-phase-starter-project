package infrastructure

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

// Genratetoken generates a JWT token for the given user and password.
func GenerateToken(user *domain.User, pwd string) (string, string, error) {
	configjwt,err := config.LoadConfig()
	if err != nil {
		return "", "", err
	}
	var jwtSecret = []byte(configjwt.JWTKey)

	// User login logic
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
		return "", "", errors.New("Invalid username or password")
	}

	// Access token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "",
			Subject:   "",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh token
	expirationTime = time.Now().Add(7 * 24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

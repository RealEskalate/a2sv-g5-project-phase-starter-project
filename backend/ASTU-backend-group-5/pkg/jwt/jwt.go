package jwt

import (
	"blogApp/internal/config"
	"blogApp/internal/domain"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	// Load environment variables from .env file (if present)
	_ = godotenv.Load(".env")

	// Retrieve the secret key from the environment variable
	secretKey = []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
}

// GenerateJWT generates a new JWT token and populates it with the user's ID and email

func GenerateJWT(userID, email, username, role string) (string, error) {
	ACCESS_TOKEN_LIFE_TIME := 2 * 24 * time.Hour

	Config, err := config.Load()
	if err == nil {
		integerLifeTime, err := strconv.Atoi(Config.ACCESS_TOKEN_LIFE_TIME)
		if err == nil {
			ACCESS_TOKEN_LIFE_TIME = time.Duration(integerLifeTime) * time.Hour
		}
	}

	expirationTime := time.Now().Add(ACCESS_TOKEN_LIFE_TIME)
	if role == "" {
		role = "user" // Set default role to "user" if no role is passed
	}
	claims := &domain.Claims{
		UserID:   userID,
		Email:    email,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	return signedToken, nil // Return the signed token and no error
}

// GenerateRefreshToken generates a new JWT token and populates it with the user's ID and email

func GenerateRefreshToken(userID, email, role, username string) (string, error) {
	REFRESH_TOKEN_LIFE_TIME := 7 * 24 * time.Hour
	Config, err := config.Load()
	if err == nil {
		integerLifeTime, err := strconv.Atoi(Config.REFRESH_TOKEN_LIFE_TIME)
		if err == nil {
			REFRESH_TOKEN_LIFE_TIME = time.Duration(integerLifeTime) * time.Hour
		}
	}
	expirationTime := time.Now().Add(REFRESH_TOKEN_LIFE_TIME)
	if role == "" {
		role = "user" // Set default role to "user" if no role is passed
	}
	claims := &domain.Claims{
		UserID:   userID,
		Email:    email,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	return signedToken, nil // Return the signed token and no error
}

// ValidateToken validates the JWT token string and returns the claims if the token is valid
func ValidateToken(tokenString string) (*domain.Claims, error) {
	claims := &domain.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
// VerifyGoogleToken
func VerifyGoogleToken(token string) (*domain.User, error) {
	// Verify the token
	user, err := verifyGoogleIDToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// verifyGoogleIDToken is a function that verifies the Google ID token
func verifyGoogleIDToken(token string) (*domain.User, error) {
	// Verify the token
	claims := &domain.Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	// Extract user information from the claims
	user := &domain.User{
		Email:    claims.Email,
		Role:     claims.Role,
	}

	return user, nil
}
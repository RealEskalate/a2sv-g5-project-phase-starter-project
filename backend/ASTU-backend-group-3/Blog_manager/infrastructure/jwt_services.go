package infrastructure

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Blacklist for invalidated tokens
var blacklistedTokens = struct {
	sync.RWMutex
	tokens map[string]bool
}{tokens: make(map[string]bool)}

// GenerateToken creates a JWT token
func GenerateToken(username, role string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// InvalidateToken adds a token to the blacklist
func InvalidateToken(tokenString string) {
	blacklistedTokens.Lock()
	defer blacklistedTokens.Unlock()
	blacklistedTokens.tokens[tokenString] = true
}

// IsTokenBlacklisted checks if a token is in the blacklist
func IsTokenBlacklisted(tokenString string) bool {
	blacklistedTokens.RLock()
	defer blacklistedTokens.RUnlock()
	return blacklistedTokens.tokens[tokenString]
}

// ParseUsernameToken parses the JWT and returns the username
func ParseUsernameToken(tokenString string) (username string, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || IsTokenBlacklisted(tokenString) {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", err
	}

	return claims.Username, nil
}

// GenerateRefreshToken creates a new refresh token
func GenerateRefreshToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

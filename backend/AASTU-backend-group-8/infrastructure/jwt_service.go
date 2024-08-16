package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTService interface
type JWTService interface {
	GenerateToken(userID, role string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, *Claims, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
}

// Claims struct to hold JWT claims
type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey        string
	issuer           string
	refreshSecretKey string
}

// NewJWTService creates a new JWTService
func NewJWTService(secretKey, issuer, refreshSecretKey string) JWTService {
	return &jwtService{
		secretKey:        secretKey,
		issuer:           issuer,
		refreshSecretKey: refreshSecretKey,
	}
}

// GenerateToken generates a new JWT token
func (j *jwtService) GenerateToken(userID, role string) (string, error) {
	claims := &Claims{
		ID:   userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // Shorter expiry time
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// GenerateRefreshToken generates a new refresh JWT token
func (j *jwtService) GenerateRefreshToken(userID string) (string, error) {
	claims := &Claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Longer expiry time for refresh token
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.refreshSecretKey))
}

// ValidateToken validates the given JWT token
func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

// ValidateRefreshToken validates the given refresh JWT token
func (j *jwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.refreshSecretKey), nil
	})
}

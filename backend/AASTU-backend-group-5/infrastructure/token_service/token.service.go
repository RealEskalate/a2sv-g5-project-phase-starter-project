package tokenservice

import (
	"errors"
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService_imp struct {
	AccessTokenSecret       string
	RefreshTokenSecret      string
	VerificationTokenSecret string
	PasswordResetSecret     string
}

func NewTokenService(accessSecret, refreshSecret, verificationSecret, resetSecret string) *TokenService_imp {
	return &TokenService_imp{
		AccessTokenSecret:       accessSecret,
		RefreshTokenSecret:      refreshSecret,
		VerificationTokenSecret: verificationSecret,
		PasswordResetSecret:     resetSecret,
	}
}

func (t *TokenService_imp) GenerateAccessToken(user domain.User) (string, error) {
	claims := domain.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.AccessTokenSecret))
}

func (t *TokenService_imp) GenerateRefreshToken(user domain.User) (string, error) {
	claims := domain.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 168).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.RefreshTokenSecret))
}

func (t *TokenService_imp) ValidateAccessToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.AccessTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	id, _ := claims["ID"].(string)
	iD, _ := primitive.ObjectIDFromHex(id)
	return &domain.User{
		ID: iD,
	}, nil
}


func (t *TokenService_imp) ValidateRefreshToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.RefreshTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	id, _ := claims["ID"].(string)
	iD, _ := primitive.ObjectIDFromHex(id)
	return &domain.User{
		ID: iD,
	}, nil
}

func (t *TokenService_imp) GenerateVerificationToken(user domain.User) (string, error) {
	claims := domain.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 48).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.VerificationTokenSecret))
}

func (t *TokenService_imp) ValidateVerificationToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.VerificationTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid verification token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	id, _ := claims["ID"].(string)
	iD, _ := primitive.ObjectIDFromHex(id)
	return &domain.User{
		ID: iD,
	}, nil
}

func (t *TokenService_imp) GeneratePasswordResetToken(user domain.User) (string, error) {
	claims := domain.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 1).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.PasswordResetSecret))
}

func (t *TokenService_imp) ValidatePasswordResetToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.PasswordResetSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid reset token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	id, _ := claims["ID"].(string)
	iD, _ := primitive.ObjectIDFromHex(id)
	return &domain.User{
		ID: iD,
	}, nil
}
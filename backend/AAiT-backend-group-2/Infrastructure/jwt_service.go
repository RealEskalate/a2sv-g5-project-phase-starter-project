package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string, email string, role string, duration time.Duration) (map[string]string, error)
	ValidateToken(token string) (map[string]string, error)
	RenewToken(refreshToken string) (string, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService(secretKey []byte) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}


func (service *jwtService) GenerateToken(userId string, email string, role string, duration time.Duration) (map[string]string, error) {
	accessToken, err := service.generateAccessToken(userId, email, role, duration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := service.generateRefreshToken(userId, duration)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (service *jwtService) ValidateToken(jwtToken string) (map[string]string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return service.secretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid jwt")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, userIDExists := claims["userID"].(string)
        userRole, userRoleExists := claims["role"].(string)
        if !userIDExists || !userRoleExists {
            return nil, fmt.Errorf("invalid jwt claims")
        }

        return map[string]string{
            "userID": userID,
            "role":   userRole,
        }, nil
    }

	if exp, ok := token.Claims.(jwt.MapClaims)["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Sub(time.Now()) > 0 {
			return nil, fmt.Errorf("token is expired")
		}
	}
	

	return nil, fmt.Errorf("invalid jwt claims")
}

func (service *jwtService) RenewToken(refreshToken string) (string, error) {
	claims, err := service.ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token")
	}

	userID := claims["userID"]

	// Create a new access token
	accessToken, err := service.generateAccessToken(userID, "", "", time.Hour*24)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}


func (service *jwtService) generateAccessToken(userID string, email string, role string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"email": email,
		"role": role,
		"exp": time.Now().Add(duration).Unix(),
	})

	accessToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *jwtService) generateRefreshToken(userID string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(duration).Unix(),
	})

	refreshToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
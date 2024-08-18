package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username string, role string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
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

func GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
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

func GetUsernameFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}
	return claims.Username, nil
}

func GetRoleFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(*Claims)
    if!ok {
        return "", jwt.ErrInvalidKey
    }
    return claims.Role, nil
}

type ResetClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}


func GenerateResetToken(username string, jwtKey []byte) (string, error) {
    expirationTime := time.Now().Add(10 * time.Minute)
    claims := &ResetClaims{
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

func ParseToken(tokenString string, jwtKey []byte) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Verify the token's signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, jwt.ErrInvalidKey
    }

    return claims, nil
}


func ParseResetToken(tokenString string, jwtKey []byte) (*ResetClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &ResetClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*ResetClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, fmt.Errorf("invalid token")
    }
}

package infrastructure

import (
    "Blog_Starter/domain"
    "Blog_Starter/utils"
    "fmt"
    "time"

    "github.com/golang-jwt/jwt"
)

// NewTokenManager implements the functions in the tokenutil.go
type NewTokenManager struct{}

// CreateAccessToken creates a new access token for a user
func (m *NewTokenManager) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
    exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
    claims := &utils.JwtCustomClaims{
        Email:         user.Email,
        UserID:        user.UserID.Hex(),
        Role:          user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: exp,
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }
    return t, nil
}

// CreateRefreshToken creates a new refresh token for a user
func (m *NewTokenManager) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
    claimsRefresh := &utils.JwtCustomRefreshClaims{
        UserID: user.UserID.Hex(),
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
    rt, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }
    return rt, nil
}

// ExtractIDFromToken extracts the user ID from a token
func (m *NewTokenManager) ExtractIDFromToken(requestToken string, secret string) (string, error) {
    token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if userID, ok := claims["user_id"].(string); ok {
            return userID, nil
        }
        return "", fmt.Errorf("user_id not found in token claims")
    }

    return "", fmt.Errorf("invalid token")
}

// ExtractRoleFromToken extracts the user role from a token
func (m *NewTokenManager) ExtractRoleFromToken(requestToken string, secret string) (string, error) {
    token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if role, ok := claims["role"].(string); ok {
            return role, nil
        }
        return "", fmt.Errorf("role not found in token claims")
    }

    return "", fmt.Errorf("invalid token")
}
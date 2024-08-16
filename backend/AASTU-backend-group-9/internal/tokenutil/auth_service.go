package tokenutil
import (
    "context"
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "blog/domain"
    "blog/repository"
)

type AuthService struct {
    tokenRepo repository.MongoTokenRepository
    jwtSecret string
}

func NewAuthService(tokenRepo repository.MongoTokenRepository, jwtSecret string) *AuthService {
    return &AuthService{
        tokenRepo: tokenRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) ValidateAccessToken(tokenString string) (*domain.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.jwtSecret), nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid access token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims.Valid() != nil {
        return nil, errors.New("invalid token claims")
    }

    return s.tokenRepo.FindTokenByAccessToken(context.Background(), tokenString)
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
    storedToken, err := s.tokenRepo.FindTokenByRefreshToken(context.Background(), refreshToken)
    if err != nil || storedToken == nil {
        return "", errors.New("invalid refresh token")
    }

    if time.Now().After(storedToken.ExpiresAt) {
        _ = s.tokenRepo.DeleteToken(context.Background(), storedToken.ID)
        return "", errors.New("refresh token expired")
    }

    newAccessToken, err := s.generateAccessToken(storedToken.UserID.Hex())
    if err != nil {
        return "", err
    }

    storedToken.AccessToken = newAccessToken
    storedToken.CreatedAt = time.Now()
    err = s.tokenRepo.SaveToken(context.Background(), storedToken)
    if err != nil {
        return "", err
    }

    return newAccessToken, nil
}

func (s *AuthService) generateAccessToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "id":        userID,
        "exp":       time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

package auth

import (
	"blog-api/domain"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authService struct {
	refreshTokenRepo domain.RefreshTokenRepository
	resetTokenRepo   domain.ResetTokenRepository
	accessSecret     string
	refreshSecret    string
	resetSecret      string
	accessExpiry     time.Duration
	refreshExpiry    time.Duration
	resetExpiry      time.Duration
}

func NewAuthService(refreshTokenRepo domain.RefreshTokenRepository, resetTokenRepo domain.ResetTokenRepository, accessSecret, refreshSecret, resetSecret string, accessExpiryHours, refreshExpiryHours, resetExpiryHours int) *authService {
	return &authService{
		refreshTokenRepo: refreshTokenRepo,
		resetTokenRepo:   resetTokenRepo,
		accessSecret:     accessSecret,
		refreshSecret:    refreshSecret,
		resetSecret:      resetSecret,
		accessExpiry:     time.Duration(accessExpiryHours) * time.Hour,
		refreshExpiry:    time.Duration(refreshExpiryHours) * time.Hour,
		resetExpiry:      time.Duration(resetExpiryHours) * time.Hour,
	}
}

func (a *authService) GenerateAccessToken(ctx context.Context, user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
		"exp":      time.Now().Add(a.accessExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.accessSecret))
}

func (a *authService) GenerateAndStoreRefreshToken(ctx context.Context, user domain.User) (string, error) {
	userID := user.ID.Hex()
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": user.Username,
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
		"exp":      time.Now().Add(a.refreshExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.refreshSecret))
	if err != nil {
		return "", err
	}
	err = a.refreshTokenRepo.StoreRefreshToken(ctx, userID, tokenString, time.Now().Add(a.refreshExpiry))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.refreshSecret), nil
	})
}
func (a *authService) ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.accessSecret), nil
	})
}

func (a *authService) RefreshTokens(ctx context.Context, accessToken string) (*domain.RefreshResponse, error) {

	token, err := a.ValidateAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid user id claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email claims")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid username claims")
	}
	isAdmin, ok := claims["isAdmin"].(bool)
	if !ok {
		return nil, errors.New("invalid isAdmin claims")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)

	// Convert userID to primitive.ObjectID
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	// time.Now().Add(a.accessExpiry).Unix()
	user := &domain.User{
		ID:       objectID,
		Username: username,
		Email:    email,
		IsAdmin:  isAdmin,
	}

	// Check if refresh token is in the database
	refreshToken, err := a.refreshTokenRepo.GetRefreshToken(ctx, userID)
	if err != nil {
		return nil, errors.New("refresh token not found or invalid")
	}

	_, err = VerifyToken(refreshToken, a.refreshSecret)
	if err != nil {
		a.DeleteRefreshToken(ctx, userID)
		return nil, errors.New("session expired, please login again")
	}

	// Generate new tokens
	newAccessToken, err := a.GenerateAccessToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &domain.RefreshResponse{
		AccessToken: newAccessToken,
	}, nil
}

func (a *authService) DeleteRefreshToken(ctx context.Context, userID string) error {
	return a.refreshTokenRepo.DeleteRefreshToken(ctx, userID)
}

func (a *authService) GeneratePasswordResetToken(ctx context.Context, email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(), // Token valid for 1 hour
		"type":  "password_reset",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.resetSecret))
	if err != nil {
		return "", err
	}

	resetToken := domain.PasswordResetToken{
		Token:  tokenString,
		Email:  email,
		Expiry: time.Now().Add(1 * a.resetExpiry),
		Used:   false,
	}
	err = a.resetTokenRepo.StoreResetToken(ctx, resetToken)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authService) ValidateResetToken(ctx context.Context, token string) (string, error) {
	// Validate the reset token and retrieve associated email
	email, err := a.resetTokenRepo.ValidateResetToken(ctx, token)
	if err != nil {
		return "", errors.New("invalid or expired password reset token")
	}

	// If validation is successful, return the associated email
	return email, nil
}

func (a *authService) InvalidateResetToken(ctx context.Context, token string) error {
	// Invalidate the reset token in the repository
	return a.resetTokenRepo.InvalidateResetToken(ctx, token)
}

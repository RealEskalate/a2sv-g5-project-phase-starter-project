package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserUsecase struct {
	repository   AuthRepository
	emailService infrastructure.EmailService
}

func NewAuthUserUsecase(repository AuthRepository, emailService infrastructure.EmailService) AuthServices {
	return &AuthUserUsecase{
		repository:   repository,
		emailService: emailService,
	}

}

func (au *AuthUserUsecase) Login(ctx context.Context, info LoginForm) (string, string, error) {
	// panic("not implemented") // TODO: Implement
	// var userInfo LoginForm
	// find the user name and match the hashed password with ith info.password
	user, err := au.repository.GetUserByUsername(ctx, info.Username)
	if err != nil {
		return "", "", ErrNoUserWithUsername
	}
	if !user.IsActive {
		return "", "", ErrAccountNotActivated
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password))
	if err != nil {
		return "", "", ErrIncorrectPassword
	}

	refToken, err := au.GenerateToken(user, "refresh")
	if err != nil {
		return "", "", err
	}
	accessToken, err := au.GenerateToken(user, "access")
	if err != nil {
		return "", "", err
	}
	err = au.repository.RegisterRefreshToken(ctx, user.ID, refToken)
	if err != nil {
		return "", "", err
	}

	return refToken, accessToken, nil
}

func (au *AuthUserUsecase) RegisterUser(ctx context.Context, user User) error {
	// var newUser User
	_, err := au.repository.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return ErrUserExistWithThisEmail
	}
	_, err = au.repository.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return ErrUserExistWithThisUsername
	}

	hashedpasswors, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedpasswors)
	user.IsActive = false
	// user.IsAdmin = false	activationLink := fmt.Sprintf("http://localhost/activate/%s/%s", user.ID, tokenString)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = au.repository.CreateUser(ctx, user)
	if err != nil {
		return ErrCantCreateUser
	}
	tokenString, err := au.GenerateActivateToken(user.Password, user.UpdatedAt)
	if err != nil {
		return err
	}
	from := os.Getenv("FROM")
	activationLink := fmt.Sprintf("http://localhost/activate/%s/%s", user.ID, tokenString)
	au.emailService.SendEmail(from, user.Email, "click the link to activate you account"+activationLink)

	return nil
}

func (au *AuthUserUsecase) UpdateProfile(ctx context.Context, user User) error {
	_, err := au.repository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthUserUsecase) Activate(ctx context.Context, userID string, token string) error {
	// refreshToekn, err := au.repository.GetRefreshToken(ctx, token)
	user, err := au.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	expectedToken, err := au.GenerateActivateToken(user.Password, user.UpdatedAt)
	if err != nil {
		return err
	}
	if expectedToken != token {
		return err
	}

	user.IsActive = true
	user.UpdatedAt = time.Now()

	_, err = au.repository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthUserUsecase) Logout(ctx context.Context, userID string) {
	token, err := au.repository.GetRefreshToken(ctx, userID)
	if err != nil {
		return
	}
	au.repository.DeleteRefreshToken(ctx, token)
}

func (au *AuthUserUsecase) GenerateActivateToken(hashedpassword string, updatedat time.Time) (string, error) {
	token := hashedpassword + updatedat.String()
	tokenbyte, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	token = base64.StdEncoding.EncodeToString(tokenbyte)
	return token, nil
}

func (au *AuthUserUsecase) GenerateToken(user User, tokenType string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"isadmin":  user.IsAdmin,
		"isactive": user.IsActive,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"type":     tokenType,
	}
	if tokenType == "refresh" {
		claims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

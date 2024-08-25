package auth

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
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

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedpassword)
	user.IsActive = false
	user.Email = strings.ToLower(user.Email)
	// user.IsAdmin = false	activationLink := fmt.Sprintf("http://localhost/activate/%s/%s", user.ID, tokenString)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// if the user is first user is make it admin and super admin
	count, err := au.repository.GetCollectionCount(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		user.IsAdmin = true
		user.IsSupper = true
	}
	id, err := au.repository.CreateUser(ctx, user)
	if err != nil {
		return ErrCantCreateUser
	}
	user.ID = id

	from := os.Getenv("FROM")
	tokenString := au.GenerateActivateToken(user.Password, user.UpdatedAt)

	activationLink := fmt.Sprintf("http://localhost/activate/%s/%s", user.ID, tokenString)
	au.emailService.SendEmail(from, user.Email, fmt.Sprintf("click the link to activate you account %s", activationLink), "Account Activation")

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
	user, err := au.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	expectedToken := au.GenerateActivateToken(user.Password, user.UpdatedAt)

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

func (au *AuthUserUsecase) GenerateActivateToken(hashedpassword string, updatedat time.Time) string {
	token := hashedpassword + updatedat.String()
	hasher := sha1.New()
	hasher.Write([]byte(token))

	token = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return token
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

func (au *AuthUserUsecase) PromoteUser(ctx context.Context, userID string) error {
	user, err := au.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if !user.IsAdmin {
		user.IsAdmin = true
	} else {
		return errors.New("the user was an admin")
	}
	_, err = au.repository.UpdateUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func (au *AuthUserUsecase) DemoteUser(ctx context.Context, userID string) error {
	user, err := au.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if !user.IsSupper {
		user.IsAdmin = false
	} else {
		return errors.New("you don't have the previlage to delete this super admin ")
	}
	_, err = au.repository.UpdateUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}

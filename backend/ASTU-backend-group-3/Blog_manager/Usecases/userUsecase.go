package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	Register(input Domain.RegisterInput) (*Domain.User, error)
	UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error
	DeleteUser(username string) error
	Login(c *gin.Context, LoginUser *Domain.LoginInput) (string, error)
	Logout(username string, tokenString string) error
	ForgotPassword(username string) (string, error)
	Reset(token string) (string, error)
	UpdatePassword(username string, newPassword string) error
}

type userUsecase struct {
	userRepo        Repository.UserRepository
	emailService    *infrastructure.EmailService
	passwordService *infrastructure.PasswordService
}

func NewUserUsecase(userRepo Repository.UserRepository, emailService *infrastructure.EmailService) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		emailService:    emailService,
		passwordService: infrastructure.NewPasswordService(),
	}
}

func (u *userUsecase) Register(input Domain.RegisterInput) (*Domain.User, error) {
	if strings.Contains(input.Username, "@") {
		return nil, errors.New("username must not contain '@'")
	}

	if _, err := u.userRepo.FindByUsername(input.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	if _, err := u.userRepo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := u.passwordService.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &Domain.User{
		Id:             primitive.NewObjectID(),
		Name:           input.Name,
		Username:       input.Username,
		Email:          input.Email,
		Password:       string(hashedPassword),
		ProfilePicture: input.ProfilePicture,
		Bio:            input.Bio,
		Gender:         input.Gender,
		Address:        input.Address,
		Role:           "user",
		IsActive:       true,
		PostsIDs:       []string{},
	}

	err = u.userRepo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	subject := "Welcome to Our Service!"
	body := fmt.Sprintf("Hi %s, welcome to our platform!", input.Username)
	err = u.emailService.SendEmail(input.Email, subject, body)
	if err != nil {
		return nil, fmt.Errorf("failed to send welcome email: %v", err)
	}

	return user, nil
}

func (u *userUsecase) UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error {
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return errors.New("user not found")
	}

	updateFields := bson.M{}

	if updatedUser.Username != "" {
		if strings.Contains(updatedUser.Username, "@") {
			return errors.New("username must not contain '@'")
		}
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Password != "" {
		hashedPassword, err := u.passwordService.HashPassword(updatedUser.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %v", err)
		}
		updateFields["password"] = hashedPassword
	}
	if updatedUser.ProfilePicture != "" {
		updateFields["profile_picture"] = updatedUser.ProfilePicture
	}
	if updatedUser.Bio != "" {
		updateFields["bio"] = updatedUser.Bio
	}
	if updatedUser.Address != "" {
		updateFields["address"] = updatedUser.Address
	}

	err = u.userRepo.Update(username, updateFields)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (u *userUsecase) DeleteUser(username string) error {
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	err = u.userRepo.Delete(username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (u *userUsecase) Login(c *gin.Context, LoginUser *Domain.LoginInput) (string, error) {
	user, err := u.userRepo.FindByUsername(LoginUser.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = u.passwordService.ComparePasswords(user.Password, LoginUser.Password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	accessToken, err := infrastructure.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %v", err)
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	err = u.userRepo.InsertToken(user.Username, accessToken, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to store tokens: %v", err)
	}

	return accessToken, nil
}

func (u *userUsecase) Logout(username string, tokenString string) error {
	err := u.userRepo.DeleteToken(username)
	if err != nil {
		return err
	}

	claims, err := infrastructure.ParseToken(tokenString, []byte("BlogManagerSecretKey"))
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return err
	}
	claims.ExpiresAt = time.Now().Unix()

	infrastructure.Blacklist.AddToBlacklist(tokenString)
	fmt.Println(infrastructure.Blacklist.IsTokenBlacklisted(tokenString))

	return nil
}

func (u *userUsecase) ForgotPassword(username string) (string, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	resetToken, err := infrastructure.GenerateResetToken(user.Username, []byte("BlogManagerSecretKey"))
	if err != nil {
		return "", err
	}

	return resetToken, nil
}

func (u *userUsecase) Reset(token string) (string, error) {

	claims, err := infrastructure.ParseResetToken(token, []byte("BlogManagerSecretKey"))
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return "", err
	}

	user, err := u.userRepo.FindByUsername(claims.Username)

	if err != nil {
		return "", errors.New("user not found")
	}
	access_token, err := infrastructure.GenerateJWT(user.Username, user.Role)

	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	return access_token, nil
}

func (u *userUsecase) UpdatePassword(username string, newPassword string) error {

	hashedPassword, err := u.passwordService.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	err = u.userRepo.Update(username, bson.M{"password": hashedPassword})
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

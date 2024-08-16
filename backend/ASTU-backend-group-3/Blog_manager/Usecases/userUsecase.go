package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	Register(input Domain.RegisterInput) (*Domain.User, error)
	UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error
	DeleteUser(username string) error
	Logout(username string) error
	Login(LoginUser *Domain.LoginInput) (string, error)
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
		passwordService: infrastructure.NewPasswordService(), // Initialize PasswordService here
	}
}

func (u *userUsecase) Register(input Domain.RegisterInput) (*Domain.User, error) {
	// Validate username: must not contain '@'
	if strings.Contains(input.Username, "@") {
		return nil, errors.New("username must not contain '@'")
	}

	// Check if username or email already exists
	if _, err := u.userRepo.FindByUsername(input.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	if _, err := u.userRepo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password using PasswordService
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

	// Save the user to the repository
	err = u.userRepo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	// Send a welcome email using EmailService
	subject := "Welcome to Our Service!"
	body := fmt.Sprintf("Hi %s, welcome to our platform!", input.Username)
	err = u.emailService.SendEmail(input.Email, subject, body)
	if err != nil {
		return nil, fmt.Errorf("failed to send welcome email: %v", err)
	}

	return user, nil
}

func (u *userUsecase) UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error {
	// Retrieve the existing user from the repository
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return errors.New("user not found")
	}

	// Validate the new username if it is being updated
	if updatedUser.Username != "" {
		if strings.Contains(updatedUser.Username, "@") {
			return errors.New("username must not contain '@'")
		}
	}

	// Prepare update fields
	updateFields := bson.M{}

	if updatedUser.Username != "" {
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Password != "" {
		// Hash the new password before updating using PasswordService
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

	// Call repository to update the user
	err = u.userRepo.Update(username, updateFields)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (u *userUsecase) DeleteUser(username string) error {
	// Check if the user exists before attempting to delete
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Call repository to delete the user
	err = u.userRepo.Delete(username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (u *userUsecase) Login(LoginUser *Domain.LoginInput) (string, error) {
	user, err := u.userRepo.FindByUsername(LoginUser.Username)
	if err != nil {
		return "", err
	}

	storedPassword := user.Password

	// Compare the stored hashed password with the provided one
	err = u.passwordService.ComparePasswords(storedPassword, LoginUser.Password)
	if err != nil {
		return "", err
	}

	accessToken, err := infrastructure.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(user.Username)
	if err != nil {
		return "", err
	}

	err = u.userRepo.InsertToken(user.Username, accessToken, refreshToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *userUsecase) Logout(username string) error {
	return u.userRepo.DeleteToken(username)
}

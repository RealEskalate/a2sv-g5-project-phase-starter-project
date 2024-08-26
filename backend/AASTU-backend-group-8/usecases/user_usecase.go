package usecases

import (
	"errors"
	"fmt"
	"meleket/domain"
	"meleket/infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo  domain.UserRepositoryInterface
	tokenRepo domain.TokenRepositoryInterface
	jwtSvc    infrastructure.JWTService
	// passwordSvc infrastructure.PasswordService
	// emailSvc	infrastructure.EmailService
}

func NewUserUsecase(ur domain.UserRepositoryInterface, tr domain.TokenRepositoryInterface, jr infrastructure.JWTService) *UserUsecase {
	return &UserUsecase{
		userRepo:  ur,
		tokenRepo: tr,
		jwtSvc:    jr,
	}
}

// Register registers a new user
func (u *UserUsecase) Register(user *domain.User) error {
	// Hash the user's password before storing it
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save the user to the repository
	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}

func (u *UserUsecase) GetUserByEmail(email *string) (*domain.User, error) {
	return u.userRepo.GetUserByEmail(*email)
}

// Login authenticates a user and returns JWT and refresh tokens if successful
func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
	fmt.Println("authuser: ", authUser)
	user, err := u.userRepo.GetUserByUsername(authUser.Username)
	if err != nil {
		return "", "", err
	}
	if user.IsOAuth {
		return "", "", errors.New("invalid username or password")
	}

	fmt.Println("user: ", user)

	if err := infrastructure.CheckPasswordHash(user.Password, authUser.Password); err != nil {
		return "", "", errors.New("invalid username or password2")
	}

	// Generate JWT and refresh tokens for the authenticated user
	token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.jwtSvc.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshedTokenClaim := &domain.RefreshToken{
		UserID:    user.ID,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	// Save the refresh token in the database
	err = u.tokenRepo.SaveRefreshToken(refreshedTokenClaim)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *UserUsecase) LoginWithProvider(user *domain.User) (string, string, error) {
	// Generate JWT and refresh tokens for the authenticated user
	token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	userFetched, err := u.userRepo.GetUserByEmail(user.Email)

	refreshToken, err := u.jwtSvc.GenerateRefreshToken(userFetched.ID, userFetched.Role)
	if err != nil {
		return "", "", err
	}

	refreshedTokenClaim := &domain.RefreshToken{
		UserID:    user.ID,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	// Save the refresh token in the database
	err = u.tokenRepo.SaveRefreshToken(refreshedTokenClaim)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *UserUsecase) UpdateUser(username, newPassword string) error {
	existingUser, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	hashedPassword, err := infrastructure.HashPassword(newPassword)
	if err != nil {
		return err
	}
	existingUser.Password = hashedPassword

	return u.userRepo.UpdateUser(username, existingUser)
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.userRepo.GetAllUsers()
	return users, err
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(objectID)
}

func (u *UserUsecase) PromoteToAdmin(username string) error {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	if user.Role == "admin" || user.Role == "root" {
		return errors.New("user is already an admin")
	}

	return u.userRepo.UpdateRole(username, "admin")
}

func (u *UserUsecase) DemoteToUser(username string) error {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	if user.Role == "user" {
		return errors.New("user is already a user")
	}

	if user.Role == "root" {
		return errors.New("cannot demote a root user")
	}

	return u.userRepo.UpdateRole(username, "user")
}


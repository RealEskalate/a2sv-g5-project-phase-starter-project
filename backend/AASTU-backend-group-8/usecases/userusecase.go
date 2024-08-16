package usecases

import (
	"errors"
	"meleket/domain"
	"meleket/infrastructure"
	"meleket/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo    repository.UserRepositoryInterface
	passwordSvc infrastructure.PasswordService
	jwtSvc      infrastructure.JWTService
}

func NewUserUsecase(ur repository.UserRepositoryInterface, ps infrastructure.PasswordService, js infrastructure.JWTService) *UserUsecase {
	return &UserUsecase{
		userRepo:    ur,
		passwordSvc: ps,
		jwtSvc:      js,
	}
}

// Register registers a new user
func (u *UserUsecase) Register(user *domain.User) error {
	// Hash the user's password before storing it
	hashedPassword, err := u.passwordSvc.HashPassword(user.Password)
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

// Login authenticates a user and returns JWT and refresh tokens if successful
func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
	// Retrieve the user by username
	user, err := u.userRepo.FindByUsername(&authUser.Username)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Compare the provided password with the stored hashed password
	if err := u.passwordSvc.CheckPasswordHash(user.Password, authUser.Password); err != nil {
		return "", "", errors.New("invalid username or password")
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

	// Save the refresh token in the repository
	err = u.userRepo.SaveToken(authUser.Username, refreshToken)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

// DeleteRefreshToken deletes the refresh token for a user
func (u *UserUsecase) DeleteRefreshToken(userID primitive.ObjectID) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	return u.userRepo.DeleteToken(user.userID)
}

// ForgotPassword handles the forgot password logic
func (u *UserUsecase) ForgotPassword(email *string) error {
	_, err := u.userRepo.FindByUsername(email)
	if err != nil {
		return errors.New("email not found")
	}

	// Generate and send a password reset token logic would go here
	return nil
}

// GetProfile retrieves a user's profile by ID
func (u *UserUsecase) GetProfile(objectID primitive.ObjectID) (*domain.User, error) {
	user, err := u.userRepo.FindByID(objectID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateProfile updates a user's profile
func (u *UserUsecase) UpdateProfile(objectID primitive.ObjectID, profile *domain.Profile) (*domain.User, error) {
	user, err := u.userRepo.FindByID(objectID)
	if err != nil {
		return nil, err
	}

	user.Name = profile.Bio
	user.avatar_url = profile.AvatarURL

	updatedUser, err := u.userRepo.UpdateProfile(objectID.Hex(), user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// GetAllUsers retrieves all users from the repository
func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
	return u.userRepo.Delete(objectID)
}

// RefreshToken refreshes a user's JWT token
func (u *UserUsecase) RefreshToken(refreshToken *domain.RefreshToken) (string, error) {
	storedToken, err := u.userRepo.FindRefreshToken(refreshToken.UserID)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// Assuming that the storedToken contains userID and Role, you would generate a new token
	newToken, err := u.jwtSvc.GenerateToken(storedToken.UserID, storedToken.ExpiresAt.String())
	if err != nil {
		return "", err
	}

	return newToken, nil
}

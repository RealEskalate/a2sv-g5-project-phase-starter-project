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
	userRepo    domain.UserRepositoryInterface
	tokenRepo  	domain.TokenRepositoryInterface
	jwtSvc      infrastructure.JWTService
	// passwordSvc infrastructure.PasswordService
	// emailSvc	infrastructure.EmailService
}

// func NewUserUsecase(ur domain.UserRepositoryInterface,js infrastructure.JWTService) *UserUsecase {  //ps infrastructure.PasswordService, js infrastructure.JWTService)
// 	return &UserUsecase{
// 		userRepo:    ur,
// 		jwtSvc:      js,
// 		// passwordSvc: ps,
// 		// emailSvc: 	 es,
// 	}
// }

func NewUserUsecase(ur domain.UserRepositoryInterface,tr domain.TokenRepositoryInterface, jr infrastructure.JWTService) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
		tokenRepo: tr,
		jwtSvc: jr,
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
	return u.userRepo.GetUserByEmail(email)
}

// Login authenticates a user and returns JWT and refresh tokens if successful
func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
	fmt.Println("authuser: ", authUser)
	user, err := u.userRepo.GetUserByUsername(authUser.Username)
	if err != nil {
		return "", "", err
	}

	if err := infrastructure.CheckPasswordHash(user.Password, authUser.Password); err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Generate JWT and refresh tokens for the authenticated user
	// token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
	token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.jwtSvc.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshedTokenClaim := &domain.RefreshToken{
		UserID: user.ID,
		Role: user.Role,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	// Save the refresh token in the database
	err = u.tokenRepo.SaveRefreshToken(refreshedTokenClaim)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *UserUsecase) GetProfile(objectID primitive.ObjectID) (*domain.Profile, error) {
	return nil, nil
}

func (u *UserUsecase) UpdateProfile(objectID primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
	updatedProfile, err := u.userRepo.UpdateProfile(objectID, profile)
	return updatedProfile, err
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.userRepo.GetAllUsers()
	return users, err
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(objectID)
}



// DeleteRefreshToken deletes the refresh token for a user
func (u *UserUsecase) DeleteRefreshToken(userID primitive.ObjectID) error {
	// user, err := u.userRepo.GetUserByID(userID)
	// if err != nil {
	// 	return err
	// }

	// return u.userRepo.DeleteToken(user.ID)
	return nil
}



// // UpdateProfile updates a user's profile
// func (u *UserUsecase) UpdateProfile(objectID primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
// 	// user, err := u.userRepo.FindByID(objectID)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// user.Name = profile.Bio
// 	// user.avatar_url = profile.AvatarURL

// 	updatedprofile, err := u.userRepo.UpdateProfile(objectID, profile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return updatedprofile, nil
// }

package usecase

import (
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	userRepo       domain.UserRepository
	TokenManager   utils.TokenManager
	contextTimeout time.Duration
	Env            *config.Env
}

func NewSignUpUsecase(userRepo domain.UserRepository, tokenManager utils.TokenManager, env *config.Env, timeout time.Duration) domain.SignupUsecase {
	return &SignupUsecase{
		userRepo:       userRepo,
		TokenManager:   tokenManager,
		Env:            env,
		contextTimeout: timeout,
	}
}

// CreateUser implements domain.SignupUsecase.
func (s *SignupUsecase) CreateUser(c context.Context, user *domain.UserSignUp) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Email = strings.ToLower(user.Email)

	userCreate, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		userCreate = &domain.User{
			UserID:      primitive.NewObjectID(),
			Username:    user.Username,
			Password:    user.Password,
			Email:       user.Email,
			CreatedAt:   time.Now(),
			IsActivated: false,
		}

		_, err = s.userRepo.CreateUser(ctx, userCreate)
		if err != nil {
			return nil, err
		}
	} else if !userCreate.IsActivated {
		userCreate.Email = user.Email
		userCreate.Username = user.Username
		userCreate.Password = user.Password
		err = s.userRepo.UpdateSignup(ctx, userCreate)
		if err != nil {
			return nil, err
		}
	} else if userCreate.IsActivated {
		return nil, errors.New("user already exists")
	}
	return userCreate, nil

}

// VerifyEmail implements domain.SignupUsecase.
func (s *SignupUsecase) VerifyEmail(c context.Context, req *domain.VerifyEmailRequest) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user.IsActivated {
		return nil, errors.New("user already activated")
	}

	user.IsActivated = true
	err = s.userRepo.UpdateSignup(ctx, user)
	if err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		Name:           user.Name,
		Bio:            user.Bio,
		ContactInfo:    user.ContactInfo,
		Role:           user.Role,
		IsActivated:    user.IsActivated,
		ProfilePicture: user.ProfilePicture,
	}, nil

}

// ResendOTP implements domain.SignupUsecase.
func (s *SignupUsecase) ResendOTP(c context.Context, req *domain.ResendOTPRequest) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if user.IsActivated {
		return errors.New("failed to resend otp. User account already activated")
	}

	return nil
}

// CreateTokens implements domain.SignupUsecase.
func (s *SignupUsecase) CreateTokens(c context.Context, user *domain.User) (*domain.TokenResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	accessToken, err := s.TokenManager.CreateAccessToken(user, s.Env.AccessTokenSecret, s.Env.AccessTokenExpiryHour)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.TokenManager.CreateRefreshToken(user, s.Env.RefreshTokenSecret, s.Env.RefreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	userID := user.UserID.Hex()
	_, err = s.userRepo.UpdateToken(ctx, accessToken, refreshToken, userID)
    if err != nil {
        return nil, err
    }

	return &domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}


func (s *SignupUsecase) HandleFederatedSignup(c context.Context, token string)  (*domain.User, error) {
    // Verify the token with the federated identity provider (e.g., Google)
    userInfo, err := utils.VerifyFederatedToken(token, s.Env.GoogleClientID)
    if err != nil {
        return nil, err
    }

    userInfo.Email = strings.ToLower(userInfo.Email)

    // Check if the user already exists in the database
    existingUser, err := s.userRepo.GetUserByEmail(c, userInfo.Email)
    if err != nil {
        // User doesn't exist, create a new user
        newUser := domain.User{
            UserID:        primitive.NewObjectID(),
            Name:          userInfo.Name,
            Username:      userInfo.Name, // Set a default username or customize based on your application
            Email:         userInfo.Email,
            Password:      "",   // You may leave this empty or handle it differently for federated signup
            CreatedAt:     time.Now(),
            IsActivated:   true,                   // Assuming the user is activated upon federated signup
            ProfilePicture: userInfo.ProfilePicUrl, // Assuming the federated provider provides a profile picture URL
        }

        // Save the new user to the database
        _, err := s.userRepo.CreateUser(c, &newUser)
        if err != nil {
            return nil, err
        }

        return &newUser, nil
    }

    return existingUser, nil
}
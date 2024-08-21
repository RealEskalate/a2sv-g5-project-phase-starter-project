package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	userRepo domain.UserRepository
	contextTimeout time.Duration
}

func NewSignUpUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &SignupUsecase{
		userRepo: userRepo,
		contextTimeout: timeout,
		
	}
}

// CreateUser implements domain.SignupUsecase.
func (s *SignupUsecase) CreateUser(c context.Context, user *domain.UserSignUp) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	user.Email = strings.ToLower(user.Email)
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	
	//TODO: validation check
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	

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
		return errors.New( "failed to resend otp. User account already activated")
	}

	return nil
}



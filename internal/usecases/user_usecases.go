package usecases

import (
	"fmt"
	"loan-management/internal/domain"
	"loan-management/internal/repositories"
	"loan-management/pkg/infrastructures"
	"time"

	"github.com/sv-tools/mongoifc"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(db mongoifc.Database) domain.UserUsecases {
	repo := repositories.NewUserRepository(db)
	return &userUsecase{userRepository: repo}
}

func (uc *userUsecase) Register(user domain.User) (domain.User, error) {
	hashedPassword, err := infrastructures.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPassword
	expirationTime := time.Now().Add(1 * time.Hour)
	verificationToken, err := infrastructures.GenerateVerificationToken(user.Email, expirationTime)
	if err != nil {
		return domain.User{}, fmt.Errorf("generating token:", err)
	}
	go infrastructures.SendVerificationEmail(user.Email, verificationToken)
	return uc.userRepository.Create(user)
}

func (uc *userUsecase) VerifyEmail(token, email string) error {
	user, err := uc.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}
	if err := infrastructures.ValidateVerificationToken(token, email); err != nil {
		return err
	}
	uc.userRepository.Update(user.ID, domain.User{IsActive: true})
	return nil
}

func (uc *userUsecase) Login(email, password string) (domain.User, error) {
	user, err := uc.userRepository.GetByEmail(email)
	if err != nil {
		return user, err
	}
	if !user.IsActive {
		return domain.User{}, fmt.Errorf("user account is not activated")
	}
	if infrastructures.ComparePassword(user.Password, password) {
		return user, nil
	}
	return domain.User{}, fmt.Errorf("incorrect password or email")
}

func (uc *userUsecase) GetProfile(userID string) (domain.User, error) {
	return uc.userRepository.GetByID(userID)
}

func (uc *userUsecase) ForgetPassword(email string) (string, error) {
	return "", nil
}

func (uc *userUsecase) ResetPassword(token, newPassword string) (string, error) {
	return "", nil
}

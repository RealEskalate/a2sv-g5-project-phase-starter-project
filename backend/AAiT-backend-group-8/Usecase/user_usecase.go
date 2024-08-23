package usecase

import (
	domain "AAiT-backend-group-8/Domain"
	interfaces "AAiT-backend-group-8/Interfaces"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func generateVerificationToken() string {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(token) + time.Now().Format("20060102150405")
}
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

type UserUseCaseImpl struct {
	userRepository  interfaces.IUserRepository
	TokenService    interfaces.ITokenService
	PasswordService interfaces.IPasswordService
	TokenRepo       interfaces.ITokenRepository
	MailService     interfaces.IMailService
}

func NewUserUseCase(userRepo interfaces.IUserRepository, ts interfaces.ITokenService, ps interfaces.IPasswordService, tr interfaces.ITokenRepository, ms interfaces.IMailService) *UserUseCaseImpl {
	return &UserUseCaseImpl{userRepository: userRepo, TokenService: ts, PasswordService: ps, TokenRepo: tr, MailService: ms}
}

func (userUseCase *UserUseCaseImpl) VerifyEmail(token string) error {
	user, err := userUseCase.userRepository.GetUserByVerificationToken(token)
	if err != nil {
		return err
	}

	if user.Verified {
		return errors.New("user is already verified")
	}

	user.Verified = true
	user.VerificationToken = ""
	err = userUseCase.userRepository.VerifyUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (userUseCase *UserUseCaseImpl) RegisterUser(user *domain.User) error {
	// Check if this is the first user
	userCount, err := userUseCase.userRepository.GetUserCount()
	if err != nil {
		return err
	}

	if userCount == 0 {
		user.Role = "super-admin"
	} else {
		user.Role = "user" // Default role for non-first users
	}

	// Check if email already exists
	existingUser, err := userUseCase.userRepository.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Set other user details
	user.CreatedAt = time.Now()
	user.Verified = false
	user.VerificationToken = generateVerificationToken()

	// Hash password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = userUseCase.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	err = userUseCase.MailService.SendVerificationEmail(user.Email, user.VerificationToken)
	if err != nil {
		return err
	}

	return nil
}

func (userUseCase *UserUseCaseImpl) GetSingleUser(email string) (*domain.User, error) {
	var user *domain.User

	user, err := userUseCase.userRepository.GetUserByEmail(email)

	return user, err
}

func (userUseCase *UserUseCaseImpl) RefreshToken(email, refresher string) (string, error) {
	//Check the validity of the refresher token
	_, err := userUseCase.TokenService.ValidateToken(refresher)

	if err != nil {
		return "", err
	}
	existingRefresher, err := userUseCase.TokenRepo.GetRefresher(email)

	if err != nil {
		return "", err
	}

	if existingRefresher != refresher {
		return "", errors.New("invalid refresher token")
	}

	var user *domain.User
	user, err = userUseCase.userRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	//generate a new token
	tokenExp := time.Now().Add(time.Minute * 5).Unix()
	token, err := userUseCase.TokenService.GenerateToken(email, user.Id, user.Role, user.Name, tokenExp)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (userUseCase *UserUseCaseImpl) Login(email string, password string) (string, string, error) {
	//Get user's hashedPassword from the database
	fmt.Print("email and password", email, password)
	user, err := userUseCase.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("incorrect email or password")
	}

	if !user.Verified {
		return "", "", errors.New("not a verified user")
	}

	hashedPassword := user.Password

	err = userUseCase.PasswordService.VerifyPassword(hashedPassword, password)
	if err != nil {
		return "", "", errors.New("incorrect email or password")
	}

	//Generate a token for the user
	tokenExp := time.Now().Add(time.Hour * 50).Unix()
	token, err := userUseCase.TokenService.GenerateToken(user.Email, user.Id, user.Role, user.Name, tokenExp)

	if err != nil {
		return "", "", err
	}

	refresherExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refresher, err := userUseCase.TokenService.GenerateToken(user.Email, user.Id, user.Role, user.Name, refresherExp)

	if err != nil {
		return "", "", err
	}
	credentials := domain.Credential{Email: email, Refresher: refresher}
	err = userUseCase.TokenRepo.InsertRefresher(credentials)

	if err != nil {
		return "", "", err
	}

	return token, refresher, nil
}

func (userUseCase *UserUseCaseImpl) GenerateResetPasswordToken(email string) error {
	user, err := userUseCase.userRepository.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	resetToken, err := userUseCase.TokenService.GenerateToken(user.Email, user.Id, "reset_password", "", time.Now().Add(1*time.Hour).Unix())
	if err != nil {
		return err
	}

	err = userUseCase.MailService.SendPasswordResetEmail(user.Email, resetToken)
	if err != nil {
		return err
	}

	return nil
}

func (userUseCase *UserUseCaseImpl) StoreToken(token string) error {
	claims, err := userUseCase.TokenService.ValidateToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token payload")
	}

	err = userUseCase.userRepository.StoreResetToken(email, token)
	if err != nil {
		return err
	}

	return nil
}

func (userUseCase *UserUseCaseImpl) ResetPassword(token string, newPassword string) error {
	claims, err := userUseCase.TokenService.ValidateToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token payload")
	}

	storedToken, err := userUseCase.userRepository.GetResetTokenByEmail(email)
	if err != nil {
		return err
	}

	if storedToken != token {
		return errors.New("invalid or mismatched token")
	}

	hashedPassword, err := userUseCase.PasswordService.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = userUseCase.userRepository.UpdatePasswordByEmail(email, hashedPassword)
	if err != nil {
		return err
	}

	err = userUseCase.userRepository.InvalidateResetToken(email)
	if err != nil {
		return err
	}

	return nil
}

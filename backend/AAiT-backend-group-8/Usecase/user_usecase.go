package Usecase

import (
	domain "AAiT-backend-group-8/Domain"

	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	userRepository  domain.IUserRepository
	TokenService    domain.ITokenService
	PasswordService domain.IPasswordService
	TokenRepo       domain.ITokenRepository
	MailService     domain.IMailService
}

func NewUserUseCase(userRepo domain.IUserRepository, ts domain.ITokenService, ps domain.IPasswordService, tr domain.ITokenRepository, ms domain.IMailService) *UserUseCaseImpl {
	return &UserUseCaseImpl{userRepository: userRepo, TokenService: ts, PasswordService: ps, TokenRepo: tr, MailService: ms}
}

func (u *UserUseCaseImpl) VerifyEmail(token string) error {
	user, err := u.userRepository.GetUserByVerificationToken(token)
	if err != nil {
		return err
	}

	if user.Verified {
		return errors.New("user is already verified")
	}

	user.Verified = true
	user.VerificationToken = ""
	err = u.userRepository.VerifyUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCaseImpl) RegisterUser(user *domain.User) error {
	// Check if this is the first user
	userCount, err := u.userRepository.GetUserCount()
	if err != nil {
		return err
	}

	if userCount == 0 {
		user.Role = "superadmin"
	} else {
		user.Role = "user" // Default role for non-first users
	}

	// Check if email already exists
	existingUser, err := u.userRepository.GetUserByEmail(user.Email)
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

	err = u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	err = u.MailService.SendVerificationEmail(user.Email, user.VerificationToken)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCaseImpl) GetSingleUser(email string) (*domain.User, error) {
	var user *domain.User

	user, err := uuc.userRepository.GetUserByEmail(email)

	return user, err
}

func (uuc *UserUseCaseImpl) RefreshToken(email, refresher string) (string, error) {
	//Check the validity of the refresher token
	err := uuc.TokenService.ValidateToken(refresher)

	if err != nil {
		return "", err
	}
	//Grasp the user's refresher token from the database
	existingRefresher, err := uuc.TokenRepo.GetRefresher(email)

	if err != nil {
		return "", err
	}

	//cross-check the given refresher with the existing one
	if existingRefresher != refresher {
		return "", errors.New("invalid refresher token")
	}

	//get user data from the db
	var user *domain.User
	user, err = uuc.userRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	//generate a new token
	tokenExp := time.Now().Add(time.Minute * 5).Unix()
	token, err := uuc.TokenService.GenerateToken(email, user.Id, user.Role, user.Name, tokenExp)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (uuc *UserUseCaseImpl) Login(email string, password string) (string, string, error) {
	//Get user's hashedPassword from the database
	user, err := uuc.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("incorrect email or password")
	}

	if !user.Verified {
		return "", "", errors.New("not a verified user")
	}

	hashedPassword := user.Password

	//Verify the password and the hashedPassword alignment
	err = uuc.PasswordService.VerifyPassword(hashedPassword, password)
	if err != nil {
		return "", "", errors.New("incorrect email or password")
	}

	//Generate a token for the user
	tokenExp := time.Now().Add(time.Minute * 5).Unix()
	token, err := uuc.TokenService.GenerateToken(user.Email, user.Id, user.Role, user.Name, tokenExp)

	if err != nil {
		return "", "", err
	}

	//Define and Generate a refresher token for the user
	refresherExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refresher, err := uuc.TokenService.GenerateToken(user.Email, user.Id, user.Role, user.Name, refresherExp)

	if err != nil {
		return "", "", err
	}
	//Store the refresher token in the database
	credentials := domain.Credential{Email: email, Refresher: refresher}
	err = uuc.TokenRepo.InsertRefresher(credentials)

	if err != nil {
		return "", "", err
	}

	//return the token and the refresher token
	return token, refresher, nil
}

package usecase

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
	//get user data from the db
	var user *domain.User
	user, err := uuc.userRepository.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("invalid email")
	}

	//Check the validity of the refresher token
	_, err = uuc.TokenService.ValidateToken(refresher)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	//Grasp the user's refresher token from the database
	err = uuc.TokenRepo.CheckRefresher(email, refresher)

	if err != nil {
		return "", errors.New("invalid refresher token")
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

func (uuc *UserUseCaseImpl) GenerateResetPasswordToken(email string) error {
	user, err := uuc.userRepository.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	resetToken, err := uuc.TokenService.GenerateToken(user.Email, user.Id, "reset_password", "", time.Now().Add(1*time.Hour).Unix())
	if err != nil {
		return err
	}

	err = uuc.MailService.SendPasswordResetEmail(user.Email, resetToken)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCaseImpl) StoreToken(token string) error {
	claims, err := uuc.TokenService.ValidateToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token payload")
	}

	err = uuc.userRepository.StoreResetToken(email, token)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCaseImpl) ResetPassword(token string, newPassword string) error {
	claims, err := uuc.TokenService.ValidateToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token payload")
	}

	storedToken, err := uuc.userRepository.GetResetTokenByEmail(email)
	if err != nil {
		return err
	}

	if storedToken != token {
		return errors.New("invalid or mismatched token")
	}

	hashedPassword, err := uuc.PasswordService.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = uuc.userRepository.UpdatePasswordByEmail(email, hashedPassword)
	if err != nil {
		return err
	}

	err = uuc.userRepository.InvalidateResetToken(email)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCaseImpl) PromoteUser(email string) error {
	return uuc.userRepository.PromoteUser(email)
}

func (uuc *UserUseCaseImpl) DemoteUser(email string) error {
	return uuc.userRepository.DemoteUser(email)
}

func (uuc *UserUseCaseImpl) DeleteUser(email string) error {
	//Delete the user
	user_err := uuc.userRepository.DeleteUser(email)
	if user_err != nil {
		return errors.New("user not found")
	}

	//Delete the refresher associated with the user
	refresher_err := uuc.TokenRepo.DeleteAllRefreshers(email)
	if refresher_err != nil {
		return errors.New("refresher not found")
	}

	return nil
}

// func (uuc *UserUseCaseImpl) DeleteRefresher(email, refresher string) error {
// 	return uuc.TokenRepo.DeleteRefresher(email, refresher)
// }

func (uuc *UserUseCaseImpl) Logout(email, refresher string) error{
	return uuc.TokenRepo.DeleteRefresher(email, refresher)
}

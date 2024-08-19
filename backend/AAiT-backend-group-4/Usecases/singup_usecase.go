package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type singupUsecase struct {
	userRepository  domain.UserRepository
	tokenService    domain.TokenInfrastructure
	passwordService domain.PasswordInfrastructure
	otpUsecase      domain.OTPUsecase
	contextTimeout  time.Duration
}

// NewSingupUsecase creates a new instance of the SignupUsecase implementation.
// It takes in the following parameters:
// - userRepository: The UserRepository implementation for user data access.
// - otpUsecase: The OTPUsecase implementation for OTP-related operations.
// - timeout: The duration for the context timeout.
// - passwordService: The PasswordInfrastructure implementation for password-related operations.
// - tokenService: The TokenInfrastructure implementation for token-related operations.
// It returns a domain.SignupUsecase interface.
func NewSingupUsecase(userRepository domain.UserRepository, otpUsecase domain.OTPUsecase,
	timeout time.Duration, passwordService domain.PasswordInfrastructure, tokenService domain.TokenInfrastructure) domain.SignupUsecase {
	return &singupUsecase{
		userRepository:  userRepository,
		otpUsecase:      otpUsecase,
		tokenService:    tokenService,
		passwordService: passwordService,
		contextTimeout:  timeout,
	}
}

// Signup is a method that handles the signup process for a user.
// It takes a context and a SignupRequest object as input and returns an OTPVerificationResponse and an error.
// The function first checks if the email already exists in the system by calling the GetByEmail method.
// If the email already exists, it returns an error indicating that the email is already taken.
// Next, it checks if the username is already taken by calling the GetByUsername method.
// If the username is already taken, it returns an error indicating that the username is already taken.
// The function then hashes the user's password using the passwordService's HashPassword method.
// If there is an error while hashing the password, it returns an error indicating that it is unable to hash the password.
// It creates a new User object with the provided user details and sets the default values for other fields.
// The function then calls the CreateUser method of the userRepository to create the user in the system.
// If there is an error while creating the user, it returns the error.
// It creates a new UserOTPRequest object with the user's ID and email.
// The function then calls the GenerateOTP method of the otpUsecase to generate an OTP for the user.
// If there is an error while generating the OTP, it returns the error.
// Finally, it returns the generated OTPVerificationResponse and nil error.
func (su *singupUsecase) Signup(c context.Context, user *domain.SignupRequest) (resp domain.OTPVerificationResponse, err error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	_, err = su.GetByEmail(c, user.Email)
	if err == nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("email already exists")
	}

	_, err = su.GetByUsername(c, user.User_Name)
	if err == nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("username alread taken")
	}

	hashedPassword, err := su.passwordService.HashPassword(user.Password)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to hash password")
	}

	newUser := domain.User{
		ID:            primitive.NewObjectID(),
		First_Name:    user.First_Name,
		Last_Name:     user.Last_Name,
		Username:      user.User_Name,
		Email:         user.Email,
		Password:      hashedPassword,
		User_Role:     "USER",
		Access_Token:  "",
		Refresh_Token: "",
		Verified:      false,
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
	}
	err = su.userRepository.CreateUser(ctx, &newUser)
	if err != nil {
		return domain.OTPVerificationResponse{}, err
	}

	newUserOtpRequest := domain.UserOTPRequest{
		UserID: newUser.ID.Hex(),
		Email:  newUser.Email,
	}

	newOtpVericationReponse, err := su.otpUsecase.GenerateOTP(c, &newUserOtpRequest)
	if err != nil {
		return domain.OTPVerificationResponse{}, err
	}

	return newOtpVericationReponse, nil
}

// GetByEmail retrieves a user by their email address.
// It takes a context.Context and the email string as parameters.
// It returns the user domain.User and an error if any.
func (su *singupUsecase) GetByEmail(c context.Context, email string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

// GetByUsername retrieves a user by their username.
// It takes a context.Context and the username as parameters.
// It returns the user domain object and an error if any.
func (su *singupUsecase) GetByUsername(c context.Context, userName string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByUsername(ctx, userName)
}

// CreateAllTokens generates access and refresh tokens for the given user using the provided access and refresh secrets.
// It sets the expiry time for both access and refresh tokens.
// The generated access and refresh tokens are returned along with any error encountered during the token creation process.
func (su *singupUsecase) CreateAllTokens(user *domain.User, accessSecret string, refreshSecret string,
	accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error) {
	accessToken, refreshToken, err = su.tokenService.CreateAllTokens(
		user, accessSecret, refreshSecret, accessExpiry, refreshExpiry)

	if err != nil {
		return "", "", fmt.Errorf("unable to create tokens")
	}

	return accessToken, refreshToken, nil
}

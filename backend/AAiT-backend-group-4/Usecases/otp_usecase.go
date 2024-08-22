package usecases

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type otpUsecase struct {
	otpRepository    domain.OTPRepository
	otpServices      domain.OtpInfrastructure
	passwordServices domain.PasswordInfrastructure
	userRepository   domain.UserRepository
	contextTimeout   time.Duration
	Env              *bootstrap.Env
}

// NewOtpUsecase creates a new instance of the OTP use case.
// It takes an OTP repository, a timeout duration, OTP services, password services,
// an environment configuration, and a user repository as parameters.
// It returns an instance of the domain.OTPUsecase interface.
func NewOtpUsecase(otpRepository domain.OTPRepository, timeout time.Duration,
	otpServices domain.OtpInfrastructure, passwordServices domain.PasswordInfrastructure,
	env bootstrap.Env, userRepository domain.UserRepository) domain.OTPUsecase {
	return &otpUsecase{
		otpRepository:    otpRepository,
		otpServices:      otpServices,
		passwordServices: passwordServices,
		contextTimeout:   timeout,
		userRepository:   userRepository,
		Env:              &env,
	}
}

// GenerateOTP generates a one-time password (OTP) for the given user and performs the necessary operations to store and send the OTP.
// It takes a context.Context object and a *domain.UserOTPRequest object as parameters.
// The function returns a domain.OTPVerificationResponse object and an error.
// The OTP is created using the otpServices.CreateOTP method and then hashed using the passwordServices.HashPassword method.
// The hashed OTP, along with other user information, is stored in the database using the otpRepository.CreateOTP method.
// An email containing the OTP is sent to the user's email address using the otpServices.SendEmail method.
// The function returns a domain.OTPVerificationResponse object with the status "Not Verified" and a message indicating that the OTP has been sent to the user's email for verification.
// If any error occurs during the process, an error is returned with an appropriate error message.
func (ou *otpUsecase) GenerateOTP(c context.Context, user *domain.UserOTPRequest) (otp domain.OTPVerificationResponse, err error) {
	otpCode, err := ou.otpServices.CreateOTP(user)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to create otp code")
	}

	hashedOTPCode, err := ou.passwordServices.HashPassword(otpCode)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to hash otp code")
	}

	NewOTP := domain.UserOTPVerification{
		ID:         primitive.NewObjectID(),
		User_ID:    user.UserID,
		Email:      user.Email,
		OTP:        hashedOTPCode,
		Created_At: time.Now(),
		Expires_At: time.Now().Add(15 * time.Minute),
	}

	err = ou.otpRepository.CreateOTP(c, &NewOTP)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to create and save otp")
	}

	err = ou.otpServices.SendEmail(user.Email, "Email Verification", ou.Env.EmailApiKey, otpCode)
	if err != nil {
		return domain.OTPVerificationResponse{}, err
	}

	NewOtpResponse := domain.OTPVerificationResponse{
		Status:  "Not Verified",
		Message: "OTP Sent to your email please verify your email",
	}

	return NewOtpResponse, nil

}

// VerifyOTP verifies the OTP code provided by the user and updates the user's account status.
// It takes the following parameters:
// - c: The context.Context object for the request.
// - user: The UserOTPRequest object containing user information.
// - otp: The OTP code provided by the user.
// It returns an OTPVerificationResponse object and an error.
// The function first retrieves the OTP code associated with the user's email from the otpRepository.
// If the OTP code is not found, it returns an error indicating that a new OTP verification request should be sent.
// It then compares the provided OTP code with the retrieved OTP code using the passwordServices.ComparePasswords method.
// If the OTP codes do not match, it returns an error indicating that the OTP code is incorrect.
// Next, it deletes the OTP code from the otpRepository using the otpRepository.DeleteOTPByEmail method.
// If the deletion fails, it returns an error indicating that the OTP code could not be removed.
// It then updates the user's account status to verified by creating a UserUpdate object with the Verified field set to true.
// If the user update fails, it returns an error indicating that the user could not be verified.
// Finally, it creates a new OTPVerificationResponse object with the status set to "Verified" and the message set to "Congrats your account is now verified",
// and returns it along with a nil error.
func (ou *otpUsecase) VerifyOTP(c context.Context, email string, otp string) (resp domain.OTPVerificationResponse, err error) {
	userFound, err := ou.userRepository.GetByEmail(c, email)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("user not found")
	}

	user := domain.UserOTPRequest{
		UserID: userFound.ID.Hex(),
		Email:  userFound.Email,
	}

	otpFound, err := ou.otpRepository.GetOTPByEmail(c, user.Email)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("otp not found please send a new otp verification request")
	}

	if otpFound.Expires_At.Before(time.Now()) {
		err = ou.otpRepository.DeleteOTPByEmail(c, otpFound.Email)
		if err != nil {
			return domain.OTPVerificationResponse{}, fmt.Errorf("unable to remove otp")
		}

		return domain.OTPVerificationResponse{}, fmt.Errorf("your otp code has expired request new code")
	}

	err = ou.passwordServices.ComparePasswords(otp, otpFound.OTP)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("incorrect otp code")
	}

	err = ou.otpRepository.DeleteOTPByEmail(c, otpFound.Email)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to remove otp")
	}

	validated := true
	updateUser := domain.UserUpdate{
		Verified: &validated,
	}

	_, err = ou.userRepository.UpdateUser(c, user.UserID, updateUser)
	if err != nil {
		return domain.OTPVerificationResponse{}, err
	}

	newOtpVerificationResponse := domain.OTPVerificationResponse{
		Status:  "Verified",
		Message: "Congrats your account is now verified",
	}

	return newOtpVerificationResponse, nil
}

func (ou *otpUsecase) ResendOTP(c context.Context, email string) (resp domain.OTPVerificationResponse, err error) {
	userFound, err := ou.userRepository.GetByEmail(c, email)
	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("user not found")
	}

	user := domain.UserOTPRequest{
		UserID: userFound.ID.Hex(),
		Email:  userFound.Email,
	}

	response, err := ou.GenerateOTP(c, &user)

	if err != nil {
		return domain.OTPVerificationResponse{}, fmt.Errorf("unable to resend otp")
	}

	return response, nil
}

func (ou *otpUsecase) SendPasswordResetEmail(c context.Context, email string) (otp string, err error) {
	return "", nil
}

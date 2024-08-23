package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	utils "blogs/Utils"
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupUseCase struct {
	SignupRepository         domain.SignupRepository
	UnverifiedUserRepository domain.UnverifiedUserRepository
	contextTimeout           time.Duration
	passwordService          domain.PasswordService
}

func NewSignupUseCase(SignupRepository domain.SignupRepository, uvu domain.UnverifiedUserRepository, timeout time.Duration, passwordService domain.PasswordService) domain.SignupUseCase {
	return &SignupUseCase{
		SignupRepository:         SignupRepository,
		UnverifiedUserRepository: uvu,
		contextTimeout:           timeout,
		passwordService:          passwordService}
}

func (u *SignupUseCase) Create(c context.Context, user domain.User) interface{} {
	// check empty fields
	if user.Email == "" || user.Username == "" || user.Password == "" {
		return &domain.ErrorResponse{Message: "All fields are required", Status: 400}
	}

	// CHECK EMAIL VALIDITY
	if u.passwordService.ValidateEmail(user.Email) != nil {
		return &domain.ErrorResponse{Message: "Invalid email format", Status: 400}
	}

	// CHECK PASSWORD VALIDITY
	if err := u.passwordService.ValidatePassword(user.Password); err != nil {
		return &domain.ErrorResponse{Message: err.Error(), Status: 400}
	}

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	idofNumber := primitive.NewObjectID()
	user.ID = idofNumber

	// check if user already exists
	existingUser, err := u.SignupRepository.FindUserByEmail(ctx, user.Email)

	if err == nil {
		return &domain.ErrorResponse{Message: "User already exists", Status: 400}
	} else if err == nil && !existingUser.Verified {
		return u.HandleUnverifiedUser(c, existingUser)
	}

	// hash the password
	hashedPassword, err := u.passwordService.HashPassword(user.Password)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}

	user.Password = hashedPassword

	// 15 minute for expiration
	user.ExpiresAt = time.Now().Add(time.Minute * 2)

	// send OTP
	otp, err := infrastructure.GenerateOTP()
	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OTP", Status: 500}
	}

	// save OTP to db
	user.PostsID = utils.MakePrimitiveList(0)
	var newuser domain.UnverifiedUser
	newuser.Email = user.Email
	newuser.OTP = otp
	exp:=time.Now().Add(time.Minute * 10)
	unverifiedClaim := domain.UnverifiedUserClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
		},
	}
	newuser.UserToken, err = infrastructure.CreateToken(unverifiedClaim, "unverified")
	if err != nil {
		return &domain.ErrorResponse{Message: "Error creating token", Status: 500}
	}

	err = u.UnverifiedUserRepository.StoreUnverifiedUser(ctx, newuser)
	if err != nil {
		return &domain.ErrorResponse{Message: "Error creating user", Status: 500}
	}

	err = u.SignupRepository.SetOTP(ctx, user.Email, otp)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error saving OTP", Status: 500}
	}

	err = infrastructure.SendOTPEmail(user.Email, otp)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error sending OTP", Status: 500}
	}

	return &domain.SuccessResponse{Message: "Registerd Sucessfully Verify your account", Data: "", Status: 201}
}

func (u *SignupUseCase) VerifyOTP(c context.Context, otp domain.OtpToken) interface{} {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check if OTP is correct

	user, err := u.SignupRepository.FindUserByEmail(ctx, otp.Email)
	if err != nil {
		return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}

	if user.OTP != otp.OTP {
		return &domain.ErrorResponse{Message: "Invalid OTP", Status: 400}
	}

	// check if OTP is expired
	if time.Now().After(user.ExpiresAt) {
		return &domain.ErrorResponse{Message: "OTP expired", Status: 400}
	}

	// update user
	user.Verified = true
	user.OTP = ""

	verifiedUser, err := u.SignupRepository.VerifyUser(ctx, user)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error verifying user", Status: 500}
	}

	return &domain.SuccessResponse{Message: "Account verified successfully", Data: verifiedUser, Status: 200}

}

func (u *SignupUseCase) ForgotPassword(c context.Context, email domain.ForgotPasswordRequest) interface{} {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check if user exists
	existing, err := u.SignupRepository.FindUserByEmail(ctx, email.Email)
	if err != nil {
		return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}

	// generate token

	// check if token is already set and the expiration time is not passed
	if existing.ResetPasswordToken != "" && time.Now().Before(existing.ResetPasswordExpires) {
		difftime := existing.ResetPasswordExpires.Sub(time.Now())
		return &domain.ErrorResponse{Message: "Reset token already sent Please wait for " + strconv.FormatFloat(difftime.Minutes(), 'f', -1, 64) + " to resend reset token", Status: 400}
	}

	token, err := infrastructure.GenerateResetToken()

	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating reset token", Status: 500}
	}

	// save token to db
	// expiration time 15 minutes

	expiration := time.Now().Add(time.Minute * 3)

	_, err = u.SignupRepository.SetResetToken(ctx, email, token, expiration)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error saving reset token", Status: 500}
	}

	// send reset email

	err = infrastructure.SendResetEmail(email.Email, token)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error sending reset email", Status: 500}
	}

	return &domain.SuccessResponse{Message: "Reset email sent", Data: "", Status: 200}

}

func (u *SignupUseCase) ResetPassword(c context.Context, password domain.ResetPasswordRequest, token string) interface{} {

	// check the password validity
	if err := u.passwordService.ValidatePassword(password.Password); err != nil {
		return &domain.ErrorResponse{Message: err.Error(), Status: 400}
	}

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check if the ResetToken is Set

	user, err := u.SignupRepository.FindUserByResetToken(ctx, token)

	if err != nil {
		return &domain.ErrorResponse{Message: "Invalid reset token", Status: 400}
	}

	// check if token is expired

	if time.Now().After(user.ResetPasswordExpires) {
		return &domain.ErrorResponse{Message: "Reset token expired", Status: 400}
	}
	hashedPassword, err := u.passwordService.HashPassword(password.Password)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}

	// hash the password

	user.Password = hashedPassword
	user.ResetPasswordToken = ""
	user.ResetPasswordExpires = time.Time{}

	// update user
	_, err = u.SignupRepository.UpdateUser(ctx, user)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error in Reseting the Password", Status: 500}
	}

	return &domain.SuccessResponse{Message: "Password Reset Sucessfully", Status: 200}
}

func (u *SignupUseCase) HandleUnverifiedUser(c context.Context, user domain.User) interface{} {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if user.Username == "" {
		existingUser, err := u.SignupRepository.FindUserByEmail(ctx, user.Email)
		if err != nil {
			return &domain.ErrorResponse{Message: "User not found", Status: 404}
		}
		user = existingUser
	}
	if user.Verified {
		return &domain.ErrorResponse{Message: "User already verified", Status: 400}
	}
	// check if the user send the register button again Not to send the OTP again before the expiration time

	if time.Now().Before(user.ExpiresAt) {
		difftime := user.ExpiresAt.Sub(time.Now())

		return &domain.ErrorResponse{Message: "Otp already Sent Please wait for " + strconv.FormatFloat(difftime.Minutes(), 'f', -1, 64) + " to resend OTP", Status: 400}

	}

	user.ExpiresAt = time.Now().Add(time.Minute * 10)
	_, err := u.SignupRepository.UpdateUser(ctx, user)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error in setting Expiration time", Status: 500}
	}
	// Generate OTP
	otp, err := infrastructure.GenerateOTP()

	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OTP", Status: 500}
	}

	// SaveOTP to the DB
	err = u.SignupRepository.SetOTP(ctx, user.Email, otp)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error saving OTP", Status: 500}
	}

	// Send The email
	err = infrastructure.SendOTPEmail(user.Email, otp)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error sending OTP", Status: 500}
	}

	return &domain.SuccessResponse{Message: "OTP send to your Email Verify Your Account", Data: "", Status: 201}
}

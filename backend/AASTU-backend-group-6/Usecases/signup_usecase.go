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
	_, err := u.SignupRepository.FindUserByEmail(ctx, user.Email)

	if err == nil {
		return &domain.ErrorResponse{Message: "User already exists please login", Status: 400}
	}
	existingUnverifiedUser, err := u.UnverifiedUserRepository.FindUnverifiedUser(ctx, user.Email)

	if existingUnverifiedUser.Email != "" && err == nil {
		return &domain.ErrorResponse{Message: "User already Registerd Verify Your account", Status: 400}

	}
	// hash the password
	hashedPassword, err := u.passwordService.HashPassword(user.Password)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}

	user.Password = hashedPassword


	// Generate OTP
	otp, err := infrastructure.GenerateOTP()
	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OTP", Status: 500}
	}

	// save OTP to db
	user.PostsID = utils.MakePrimitiveList(0)
	var newuser domain.UnverifiedUser
	newuser.Email = user.Email
	newuser.OTP = otp
	newuser.ExpiresAt = time.Now().Add(time.Minute * 10)
	exp := time.Now().Add(time.Hour * 10)
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
	newuser.Created_at = time.Now()
	err = infrastructure.SendOTPEmail(user.Email, otp)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error sending OTP", Status: 500}
	}
	err = u.UnverifiedUserRepository.StoreUnverifiedUser(ctx, newuser)
	if err != nil {
		return &domain.ErrorResponse{Message: "Error creating user", Status: 500}
	}

	return &domain.SuccessResponse{Message: "Registerd Sucessfully Verify your account", Data: "", Status: 201}
}

func (u *SignupUseCase) VerifyOTP(c context.Context, otp domain.OtpToken) interface{} {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check if OTP is correct

	unverifiedUser, err := u.UnverifiedUserRepository.FindUnverifiedUser(ctx, otp.Email)
	if err != nil {
		return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}
	user, err := infrastructure.ExtractFromToken(unverifiedUser.UserToken, "unverified")
	if err != nil {
		u.UnverifiedUserRepository.DeleteUnverifiedUser(ctx, otp.Email)
		return &domain.ErrorResponse{Message: "Register Again", Status: 500}
	}

	if unverifiedUser.OTP != otp.OTP {
		return &domain.ErrorResponse{Message: "Invalid OTP", Status: 400}
	}

	// check if OTP is expired
	if time.Now().After(unverifiedUser.ExpiresAt) {
		return &domain.ErrorResponse{Message: "OTP expired resend OTP", Status: 400}
	}

	// create Verified user
	user.PostsID = utils.MakePrimitiveList(0)
	user.Role = "user"
	user.CommentsID = utils.MakePrimitiveList(0)
	user.LikedPostsID = utils.MakePrimitiveList(0)
	user.DisLikePostsID = utils.MakePrimitiveList(0)
	verifiedUser, err := u.SignupRepository.Create(ctx, user)
	verifiedUser.Password = ""
	
	
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

	if existing.GoogleID != "" {
		return &domain.ErrorResponse{Message: "User is registered with Google", Status: 400}
	}

	// generate token

	// check if token is already set and the expiration time is not passed
	if existing.ResetPasswordToken != "" && time.Now().Before(existing.ResetPasswordExpires) {
		difftime := time.Until(existing.ResetPasswordExpires)
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
func (u *SignupUseCase) HandleUnverifiedUser(c context.Context, user domain.Email) interface{} {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
  
	email := user.Email
  
	// check if the user is already verified
	_ , err := u.SignupRepository.FindUserByEmail(ctx, email)
  
	if err == nil { 
	  return &domain.ErrorResponse{Message: "User Already Verified", Status: 404}
	}
  
	// check if the user is already registered
  
	existingUser, err := u.UnverifiedUserRepository.FindUnverifiedUser(ctx, email)
  
	if err != nil { 
	  return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}
  
	// check if the user send the register button again Not to send the OTP again before the expiration time
  
	if time.Now().Before(existingUser.ExpiresAt) {
	  difftime := existingUser.ExpiresAt.Sub(time.Now())
	  return &domain.ErrorResponse{Message: "Otp already Sent Please wait for " + strconv.FormatFloat(difftime.Minutes(), 'f', -1, 64)[:2] + " to resend OTP", Status: 400}
  
	}
  
	// Generate OTP
	otp, err := infrastructure.GenerateOTP()
  
	if err != nil { 
	  return &domain.ErrorResponse{Message: "Error in Generating OTP", Status: 500}
	}
  
	expiry := time.Now().Add(time.Minute * 10)
  
	_ , err  = u.UnverifiedUserRepository.UpdateOTP(ctx, email , otp , expiry)
	if err != nil { 
	  return &domain.ErrorResponse{Message: "Error in setting OTP ", Status: 500}
	}
  
	// send OTP
	err = infrastructure.SendOTPEmail(email, otp)
  
	if err != nil { 
	  return &domain.ErrorResponse{Message: "Error sending OTP", Status: 500}
	}
  
	return &domain.SuccessResponse{Message: "OTP send to your Email Verify Your Account", Data: "", Status: 201}
  
  }
  

//   background Task that delete the user that stays unverified over 2 week

func (u *SignupUseCase) DeleteOldUnverifiedUsers(c context.Context , days int) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), u.contextTimeout)
	defer cancel()

    cutoffDate := time.Now().AddDate(0, 0, -days)
    return u.UnverifiedUserRepository.DeleteUnverifiedUsersBefore(ctx , cutoffDate)
	
}
	


  
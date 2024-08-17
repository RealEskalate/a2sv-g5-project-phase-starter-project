package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupUseCase struct {
	SignupRepository domain.SignupRepository
	contextTimeout time.Duration

}

func NewSignupUseCase(SignupRepository domain.SignupRepository , timeout time.Duration) domain.SignupUseCase {
	return &SignupUseCase{SignupRepository: SignupRepository,
							contextTimeout: timeout}	
}

func (u *SignupUseCase) Create(c context.Context , user domain.User) interface{} {
	// check empty fields
	if user.Email == "" || user.Username == "" || user.Password == "" {
		return &domain.ErrorResponse{Message: "All fields are required" , Status: 400}
	}

	ctx , cancel :=context.WithTimeout(c , u.contextTimeout)
	defer cancel()
	idofNumber := primitive.NewObjectID()
	user.ID = idofNumber

	// check if user already exists
	_, err := u.SignupRepository.FindUserByEmail(ctx , user.Email)
	if err == nil {
		return &domain.ErrorResponse{Message: "User already exists", Status: 400}
	}
	// hash the password
	hashedPassword, err := infrastructure.HashPassword(user.Password)

	if err != nil {
		return &domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}

	user.Password = hashedPassword

	// 15 minute for expiration 
	user.ExpiresAt = time.Now().Add(time.Minute  * 10)

	// send OTP
	otp , err := infrastructure.GenerateOTP()

	if err != nil {
		return &domain.ErrorResponse{Message: "Error generating OTP", Status: 500}
	}
	

	// save OTP to db
	_ , err = u.SignupRepository.Create(ctx  , user)
	if err != nil {
		return &domain.ErrorResponse{Message: "Error creating user", Status: 500}
	}

	err = u.SignupRepository.SetOTP(ctx , user.Email , otp)

	fmt.Println(err)
	if err != nil {
		return &domain.ErrorResponse{Message: "Error saving OTP", Status: 500}
	}

	err = infrastructure.SendOTPEmail(user.Email, otp)

	if err != nil { 
		return &domain.ErrorResponse{Message: "Error sending OTP", Status: 500}
	}



	return &domain.SuccessResponse{Message: "Registerd Sucessfully Verify your account" , Status: 201}
}


func (u *SignupUseCase) VerifyOTP(c context.Context , otp domain.OtpToken) interface{} { 

	ctx , cancel := context.WithTimeout(c , u.contextTimeout)
	defer cancel()

	// check if OTP is correct
	user, err := u.SignupRepository.FindUserByEmail(ctx , otp.Email)
	if err != nil {
		return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}


	if user.OTP != otp.OTP {
		return &domain.ErrorResponse{Message: "Invalid OTP", Status: 400}
	}

	// check if OTP is expired
	if time.Now().After(user.ExpiresAt) {
		return   &domain.ErrorResponse{Message: "OTP expired", Status: 400}
	}

	// update user
	user.Verified = true

	verifiedUser , err := u.SignupRepository.VerifyUser(ctx , user)

	if err != nil { 
		return domain.ErrorResponse{Message: "Error verifying user", Status: 500}
	}

	return domain.SuccessResponse{Message: "OTP verified successfully", Data: verifiedUser, Status: 200}

}
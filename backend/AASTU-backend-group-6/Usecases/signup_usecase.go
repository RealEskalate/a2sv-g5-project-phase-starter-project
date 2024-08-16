package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"time"
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
		return domain.ErrorResponse{Message: "All fields are required" , Status: 400}
	}

	ctx , cancel :=context.WithTimeout(c , u.contextTimeout)
	defer cancel()


	// check if user already exists
	_, err := u.SignupRepository.FindUserByEmail(ctx , user.Email)
	if err == nil {
		return domain.ErrorResponse{Message: "User already exists", Status: 400}
	}
	// hash the password
	hashedPassword, err := infrastructure.HashPassword(user.Password)

	if err != nil {
		return domain.ErrorResponse{Message: "Error hashing password", Status: 500}
	}

	user.Password = hashedPassword

	// create user
	createdUser, err := u.SignupRepository.Create(user)

	if err != nil {
		return domain.ErrorResponse{Message: "Error creating user", Status: 500}
	}

	return domain.SuccessResponse{Message: "User created successfully", Data: createdUser, Status: 201}
}

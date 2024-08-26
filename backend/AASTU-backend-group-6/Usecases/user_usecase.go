package usecases

import (
	domain "blogs/Domain"
	"context"
	"time"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
	passwordService domain.PasswordService
	
}

func NewUserUseCase(UserRepository domain.UserRepository, timeout time.Duration , passwordService domain.PasswordService) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: UserRepository,
		contextTimeout: timeout,
		passwordService: passwordService,
	}
}


func (uc *UserUseCase) UpdateUser(c context.Context, req domain.UserUpdateRequest) interface{} {
	// Fetch the current user data from the database
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	getUser, err := uc.UserRepository.FindUserByID(c, req.ID)
	if err != nil {
		return &domain.ErrorResponse{Message: "User not found", Status: 404}
	}

	// Check if the user is updating their own account
	if getUser.ID.Hex() != req.ID {
		return &domain.ErrorResponse{Message: "Unauthorized to update this user", Status: 403}
	}

	// Check if the new username already exists (excluding the current user)
	if req.Username != "" {
		existingUser, err := uc.UserRepository.FindUserByUsername(ctx, req.Username)
		if err == nil && existingUser.ID.Hex() != getUser.ID.Hex() {
			return &domain.ErrorResponse{Message: "Username is already taken", Status: 409}
		}
	}

	
	// Update only the fields that are set in the request
	if req.Full_Name != "" {
		getUser.Full_Name = req.Full_Name
	}
	if req.Username != "" {
		getUser.Username = req.Username
	}
	if req.Password != ""{
		if getUser.GoogleID == "" {

			// Validate the password
			if err := uc.passwordService.ValidatePassword(req.Password); err != nil {
				return &domain.ErrorResponse{Message: err.Error(), Status: 400}
			}
			
			hashedpassword , err := uc.passwordService.HashPassword(req.Password)
			if err != nil{
				return &domain.ErrorResponse{Message: "Error in hashing the passeword" , Status:500}
			}
			getUser.Password = hashedpassword

		}else{
			return &domain.ErrorResponse{Message: "Cannot change password for google users" , Status:400}
		}
				
	}
	if req.Profile_image_url != "" {
		getUser.Profile_image_url = req.Profile_image_url
	}
	if req.Contact != "" {
		getUser.Contact = req.Contact
	}
	if req.Bio != "" {
		getUser.Bio = req.Bio
	}

	// Update the user in the repository
	updatedUser, err := uc.UserRepository.UpdateUser(ctx, getUser)
	updatedUser.Password = ""
	if err != nil {
		return &domain.ErrorResponse{Message: "Error updating user", Status: 500}
	}

	return &domain.SuccessResponse{Message: "User updated successfully", Status: 200, Data: updatedUser}
}


func (uc *UserUseCase) PromoteandDemoteUser(c context.Context , userId string , promotion domain.UserPromotionRequest) interface{} {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	

	// Fetch the user from the database
	_, err := uc.UserRepository.FindUserByID(ctx, userId)

	if err != nil {
		return &domain.ErrorResponse{Message: err.Error(), Status: 404}
	}

	if promotion.Action == "demote"{
		err := uc.UserRepository.PromoteandDemoteUser(ctx, userId, "user")
		if err != nil {
			return &domain.ErrorResponse{Message: err.Error(), Status: 500}
		}
		return &domain.SuccessResponse{Message: "User demoted successfully", Status: 200}
	}else if promotion.Action != "promote"{
		return &domain.ErrorResponse{Message: "Invalid action", Status: 400}
	}
	err = uc.UserRepository.PromoteandDemoteUser(ctx, userId, "admin")
	if err != nil {
		return &domain.ErrorResponse{Message: err.Error(), Status: 500}
	}
	return &domain.SuccessResponse{Message: "User promoted successfully", Status: 200}
	
}

func (uc *UserUseCase) FindUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.UserRepository.FindUserByID(ctx, id)
}
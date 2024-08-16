package usecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"
)

func (u *UserUsecase) UpdateProfile(usernameoremail string, user *domain.User) error {
	// Get existing user details
	existingUser, err := u.UserRepo.GetUserByUsernameorEmail(usernameoremail)
	if err != nil {
		return err
	}

	// Check if the first name is present
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}

	// Check if the last name is present
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	// Check if the bio is present
	if user.Bio != "" {
		existingUser.Bio = user.Bio
	}

	// Check if the avatar is present
	if user.Avatar != "" {
		existingUser.Avatar = user.Avatar
	}

	// Check if the address is present
	if user.Address != "" {
		existingUser.Address = user.Address
	}

	// Check if the email is present and is unique and valid
	if user.Email != "" {
		// Validate the email
		err = config.IsValidEmail(user.Email)
		if err != nil {
			return err
		}

		// Check if the email is unique
		err = u.UserRepo.CheckUsernameAndEmail(existingUser.Username, user.Email)
		if err != nil {
			return err
		}

		existingUser.Email = user.Email
	}

	// Check if the password is present and is strong
	if user.Password != "" {
		// Validate the password
		err = config.IsStrongPassword(user.Password)
		if err != nil {
			return err
		}

		// Hash the new password
		hashedPassword, err := config.HashPassword(user.Password)
		if err != nil {
			return err
		}

		existingUser.Password = hashedPassword
	}

	// Check if the role is present
	if user.Role != "" {
		return errors.New("role cannot be updated")
	}

	// Check if the joined date is present
	if !user.JoinedDate.IsZero() {
		return errors.New("joined date cannot be updated")
	}

	// Update the user profile in the repository
	err = u.UserRepo.UpdateProfile(usernameoremail, existingUser)
	if err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"errors"
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"time"
)


func (u *UserUsecase) Login(user *domain.User, deviceID string) (domain.LogInResponse, error) {
    if u.UserRepo == nil {
        return domain.LogInResponse{}, errors.New("UserRepo is not initialized")
    }
    if u.PasswordSvc == nil {
        return domain.LogInResponse{}, errors.New("PasswordSvc is not initialized")
    }
    if u.TokenGen == nil {
        return domain.LogInResponse{}, errors.New("TokenGen is not initialized")
    }

    existingUser, err := u.UserRepo.Login(user)
    if err != nil {
        return domain.LogInResponse{}, errors.New("invalid credentials")
    }

    if !u.PasswordSvc.CheckPasswordHash(user.Password, existingUser.Password) {
        return domain.LogInResponse{}, errors.New("invalid credentials")
    }

    refreshToken, err := u.TokenGen.GenerateRefreshToken(*existingUser)
    if err != nil {
        return domain.LogInResponse{}, err
    }

    newRefreshToken := domain.RefreshToken{
        Token:     refreshToken,
        DeviceID:  deviceID,
        CreatedAt: time.Now(),
    }	
	
	for i, rt := range existingUser.RefreshTokens {
		if rt.DeviceID == deviceID  {
			existingUser.RefreshTokens = append(existingUser.RefreshTokens[:i], existingUser.RefreshTokens[i+1:]...)
			break
		}
	}
	
	existingUser.RefreshTokens = append(existingUser.RefreshTokens, newRefreshToken)
	
	

    err = u.UserRepo.UpdateUser(existingUser)
    if err != nil {
        return domain.LogInResponse{}, err
    }

    accessToken, err := u.TokenGen.GenerateToken(*existingUser)
    if err != nil {
        return domain.LogInResponse{}, err
    }

    return domain.LogInResponse{
        AccessToken:  accessToken,
        RefreshToken: newRefreshToken.Token,
    }, nil
}


func (u *UserUsecase) Logout(userID, deviceID, token string) error {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	for i, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("invalid token")
}

func (u *UserUsecase) LogoutAllDevices(userID string) error {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	user.RefreshTokens = []domain.RefreshToken{}
	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) LogoutDevice(userID, deviceID string) error {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	for i, rt := range user.RefreshTokens {
		if rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("device not found")
}

func (u *UserUsecase) GetDevices(userID string) ([]string, error) {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	var devices []string
	for _, rt := range user.RefreshTokens {
		devices = append(devices, rt.DeviceID)
	}
	return devices, nil
}

func (u *UserUsecase) RefreshToken(userID, deviceID, token string) (domain.LogInResponse, error) {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.LogInResponse{}, errors.New("user not found")
	}

	for i, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {

			// _, tokenErr := infrastracture.RefreshToken(token)
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)

			
			// if tokenErr != nil {
			// 	return domain.LogInResponse{}, errors.New("invalid token")
			// }

			refreshToken, err := u.TokenGen.GenerateRefreshToken(user)
			if err != nil {
				return domain.LogInResponse{}, err
			}

			newRefreshToken := domain.RefreshToken{
				Token:     refreshToken,
				DeviceID:  deviceID,
				CreatedAt: time.Now(),
			}

			user.RefreshTokens = append(user.RefreshTokens, newRefreshToken)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return domain.LogInResponse{}, err
			}

			accessToken, err := u.TokenGen.GenerateToken(user)
			if err != nil {
				return domain.LogInResponse{}, err
			}

			return domain.LogInResponse{
				AccessToken:  accessToken,
				RefreshToken: newRefreshToken.Token,

			}, err
			

		}
	}

	return domain.LogInResponse{}, errors.New("invalid token")
}

func (u *UserUsecase) Register(user domain.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("all fields are required")
	}

	if !infrastracture.IsValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if !infrastracture.IsValidPassword(user.Password) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one digit, one special character and minimum length of 8 characters")
	}

	_, err := u.UserRepo.GetUserByUsernameOrEmail(user.Username, user.Email)
	if err == nil {
		return errors.New("username or email already exists")
	}

	user.Role = "user"

	// Hash password
	hashedPassword, err := u.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return errors.New("could not hash password")
	}

	token, err := infrastracture.GenerateActivationToken()
	if err != nil {
		return errors.New("could not generate activation token")
	}

	user.Password = hashedPassword
	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	// Create user account in the database
	err = u.UserRepo.Register(user)
	if err != nil {
		return err
	}

	// Send activation email or link to the user
	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetUserByUsernameOrEmail(username, email string) (domain.User, error) {
	return u.UserRepo.GetUserByUsernameOrEmail(username, email)
}

func (u *UserUsecase) AccountActivation(token string, email string) error {
	return u.UserRepo.AccountActivation(token, email)
}

func (u *UserUsecase) SendPasswordResetLink(email string) error {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	// Generate a reset token (you can use JWT or a random string)
	resetToken, err := infrastracture.GenerateActivationToken()
	if err != nil {
		return errors.New("could not generate reset token")
	}
	user.PasswordResetToken = resetToken
	user.TokenCreatedAt = time.Now()

	// Save the reset token in the database
	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return errors.New("failed to save reset token")
	}

	// Send the email with the reset link (implement your email logic)
	err = infrastracture.SendResetLink(user.Email, resetToken)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) ResetPassword(token, newPassword string) error {
	user, err := u.UserRepo.GetUserByResetToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if time.Since(user.TokenCreatedAt) > 24*time.Hour {
		return errors.New("token has expired")
	}

	hashedPassword, err := u.PasswordSvc.HashPassword(newPassword)
	if err != nil {
		return errors.New("could not hash password")
	}

	user.Password = hashedPassword
	user.PasswordResetToken = ""
	user.TokenCreatedAt = time.Time{}

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return errors.New("failed to reset password")
	}

	return nil
}



func (u *UserUsecase) ActivateAccountMe(userID string) error {

	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	token, err := infrastracture.GenerateActivationToken()
	if err != nil {
		return errors.New("could not generate activation token")
	}

	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return errors.New("failed to save activation token")
	}

	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
	
}
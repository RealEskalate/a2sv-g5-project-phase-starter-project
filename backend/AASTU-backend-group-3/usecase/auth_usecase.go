package usecase

import (
	"errors"
	"fmt"
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"time"
)

func (u *UserUsecase) Login(user *domain.User, deviceID string) (domain.LogInResponse, error) {
	existingUser, err := u.UserRepo.Login(user)
	if err != nil {
		return domain.LogInResponse{}, errors.New("invalid credentials")
	}
	if !infrastracture.CheckPasswordHash(user.Password, existingUser.Password) {
		return domain.LogInResponse{}, errors.New("invalid credentials")
	}

	refreshToken, err := infrastracture.GenerateRefreshToken(existingUser)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	newRefreshToken := domain.RefreshToken{
		Token:     refreshToken,
		DeviceID:  deviceID,
		CreatedAt: time.Now(),
	}

	existingUser.RefreshTokens = append(existingUser.RefreshTokens, newRefreshToken)

	err = u.UserRepo.UpdateUser(existingUser)
	if err != nil {
		return domain.LogInResponse{}, err
	}

	accessToken, err := infrastracture.GenerateToken(*existingUser)
	if err != nil {
		return domain.LogInResponse{}, err
	}
	LogInResponse := domain.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.Token,
	}

	return LogInResponse, nil
}

func (au *UserUsecase) Logout(userID, deviceID, token string) error {
	user, err := au.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	for i, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = au.UserRepo.UpdateUser(&user)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("invalid token")
}

func (au *UserUsecase) LogoutAllDevices(userID string) error {
	user, err := au.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	user.RefreshTokens = []domain.RefreshToken{}
	err = au.UserRepo.UpdateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

// logout a specific device with deviceId
func (au *UserUsecase) LogoutDevice(userID, deviceID string) error {
	user, err := au.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	for i, rt := range user.RefreshTokens {
		if rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = au.UserRepo.UpdateUser(&user)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("device not found")
}

func (au *UserUsecase) GetDevices(userID string) ([]string, error) {
	user, err := au.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	var devices []string
	for _, rt := range user.RefreshTokens {
		devices = append(devices, rt.DeviceID)
	}
	return devices, nil
}

func (au *UserUsecase) RefreshToken(userID, deviceID, token string) (domain.LogInResponse, error) {
	user, err := au.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.LogInResponse{}, errors.New("user not found")
	}

	for _, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {
			t, tokenErr := infrastracture.RefreshToken(token)
			for i, v := range user.RefreshTokens {
				if v.Token == token {
					user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
					break
				}
			}
			
			if tokenErr != nil {

				return domain.LogInResponse{}, errors.New("invalid token")
			}
			refreshToken, err := infrastracture.GenerateRefreshToken(&user)
			if err != nil {
				return domain.LogInResponse{}, err
			}

			newRefreshToken := domain.RefreshToken{
				Token:     refreshToken,
				DeviceID:  deviceID,
				CreatedAt: time.Now(),
			}

			user.RefreshTokens = append(user.RefreshTokens, newRefreshToken)
			err = au.UserRepo.UpdateUser(&user)
			if err != nil {
				return domain.LogInResponse{}, err
			}

			return domain.LogInResponse{
				AccessToken:  t,
				RefreshToken: newRefreshToken.Token,
			}, err
		}
	}

	return domain.LogInResponse{}, errors.New("invalid token")
}
func (au *UserUsecase) Register(user domain.User) error {

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("all fields are required")
	}

	if !infrastracture.IsValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if !infrastracture.IsValidPassword(user.Password) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one digit, one special character and minimum length of 8 characters")
	}

	_, err := au.UserRepo.GetUserByUsernameOrEmail(user.Username, user.Email)

	if err == nil {
		return errors.New("username or email already exists")
	}

	user.Role = "user"

	// Hash password
	hashedPassword, err := infrastracture.HashPassword(user.Password)
	token, err2 := infrastracture.GenerateActivationToken()

	if err != nil {
		return errors.New("could not hash password")
	}
	if err2 != nil {
		return errors.New("could not generate activation token")
	}
	user.Password = hashedPassword
	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	// Create user account in the database
	err = au.UserRepo.Register(user)
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

func (au *UserUsecase) GetUserByUsernameOrEmail(username, email string) (domain.User, error) {
	return au.UserRepo.GetUserByUsernameOrEmail(username, email)
}

func (au *UserUsecase) AccountActivation(token string, email string) error {
	return au.UserRepo.AccountActivation(token, email)
}




// reset password

func (uc *UserUsecase) SendPasswordResetLink(email string) error {
	user, err := uc.UserRepo.GetUserByEmail(email)
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
	err = uc.UserRepo.UpdateUser(&user)
	if err != nil {
		return errors.New("failed to save reset token")
	}

	// Send the email with the reset link (implement your email logic)
	
	infrastracture.SendResetLink(user.Email, resetToken)

	return nil
}

func (uc *UserUsecase) ResetPassword(token, newPassword string) error {
	user, err := uc.UserRepo.GetUserByResetToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if time.Since(user.TokenCreatedAt) > 24*time.Hour {
		return errors.New("token has expired")
	}

	// Update the user's password
	fmt.Println("newPassword////////////", newPassword)
	hashedPassword, err  := infrastracture.HashPassword(newPassword)
	user.Password = hashedPassword
	if err != nil {
		return errors.New("could not hash password")
	}

	user.PasswordResetToken = ""
	user.TokenCreatedAt = time.Time{}

	err = uc.UserRepo.UpdateUser(&user)
	if err != nil {
		return errors.New("failed to reset password")
	}

	return nil
}
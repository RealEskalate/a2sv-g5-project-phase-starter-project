package usecase

import (
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"log"
	"time"
)

func (u *UserUsecase) Login(user *domain.User, deviceID string) (domain.LogInResponse, *domain.CustomError) {
	if u.UserRepo == nil || u.PasswordSvc == nil || u.TokenGen == nil {
		log.Fatal("Necessary services are nil")
		return domain.LogInResponse{}, domain.ErrInternalServer
	}

	existingUser, err := u.UserRepo.Login(user)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrInvalidCredentials
	}

	if !u.PasswordSvc.CheckPasswordHash(user.Password, existingUser.Password) {
		return domain.LogInResponse{}, domain.ErrInvalidCredentials
	}

	refreshToken, err := u.TokenGen.GenerateRefreshToken(*existingUser)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrInternalServer
	}

	newRefreshToken := domain.RefreshToken{
		Token:     refreshToken,
		DeviceID:  deviceID,
		CreatedAt: time.Now(),
	}

	for i, rt := range existingUser.RefreshTokens {
		if rt.DeviceID == deviceID {
			existingUser.RefreshTokens = append(existingUser.RefreshTokens[:i], existingUser.RefreshTokens[i+1:]...)
			break
		}
	}

	existingUser.RefreshTokens = append(existingUser.RefreshTokens, newRefreshToken)

	err = u.UserRepo.UpdateUser(existingUser)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrFailedToUpdateUser
	}

	accessToken, err := u.TokenGen.GenerateToken(*existingUser)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrInternalServer
	}

	return domain.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.Token,
	}, &domain.CustomError{}
}

func (u *UserUsecase) Logout(userID, deviceID, token string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.ErrNotFound
	}

	for i, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return domain.ErrFailedToUpdateUser
			}
			return &domain.CustomError{}
		}
	}

	return domain.ErrInvalidToken
}

func (u *UserUsecase) LogoutAllDevices(userID string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.ErrNotFound
	}

	user.RefreshTokens = []domain.RefreshToken{}
	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return domain.ErrFailedToUpdateUser
	}
	return &domain.CustomError{}
}

func (u *UserUsecase) LogoutDevice(userID, deviceID string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.ErrNotFound
	}

	for i, rt := range user.RefreshTokens {
		if rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return domain.ErrFailedToUpdateUser
			}
			return &domain.CustomError{}
		}
	}

	return domain.ErrDeviceNotFound
}

func (u *UserUsecase) GetDevices(userID string) ([]string, *domain.CustomError) {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	var devices []string
	for _, rt := range user.RefreshTokens {
		devices = append(devices, rt.DeviceID)
	}
	return devices, &domain.CustomError{}
}

func (u *UserUsecase) RefreshToken(userID, deviceID, token string) (domain.LogInResponse, *domain.CustomError) {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrNotFound
	}

	for i, rt := range user.RefreshTokens {
		if rt.Token == token && rt.DeviceID == deviceID {
			user.RefreshTokens = append(user.RefreshTokens[:i], user.RefreshTokens[i+1:]...)

			refreshToken, err := u.TokenGen.GenerateRefreshToken(user)
			if err != nil {
				return domain.LogInResponse{}, domain.ErrInternalServer
			}

			newRefreshToken := domain.RefreshToken{
				Token:     refreshToken,
				DeviceID:  deviceID,
				CreatedAt: time.Now(),
			}

			user.RefreshTokens = append(user.RefreshTokens, newRefreshToken)
			err = u.UserRepo.UpdateUser(&user)
			if err != nil {
				return domain.LogInResponse{}, domain.ErrFailedToUpdateUser
			}

			accessToken, err := u.TokenGen.GenerateToken(user)
			if err != nil {
				return domain.LogInResponse{}, domain.ErrInternalServer
			}

			return domain.LogInResponse{
				AccessToken:  accessToken,
				RefreshToken: newRefreshToken.Token,
			}, &domain.CustomError{}
		}
	}

	return domain.LogInResponse{}, domain.ErrInvalidToken
}

func (u *UserUsecase) Register(user domain.User) *domain.CustomError {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return domain.ErrMissingRequiredFields
	}

	if !infrastracture.IsValidEmail(user.Email) {
		return domain.ErrInvalidEmail
	}

	if !infrastracture.IsValidPassword(user.Password) {
		return domain.ErrInvalidPassword
	}

	_, err := u.UserRepo.GetUserByUsernameOrEmail(user.Username, user.Email)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	user.Role = "user"

	// Hash password
	hashedPassword, err := u.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.ErrInternalServer
	}

	token, err := infrastracture.GenerateOTP()
	if err != nil {
		return domain.ErrInternalServer
	}

	user.Password = hashedPassword
	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	// Create user account in the database
	err = u.UserRepo.Register(user)
	if err != nil {
		return domain.ErrInternalServer
	}

	// Send activation email or link to the user
	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return domain.ErrFailedToSendEmail
	}

	return &domain.CustomError{}
}

func (u *UserUsecase) GetUserByUsernameOrEmail(username, email string) (domain.User, *domain.CustomError) {
	user, err := u.UserRepo.GetUserByUsernameOrEmail(username, email)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	return user, &domain.CustomError{}
}

func (u *UserUsecase) AccountActivation(token string, email string) *domain.CustomError {
	err := u.UserRepo.AccountActivation(token, email)
	if err != nil {
		return domain.ErrActivationFailed
	}
	return &domain.CustomError{}
}

func (u *UserUsecase) SendPasswordResetLink(email string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return domain.ErrNotFound
	}

	resetToken, err := infrastracture.GenerateActivationToken()
	if err != nil {
		return domain.ErrInternalServer
	}
	user.PasswordResetToken = resetToken
	user.TokenCreatedAt = time.Now()

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return domain.ErrFailedToUpdateUser
	}

	err = infrastracture.SendResetLink(user.Email, resetToken)
	if err != nil {
		return domain.ErrFailedToSendEmail
	}

	return &domain.CustomError{}
}

func (u *UserUsecase) ResetPassword(token, newPassword string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByResetToken(token)
	if err != nil {
		return domain.ErrInvalidToken
	}

	hashedPassword, err := u.PasswordSvc.HashPassword(newPassword)
	if err != nil {
		return domain.ErrInternalServer
	}

	user.Password = hashedPassword
	user.PasswordResetToken = ""
	user.TokenCreatedAt = time.Time{}

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return domain.ErrFailedToUpdateUser
	}

	return &domain.CustomError{}
}

func (u *UserUsecase) ActivateAccountMe(userID string) *domain.CustomError {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		return domain.ErrNotFound
	}

	token, err := infrastracture.GenerateActivationToken()
	if err != nil {
		return domain.ErrInternalServer
	}

	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		return domain.ErrFailedToUpdateUser
	}

	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return domain.ErrFailedToSendEmail
	}

	return &domain.CustomError{}
}

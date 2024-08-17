package usecase

import (
	"errors"
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

	user.RefreshTokens = append(user.RefreshTokens, newRefreshToken)
	err = u.UserRepo.UpdateUser(user)
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

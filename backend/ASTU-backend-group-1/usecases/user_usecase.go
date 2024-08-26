package usecase

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"astu-backend-g1/infrastructure"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) (domain.UserUsecase, error) {
	return &userUsecase{userRepository: u}, nil
}

func (useCase *userUsecase) Get() ([]domain.User, error) {
	return useCase.userRepository.Get(domain.UserFilterOption{})
}

func (useCase *userUsecase) LoginUser(uname string, password string, email string) (string, string, error) {
	if uname != "" {
		user, err := useCase.GetByUsername(uname)
		if err != nil {
			return "", "", err
		}
		if user.IsActive == false {
			return "", "", errors.New("Account not activated")
		}
		accesstoken, refreshToken, err := infrastructure.GenerateToken(&user, password)
		if err != nil {
			return "", "", err
		}
		user.RefreshToken = refreshToken
		useCase.userRepository.Update(user.ID, domain.User{RefreshToken: refreshToken, IsAdmin: user.IsAdmin})

		return accesstoken, refreshToken, nil
	} else if email != "" {
		user, err := useCase.GetByEmail(email)

		if err != nil {
			return "", "", err
		}
		if user.IsActive == false {
			return "", "", errors.New("Account not activated")
		}
		accesstoken, refreshToken, err := infrastructure.GenerateToken(&user, password)
		if err != nil {
			return "", "", err
		}
		log.Println(user)
		user.RefreshToken = refreshToken
		useCase.userRepository.Update(user.ID, domain.User{RefreshToken: refreshToken, IsAdmin: user.IsAdmin})

		return accesstoken, refreshToken, nil
	} else {
		return "", "", errors.New("invalid login credentials")
	}
}

func (useCase *userUsecase) Logout(email string) error {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return err
	}
	user.RefreshToken = ""
	useCase.userRepository.Update(user.ID, domain.User{RefreshToken: "", IsAdmin: user.IsAdmin})
	return nil
}

func (useCase *userUsecase) ForgetPassword(email string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user.IsActive == false {
		return "", errors.New("account not activated")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}

	user.VerifyToken = string(confirmationToken)

	expirationTime := time.Now().Add(2 * time.Hour)
	useCase.userRepository.Update(user.ID, domain.User{VerifyToken: string(confirmationToken), ExpirationDate: expirationTime, IsAdmin: user.IsAdmin})
	config_domain, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	link := config_domain.Domain + "/users/resetPassword/?email=" + user.Email + "&token=" + string(confirmationToken)
	err = infrastructure.SendEmail(user.Email, "Password Reset", "This is the password reset link: ", link)
	if err != nil {
		return "", err
	}

	return "Password reset token sent to your email", nil
}

func (useCase *userUsecase) ResetPassword(email string, token string, password string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", errors.New("account not activated")
	}
	if user.VerifyToken == token {
		if user.ExpirationDate.Before(time.Now()) {
			return "Token has expired", errors.New("token expired")
		}
		user.Password, _ = infrastructure.PasswordHasher(password)
		_, err := useCase.userRepository.Update(user.ID, domain.User{Password: user.Password, IsAdmin: user.IsAdmin})
		if err != nil {
			return "password has not been updated", err
		}
		return "Password reset successful", nil
	}
	return "Invalid token", errors.New("Invalid token")
}

func (useCase *userUsecase) GetByID(userID string) (domain.User, error) {
	filter := domain.UserFilter{UserId: userID}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) GetByUsername(username string) (domain.User, error) {
	filter := domain.UserFilter{Username: username}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

// function for account verification

func (useCase *userUsecase) AccountVerification(uemail string, confirmationToken string) error {
	user, err := useCase.GetByEmail(uemail)
	if err != nil {
		return err
	}
	if user.VerifyToken == confirmationToken {
		_, err := useCase.userRepository.Update(user.ID, domain.User{IsActive: true, IsAdmin: user.IsAdmin})
		return err
	} else {
		return errors.New("invalid token")
	}
}

func (useCase *userUsecase) GetByEmail(email string) (domain.User, error) {
	filter := domain.UserFilter{Email: email}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) Create(u *domain.User) (domain.User, error) {
	u.Password, _ = infrastructure.PasswordHasher(u.Password)
	u.IsActive = false
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}
	u.VerifyToken = string(confirmationToken)
	nUser, err := useCase.userRepository.Create(u)
	config, err := config.LoadConfig()
	if err != nil {
		return domain.User{}, err
	}
	if !nUser.IsAdmin {
		// link := "`http://localhost:8000/`users/accountVerification/?email=" + u.Email + "&token=" + string(confirmationToken)
		link := config.Domain + "/users/accountVerification/?email=" + u.Email + "&token=" + string(confirmationToken)
		err = infrastructure.SendEmail(u.Email, "Registration Confirmation", "This sign up Confirmation email to verify: ", link)
		if err != nil {
			return nUser, err
		}
	} else {
		infrastructure.SendEmail(nUser.Email, "Welcome! You Are Our First Admin", "Congratulations! As the first user to join our site, you have been automatically granted admin privileges. Thank you for being an early supporter.", config.Domain)
	}
	if err != nil {
		return domain.User{}, err
	}
	return nUser, err
}

func (useCase *userUsecase) UpdateUser(userId string, updateData domain.User) (domain.User, error) {
	user, err := useCase.GetByEmail(updateData.Email)
	updateData.IsAdmin = user.IsAdmin
	updateData.IsActive = user.IsActive
	updateData.Password = ""
	if err != nil {
		return user, err
	}
	return useCase.userRepository.Update(userId, updateData)
}

func (useCase *userUsecase) ChangePassword(email string, oldPassword string, newPassword string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", errors.New("account not activated")
	}
	oldPassword, _ = infrastructure.PasswordHasher(oldPassword)
	if user.Password == oldPassword {
		user.Password, _ = infrastructure.PasswordHasher(newPassword)
		user.Password = newPassword
		_, err := useCase.userRepository.Update(user.ID, domain.User{Password: newPassword, IsAdmin: user.IsAdmin})
		if err != nil {
			return "password has not been updated", err
		}
		return "Password change successful", nil
	}
	return "Invalid password", errors.New("invalid password")
}
func (useCase *userUsecase) Delete(userId string) error {
	return useCase.userRepository.Delete(userId)
}
func (useCase *userUsecase) PromteUser(username string) (domain.User, error) {
	user, err := useCase.GetByUsername(username)
	if user.IsAdmin {
		return user, fmt.Errorf("user is already an admin")
	}
	if !user.IsActive {
		return user, fmt.Errorf("the account is not activated")
	}
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	user, err = useCase.userRepository.Update(user.ID, domain.User{IsAdmin: true})
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	config, err := config.LoadConfig()
	if err != nil {
		return domain.User{}, err
	}
	infrastructure.SendEmail(user.Email, "Promotion To Admin", "This is a message for announcing you that you have been promoted to be an admin in our site", config.Domain)
	return user, nil
}
func (useCase *userUsecase) DemoteUser(username string) (domain.User, error) {
	user, err := useCase.GetByUsername(username)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	if !user.IsAdmin {
		return user, errors.New("user is not an admin")
	}

	user, err = useCase.userRepository.Update(user.ID, domain.User{IsAdmin: false})
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	config, err := config.LoadConfig()
	if err != nil {
		return domain.User{}, err
	}
	infrastructure.SendEmail(user.Email, "Demotion Notice", "We regret to inform you that you have been demoted from your admin position on our site. Thank you for your contributions and understanding.", config.Domain)
	return user, err
}
func (useCase *userUsecase) PromteUserByEmail(email string) (domain.User, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	if user.IsAdmin {
		return user, fmt.Errorf("user is already an admin")
	}

	if !user.IsActive {
		return user, fmt.Errorf("the account is not activated")
	}
	user, err = useCase.userRepository.Update(user.ID, domain.User{IsAdmin: true})
	if err != nil {
		return user, fmt.Errorf("user not found")
	}

	return user, err
}
func (useCase *userUsecase) DemoteUserByEmail(email string) (domain.User, error) {
	user, err := useCase.GetByEmail(email)
	if !user.IsAdmin {
		return user, errors.New("user is not an admin")
	}
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	user, err = useCase.userRepository.Update(user.ID, domain.User{IsAdmin: false})
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	return user, err
}

package Usecases

import (
    "time"
    "AAiT-backend-group-8/Domain"
    "AAiT-backend-group-8/infrastructure"
	  "crypto/rand"
    "encoding/hex"
	  "golang.org/x/crypto/bcrypt"
    "errors"

)


func generateVerificationToken() string {
    token := make([]byte, 16)
    _, err := rand.Read(token)
    if err != nil {
        return ""
    }
    return hex.EncodeToString(token) + time.Now().Format("20060102150405")
}
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

type UserUseCaseImpl struct {
    userRepository Domain.IUserRepository
}

func NewUserUseCase(userRepo Domain.IUserRepository) *UserUseCaseImpl {
    return &UserUseCaseImpl{userRepository: userRepo}
}



func (u *UserUseCaseImpl) VerifyEmail(token string) error {
    user, err := u.userRepository.GetUserByVerificationToken(token)
    if err != nil {
        return err 
    }

    if user.Verified {
        return errors.New("user is already verified")
    }

    user.Verified = true
    user.VerificationToken = "" 
    err = u.userRepository.VerifyUser(user)
    if err != nil {
        return err 
    }

    return nil 
}




func (u *UserUseCaseImpl) RegisterUser(user *Domain.User) error {
    // Check if this is the first user
    userCount, err := u.userRepository.GetUserCount()
    if err != nil {
        return err
    }

    if userCount == 0 {
        user.Role = "superadmin"
    } else {
        user.Role = "user" // Default role for non-first users
    }

    // Check if email already exists
    existingUser, err := u.userRepository.GetUserByEmail(user.Email)
    if err == nil && existingUser != nil {
        return errors.New("email already exists")
    }

    // Set other user details
    user.CreatedAt = time.Now()
    user.Verified = false
    user.VerificationToken = generateVerificationToken()

    // Hash password
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    err = u.userRepository.CreateUser(user)
    if err != nil {
        return err
    }

    err = infrastructure.SendVerificationEmail(user.Email, user.VerificationToken)
    if err != nil {
        return err
    }

    return nil
}





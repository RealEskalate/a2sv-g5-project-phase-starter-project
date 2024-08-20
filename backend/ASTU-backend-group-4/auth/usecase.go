package auth

import (
	"context"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserUsecase struct {
	repository AuthRepository
}

func NewAuthUserUsecase(repository AuthRepository) AuthUserUsecase {
	return AuthUserUsecase{
		repository: repository,
	}

}

func (au *AuthUserUsecase) Login(ctx context.Context, info LoginForm) (string, error) {
	// panic("not implemented") // TODO: Implement
	// var userInfo LoginForm
	// find the user name and match the hashed password with ith info.password
	user, err := au.repository.GetUserByUsername(ctx, info.Username)
	if err != nil {
		return "", ErrNoUserWithUsername
	}
	if !user.IsActive {
		return "", ErrAccountNotActivated
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password))
	if err != nil {
		return "", ErrIncorrectPassword
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username,
		"email":     user.Email,
		"isadmin":   user.IsAdmin,
		"isactive":  user.IsActive,
		"create":    user.CreatedAt,
		"updatedat": user.UpdatedAt,
	})
	jwtToken, err := token.SignedString(os.Getenv("SECRET_KEY"))

	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (au *AuthUserUsecase) RegisterUser(ctx context.Context, user User) error {
	// var newUser User
	user, err := au.repository.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return ErrUserExistWithThisEmail
	}
	user, err = au.repository.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return ErrUserExistWithThisUsername
	}

	hashedpasswors, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedpasswors)
	user.IsActive = false
	err = au.repository.CreateUser(ctx, user)
	if err != nil {
		return ErrCantCreateUser
	}
	return nil
	// id, err := au.repository.CreateUser(ctx, user)

}

func (au *AuthUserUsecase) UpdateProfile(ctx context.Context, user User) error {
	_, err := au.repository.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (au *AuthUserUsecase) Activate(userID string, token string) {

}

func (au *AuthUserUsecase) Logout(userID string) {
	panic("not implemented") // TODO: Implement
}

func (au *AuthUserUsecase) GenerateToken(user User) (string, error) {
	panic("not implemented") // TODO: Implement
}

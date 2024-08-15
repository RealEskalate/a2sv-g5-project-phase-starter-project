package usecases

import (
    "ASTU-backend-group-3/Blog_manager/Domain"
    "ASTU-backend-group-3/Blog_manager/infrastructure"
	"ASTU-backend-group-3/Blog_manager/Repository"
    "context"

)

type UserUsecase interface {
	Register (ctx context.Context, user *Domain.User) error
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, access_token, refresh_token string) error
} 

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) Register(ctx context.Context, user *Domain.User) error {
	// implementation of Register method
}
func (u *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepository.FindByEmail(email)
    if err!= nil {
        return " ", err
    }

	storedPassword := user.Password

	err = infrastructure.ComparePasswords(storedPassword, password)

	if err != nil{
		return  " ", err
	}

	access_token, err  := infrastructure.GenerateToken(user.Username , user.Role)

	if err != nil{
		return " ", err
	}
	refresh_token , err := infrastructure.GenerateRefreshToken(user.Username)

	if err != nil{
		return " ", err
	}

	err =  u.userRepository.InsertToken(user.Username , access_token , refresh_token)
	if err != nil{
		return " ", err
	}
    
	return access_token, nil
}

func (u *userUsecase) Logout(ctx context.Context,  username  string) error {
	u.userRepository.DeleteToken (username )
}


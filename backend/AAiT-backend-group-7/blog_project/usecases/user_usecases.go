package usecases

import (
	"blog_project/domain"
	"context"
	"errors"
	"regexp"
	"time"
)

type UserUsecase struct {
	UserRepo domain.IUserRepository
}

func NewUserUsecase(userRepo domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.UserRepo.GetAllUsers(ctx)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	return user, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	existingUser, err := u.UserRepo.SearchByEmail(ctx, user.Email)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	if existingUser.ID != 0 {
		return domain.User{}, errors.New("email already in use")
	}

	existingUser, err = u.UserRepo.SearchByUsername(ctx, user.Username)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	if existingUser.ID != 0 {
		return domain.User{}, errors.New("username already in use")
	}

	if !isValidEmail(user.Email) {
		return domain.User{}, errors.New("invalid email")

	}

	if !isValidPassword(user.Password) {
		return domain.User{}, errors.New("invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8")
	}

	// user.Password = infrastructure.HashPassword(user.Password)

	user.ID = int(time.Now().UnixNano() / 1000)

	return u.UserRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	return u.UserRepo.UpdateUser(ctx, id, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.UserRepo.DeleteUser(ctx, id)
}

func (u *UserUsecase) AddBlog(ctx context.Context, userID int, blogID int) (domain.User, error) {
	return u.UserRepo.AddBlog(ctx, userID, blogID)
}

func (u *UserUsecase) Login(ctx context.Context, username, password string) (domain.User, error) {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil || user.ID == 0 {
		return domain.User{}, errors.New("invalid credentials")
	}

	// // Assume infrastructure is implemented to verify passwords
	// if !infrastructure.VerifyPassword(user.Password, password) {
	// 	return domain.User{}, errors.New("invalid credentials")
	// }

	// token, err := infrastructure.GenerateToken(user.Username)
	// if err != nil {
	// 	return domain.User{}, err
	// }

	// _, err = u.UserRepo.CreateToken(ctx, user.Username, token)
	// if err != nil {
	// 	return domain.User{}, err
	// }

	return user, nil
}

func (u *UserUsecase) ForgetPassword(ctx context.Context, email string) error {
	user, err := u.UserRepo.SearchByEmail(ctx, email)
	if err != nil || user.ID == 0 {
		return errors.New("user not found")
	}

	// Assume infrastructure is implemented to send password reset emails
	// return infrastructure.SendResetLink(user.Email)
	return nil
}

func (u *UserUsecase) ResetPassword(ctx context.Context, username, password string) error {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil || user.ID == 0 {
		return errors.New("user not found")
	}

	// Assume infrastructure is implemented to hash passwords
	// hashedPassword, err := infrastructure.HashPassword(password)
	// if err != nil {
	// 	return err
	// }

	// user.Password = hashedPassword
	_, err = u.UserRepo.UpdateUser(ctx, user.ID, user)
	return err
}

func (u *UserUsecase) Promote(ctx context.Context, userid int) error {
	user, err := u.UserRepo.GetUserByID(ctx, userid)

	if err != nil {
		return err
	}

	user.Role = "admin"

	u.UserRepo.UpdateUser(ctx, userid, user)
	return nil

}

func (u *UserUsecase) Demote(ctx context.Context, userid int) error {
	user, err := u.UserRepo.GetUserByID(ctx, userid)

	if err != nil {
		return err
	}

	user.Role = "user"

	u.UserRepo.UpdateUser(ctx, userid, user)
	return nil
}

// Email validation function
func isValidEmail(email string) bool {
	// Regex pattern for valid email format
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Password strength validation function
func isValidPassword(password string) bool {
	// Require at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8
	const passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`
	re := regexp.MustCompile(passwordRegex)
	return re.MatchString(password)
}

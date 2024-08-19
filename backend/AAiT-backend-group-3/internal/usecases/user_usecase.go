package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/infrasctructures/services"
)
type UserUsecase struct {
	userRepo repository_interface.UserRepositoryInterface
	passwordService services.IHashService
}
func NewUserUsecase(userRepo repository_interface.UserRepositoryInterface, passwordService services.IHashService) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		passwordService: passwordService,
	}
}

func (u *UserUsecase) SignUp(user *models.User) error {

	if err := passwordService.ValidatePassword(user.Password); err != nil {
		return err
	}
	if err := passwordService.ValidateEmail(user.Email); err != nil {
		return err
	}
	encryptedPassword, err := passwordService.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword
	return u.userRepo.SignUp(user)
}

func (u *UserUsecase) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	if err := passwordService.ValidateEmail(email); err != nil {
		return nil, err
	}
	return u.userRepo.GetUserByEmail(email)
}

func (u *UserUsecase) DeleteUser(userID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(userID)
}

func (u *UserUsecase) UpdateProfile(userID primitive.ObjectID, user *models.User) error {
	if err := passwordService.ValidateEmail(user.Email); err != nil {
		return err
	}
	return u.userRepo.UpdateProfile(userID, user)
}
func (u *UserUsecase) PromoteUser(userID primitive.ObjectID) error {
	return u.userRepo.PromoteUser(userID)
}
func (u *UserUsecase) DemoteUser(userID primitive.ObjectID) error {
	return u.userRepo.DemoteUser(userID)
}

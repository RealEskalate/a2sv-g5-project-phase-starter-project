package usecases

import (
	"errors"

	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	repository_interface "AAIT-backend-group-3/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserUsecase struct {
	userRepo repository_interface.UserRepositoryInterface
	passwordService services.IHashService
	validationService services.IValidationService
	emailService services.IEmailService
	jwtSevices services.IJWT
}


func NewUserUsecase(userRepo repository_interface.UserRepositoryInterface, passwordService services.IHashService, validationService services.IValidationService, emailService services.IEmailService, jwtService services.IJWT) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		passwordService: passwordService,
		validationService: validationService,
		emailService: emailService,
		jwtSevices: jwtService,
	}
}


func (u *UserUsecase) SignUp(user *models.User) error {

	if _,err := u.validationService.ValidatePassword(user.Password); err != nil {
		return err
	}
	if _,err := u.validationService.ValidateEmail(user.Email); err != nil {
		return err
	}
	encryptedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword
	return u.userRepo.SignUp(user)
}

func (u *UserUsecase) Login(user *models.User) (string, string, error) {
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
		return "", "", err
	}
	existingUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	if !u.passwordService.CompareHash(existingUser.Password,user.Password, ) {
		return "","", errors.New("invalid password")
	}

	accessToken, _ := u.jwtSevices.GenerateAccessToken(existingUser.ID.Hex(), existingUser.Role)
	refershToken, _ := u.jwtSevices.GenerateRefreshToken(existingUser.ID.Hex(), existingUser.Role)

	return accessToken, refershToken, nil
}

func (u *UserUsecase) RefreshToken(userId primitive.ObjectID ,refreshTok string ) (string, error) {

	if _, err := u.jwtSevices.ValidateRefreshToken(refreshTok); err == nil {
		return "", errors.New("invalid token")
	}

	existingUser, err := u.userRepo.GetUserByID(userId)
	if err != nil {
		return "", errors.New("user not found")
	}

	if (existingUser.RefToken != refreshTok) {
		return "", errors.New("invalid token")
	}

	accessToken, _ := u.jwtSevices.GenerateAccessToken(existingUser.ID.Hex(), existingUser.Role)

	return accessToken, nil
}

func (u *UserUsecase) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	if _, err := u.validationService.ValidateEmail(email); err != nil {
		return nil, err
	}
	return u.userRepo.GetUserByEmail(email)
}

func (u *UserUsecase) DeleteUser(userID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(userID)
}

func (u *UserUsecase) UpdateProfile(userID primitive.ObjectID, user *models.User) error {
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
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

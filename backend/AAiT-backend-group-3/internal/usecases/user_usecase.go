package usecases

import (
	"errors"
	"fmt"
	"time"

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
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return  errors.New("error while fetching")
	}
	if len(users) == 0 {
		user.Role = "ADMIN"
	} else {
		user.Role = "USER"
		existingUser, _ := u.userRepo.GetUserByEmail(user.Email)
		if existingUser != nil {
			return errors.New("user already exists")
		}
	}

	if _, err := u.validationService.ValidatePassword(user.Password); err != nil {
		return err
	}
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
		return err
	}

	encryptedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.IsVerified = false 

	regUser, err := u.userRepo.SignUp(user)
	if err != nil {
		return errors.New(err.Error())
	}

	verificationToken, err := u.jwtSevices.GenerateVerificationToken(regUser.ID.Hex())	
	if err != nil {
		return errors.New("can't generate verification token")
	}
	verificationLink := fmt.Sprintf("https://localhost:8080/verify-email?token=%s", verificationToken)

	user.VerificationToken = verificationToken
	user.TokenExp = time.Now().Add(24 * time.Hour)

	u.userRepo.UpdateProfile(regUser.ID, user)

	subject := "Email Verification"
	body := fmt.Sprintf(`
		<h1>Email Verification</h1>
		<p>To verify your email, you can click on the following link:</p>
		<a href="%s">Verify email</a>
	`, verificationLink)

	err = u.emailService.SendVerificationEmail(regUser.Email, subject, body)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil 
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

func (u *UserUsecase) RefreshToken(refreshTok string ) (string, error) {
	userId, err := u.jwtSevices.ValidateRefreshToken(refreshTok)
	if err != nil {
		return "", errors.New(err.Error())
	}

	user_id, _ := primitive.ObjectIDFromHex(userId)
	existingUser, err := u.userRepo.GetUserByID(user_id)
	if err != nil {
		return "", errors.New("user not found")
	}

	if (existingUser.RefToken != refreshTok) {
		return "", errors.New("invalid token")
	}

	accessToken, _ := u.jwtSevices.GenerateAccessToken(existingUser.ID.Hex(), existingUser.Role)

	return accessToken, nil
}

func (u *UserUsecase) VerifyEmailToken(token string) (string, string, error) {
	userId, err := u.jwtSevices.ValidateVerificationToken(token)
	if err != nil {
		return "","", errors.New(err.Error())
	}

	user_id, _ := primitive.ObjectIDFromHex(userId)
	user, err := u.userRepo.GetUserByID(user_id)
	if err != nil {
		return "", "", errors.New("invalid or expired token")
	}

	if time.Now().After(user.TokenExp) {
		return "", "", errors.New("token expired")
	}

	accessToken, _ := u.jwtSevices.GenerateAccessToken(user.ID.Hex(), user.Role)
	refershToken, _ := u.jwtSevices.GenerateRefreshToken(user.ID.Hex(), user.Role)

	user.IsVerified = true
	user.VerificationToken = "" 
	user.TokenExp = time.Time{}
	user.RefToken = refershToken

	err = u.userRepo.UpdateProfile(user.ID, user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refershToken, nil
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

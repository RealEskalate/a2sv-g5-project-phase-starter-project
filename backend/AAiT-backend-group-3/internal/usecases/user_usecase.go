package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserUsecaseInterface interface {
	SignUp(user *models.User) (*models.User, error)
	Login(user *models.User) (string, string, error)
	Logout(token string) error
	RefreshToken(refreshToken string) (string, error)
	GetUserByID(userID string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(userID string) error
	UpdateProfile(userID string, user *models.User) error
	PromoteUser(userID string) error
	DemoteUser(userID string) error
	VerifyEmailToken(token string) (string, string, error)
	
}

type UserUsecase struct {
	userRepo repository_interface.UserRepositoryInterface
	passwordService services.IHashService
	validationService services.IValidationService
	emailService services.IEmailService
	jwtSevices services.IJWT
}

func NewUserUsecase(userRepo repository_interface.UserRepositoryInterface, passwordService services.IHashService, validationService services.IValidationService, emailService services.IEmailService, jwtService services.IJWT) UserUsecaseInterface {
	return &UserUsecase{
		userRepo: userRepo,
		passwordService: passwordService,
		validationService: validationService,
		emailService: emailService,
		jwtSevices: jwtService,
	}
}


func (u *UserUsecase) SignUp(user *models.User) (*models.User, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		user.Role = "ADMIN"
	} else {
		existingUser, _ := u.userRepo.GetUserByEmail(user.Email)
		if existingUser != nil {
			return nil, err
		}
		user.Role = "USER"
	}

	if _, err := u.validationService.ValidatePassword(user.Password); err != nil {
		return nil, err
	}
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
		return nil, err
	}

	encryptedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.IsVerified = false 

	regUser, err := u.userRepo.SignUp(user)
	if err != nil {
		return nil, err
	}

	verificationToken, err := u.jwtSevices.GenerateVerificationToken(regUser.ID.Hex())	
	if err != nil {
		return nil, err
	}
	verificationLink := fmt.Sprintf("http://localhost:8080/auth/verify-email?token=%s", verificationToken)

	user.VerificationToken = verificationToken


	u.userRepo.UpdateProfile(regUser.ID.Hex(), user)

	err = u.emailService.SendVerificationEmail(user.Email, verificationLink)
	if err != nil {
		return nil, err
	}
	return user, nil 
}


func (u *UserUsecase) Login(user *models.User) (string, string, error) {
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
		return "", "", errors.New(err.Error())
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

	existingUser.RefToken = refershToken
	err = u.userRepo.UpdateProfile(existingUser.ID.Hex(), existingUser)
	if  err != nil {
		return "", "", errors.New(err.Error())
	}
	return accessToken, refershToken, nil
}


func (u *UserUsecase) Logout(token string) error {
	parsedToken, err := u.jwtSevices.ValidateAccessToken(token)
	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	expiration, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid expiration time in token claims")
	}

	remTime := time.Until(time.Unix(int64(expiration), 0))
	if remTime <= 0 {
		return errors.New("token has already expired")
	}

	err = u.userRepo.BlacklistToken(token, remTime)
	if err != nil {
		return err
	}

	existingUser, err := u.userRepo.GetUserByID(claims["userId"].(string))
	if err != nil {
		return err
	}

	existingUser.RefToken = ""
	err = u.userRepo.UpdateProfile(existingUser.ID.Hex(), existingUser)
	if err != nil {
		return errors.New("failed to update user profile")
	}
	return nil
}

func (u *UserUsecase) RefreshToken(refreshTok string ) (string, error) {
	userId, err := u.jwtSevices.ValidateRefreshToken(refreshTok)
	if err != nil {
		return "", errors.New(err.Error())
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

func (u *UserUsecase) VerifyEmailToken(token string) (string, string, error) {
	userId, err := u.jwtSevices.ValidateVerificationToken(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			claims, ok := u.jwtSevices.GetClaimsFromToken(token)
			if ok {
				userId, _ := claims["userId"].(string)
				user, err := u.userRepo.GetUserByID(userId)
				if err != nil {
					return "", "", errors.New("invalid token")
				}
				if user.IsVerified {
					return "", "", errors.New("user already verified")
				}

				verificationToken, _ := u.jwtSevices.GenerateVerificationToken(user.ID.Hex())
				verificationLink := fmt.Sprintf("http://localhost:8080/auth/verify-email?token=%s", verificationToken)

				user.VerificationToken = verificationToken
				err = u.userRepo.UpdateProfile(user.ID.Hex(), user)
				if err != nil {
					return "", "", err
				}

				err = u.emailService.SendVerificationEmail(user.Email, verificationLink)
				if err != nil {
					return "", "", err
				}

				return "", "", errors.New("verification token expired. A new verification email has been sent")
			}
		}
		return "", "", errors.New(err.Error())
	}

	user, err := u.userRepo.GetUserByID(userId)
	if err != nil{
		return "", "", errors.New("invalid token")
	} else if user.IsVerified {
		return "", "", errors.New("user already verified")
	} else if  token != user.VerificationToken {
		return "", "", errors.New("invalid token")
	}
	user.IsVerified = true
	user.VerificationToken = ""

	accessToken, _ := u.jwtSevices.GenerateAccessToken(user.ID.Hex(), user.Role)
	refershToken, _ := u.jwtSevices.GenerateRefreshToken(user.ID.Hex(), user.Role)

	user.RefToken = refershToken
	err = u.userRepo.UpdateProfile(user.ID.Hex(), user)
	if  err != nil {
		return "", "", errors.New(err.Error())
	}
	return accessToken, refershToken, nil
}


func (u *UserUsecase) GetUserByID(userID string) (*models.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	if _, err := u.validationService.ValidateEmail(email); err != nil {
		return nil, err
	}
	return u.userRepo.GetUserByEmail(email)
}

func (u *UserUsecase) DeleteUser(userID string) error {
	return u.userRepo.DeleteUser(userID)
}

func (u *UserUsecase) UpdateProfile(userID string, user *models.User) error {
	if _, err := u.validationService.ValidateEmail(user.Email); err != nil {
		return err
	}
	return u.userRepo.UpdateProfile(userID, user)
}
func (u *UserUsecase) PromoteUser(userID string) error {
	return u.userRepo.PromoteUser(userID)
}
func (u *UserUsecase) DemoteUser(userID string) error {
	return u.userRepo.DemoteUser(userID)
}
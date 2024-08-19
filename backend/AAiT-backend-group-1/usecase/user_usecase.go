package usecase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure/mail"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserUseCase struct {
	userRepo        domain.UserRepository
	sessionRepo     domain.SessionRepository
	passwordService infrastructure.PasswprdService
	jwtService      infrastructure.JWTTokenService
	mailService     mail.EmailService
}

func NewUserUseCase(userRespository domain.UserRepository, sessionRepository domain.SessionRepository, pwdService infrastructure.PasswprdService, jwtService infrastructure.JWTTokenService, mailServ mail.EmailService) UserUseCase {
	return UserUseCase{
		userRepo:        userRespository,
		sessionRepo:     sessionRepository,
		passwordService: pwdService,
		jwtService:      jwtService,
		mailService:     mailServ,
	}
}

func (userUC *UserUseCase) Register(cxt *gin.Context, user *domain.User) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	userRepo domain.UserRepository
}

func NewUserUseCase(userRespository domain.UserRepository) UserUseCase {
	return UserUseCase{
		userRepo: userRespository,
	}
}

func (userUC *UserUseCase) Register(cxt *gin.Context, user *domain.User) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()
	errValidity := user.Validate()
	if errValidity != nil {
		return &domain.CustomError{Message: errValidity.Error(), Code: http.StatusBadRequest}
	}
	if _, errRepo := userUC.userRepo.Create(context, user); errRepo != nil {
		return errRepo
	}

	verificationToken, errVerification := userUC.jwtService.GenerateVerificationToken(user.Email)
	if errVerification != nil {
		return &domain.CustomError{Message: errVerification.Error(), Code: http.StatusInternalServerError}
	}

	userUC.sessionRepo.CreateToken(context, &domain.Session{
		Username:          user.Username,
		VerificationToken: verificationToken,
	})

	errEmail := userUC.mailService.SendVerificationEmail(user.Email, user.Username, fmt.Sprintf("http://localhost:8080/verify/%s", verificationToken))

	if errEmail != nil {
		return &domain.CustomError{Message: errEmail.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

func (userUC *UserUseCase) Login(cxt context.Context, username, password string) (map[string]string, domain.Error) {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return map[string]string{}, &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}


// 	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
// 	defer cancel()

	existingUser, errExisting := userUC.userRepo.FindByUsername(context, username)
	if errExisting != nil {
		return map[string]string{}, errExisting
	}

	validPassword, errHash := userUC.passwordService.ComparePassword(existingUser.Password, password)
	if errHash != nil {
		return map[string]string{}, errExisting
	}

	if !validPassword {
		return map[string]string{}, &domain.CustomError{Message: "Invalid Password", Code: http.StatusUnauthorized}
	}

	existingSession, exsistingCheck, errSession := userUC.sessionRepo.FindTokenByUserUsername(context, username)
	if errSession != nil {
		return map[string]string{}, &domain.CustomError{Message: "error while retriving existingSession", Code: http.StatusUnauthorized}
	}

	if exsistingCheck {
		errDelete := userUC.sessionRepo.DeleteToken(context, existingSession.ID.Hex())
		if errDelete != nil {
			return map[string]string{}, &domain.CustomError{Message: "error while deleting existingSession", Code: http.StatusUnauthorized}
		}
	}

	accessToken, errAccess := userUC.jwtService.GenerateAccessTokenWithPayload(*existingUser)
	if errAccess != nil {
		return map[string]string{}, &domain.CustomError{Message: errAccess.Error(), Code: http.StatusInternalServerError}
	}

	refreshToken, errRefresh := userUC.jwtService.GenerateRefreshTokenWithPayload(*existingUser)
	if errRefresh != nil {
		return map[string]string{}, &domain.CustomError{Message: errRefresh.Error(), Code: http.StatusInternalServerError}
	}
	return map[string]string{"access_token": accessToken, "refresh_token": refreshToken}, nil
}

func (userUC *UserUseCase) ForgotPassword(cxt context.Context, email string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	existingUser, err := userUC.userRepo.FindByEmail(context, email)
	if err != nil {
		return err
	}

	passwordResetToken, errToken := userUC.jwtService.GenerateVerificationToken(email)
	if errToken != nil {
		return &domain.CustomError{Message: errToken.Error(), Code: http.StatusInternalServerError}
	}

	existingSession, existingCheck, errSession := userUC.sessionRepo.FindTokenByUserUsername(context, existingUser.Username)
	if errSession != nil {
		return errSession
	}

	if existingCheck {
		errDelete := userUC.sessionRepo.DeleteToken(context, existingSession.ID.Hex())
		if errDelete != nil {
			return errDelete
		}
	}

	_, errCreatingToken := userUC.sessionRepo.CreateToken(context, &domain.Session{
		ID:                 existingUser.ID,
		Username:           existingSession.Username,
		PasswordResetToken: passwordResetToken,
	})

	if errCreatingToken != nil {
		return errCreatingToken
	}

	errEmail := userUC.mailService.SendPasswordResetEmail(email, existingUser.Username, fmt.Sprintf("http://localhost:8080/reset/%s", passwordResetToken))
	if errEmail != nil {
		return &domain.CustomError{Message: errEmail.Error(), Code: http.StatusInternalServerError}
	}

	return nil
}

func (userUC *UserUseCase) Logout(cxt context.Context, token map[string]string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	accessToken, ok := token["access_token"]
	if !ok {
		return &domain.CustomError{Message: "access token not found", Code: http.StatusBadRequest}
	}
	_, ok = token["refresh_token"]
	if !ok {
		return &domain.CustomError{Message: "refresh token not found", Code: http.StatusBadRequest}
	}

	// invalidate the access token by adding it to the blacklist, redis can be used for this
	parsedToken, err := userUC.jwtService.ValidateAccessToken(accessToken)
	if err != nil {
		return &domain.CustomError{Message: err.Error(), Code: http.StatusUnauthorized}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return &domain.CustomError{Message: "error parsing claims", Code: http.StatusInternalServerError}
	}

	existingRefreshToken, existenceCheck, err := userUC.sessionRepo.FindTokenByUserUsername(context, claims["username"].(string))
	if err != nil {
		return &domain.CustomError{Message: "error while retriving existingSession", Code: http.StatusUnauthorized}
	}

	if existenceCheck {
		errDelete := userUC.sessionRepo.DeleteToken(context, existingRefreshToken.ID.Hex())
		if errDelete != nil {
			return &domain.CustomError{Message: "error while deleting existingSession", Code: http.StatusUnauthorized}
		}
	}

	return nil
}

func (userUC *UserUseCase) PromoteUser(cxt context.Context, userID string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	promotion := domain.User{
		Role: "admin",
	}
	return userUC.userRepo.Update(context, userID, &promotion)
}

func (userUC *UserUseCase) DemoteUser(cxt context.Context, userID string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	promotion := domain.User{
		Role: "user",
	}
	return userUC.userRepo.Update(context, userID, &promotion)
}

func (userUC *UserUseCase) UpdateProfile(cxt context.Context, userID string, user *domain.User) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	return userUC.userRepo.Update(context, userID, user)
}


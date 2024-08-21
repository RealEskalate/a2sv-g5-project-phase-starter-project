package usecases

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure/mail"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase struct {
	userRepo        domain.UserRepository
	sessionRepo     domain.SessionRepository
	passwordService infrastructure.PasswprdService
	jwtService      infrastructure.JWTTokenService
	mailService     mail.EmailService
	redisService    domain.CacheService
}

func NewUserUseCase(userRespository domain.UserRepository, sessionRepository domain.SessionRepository, pwdService infrastructure.PasswprdService, jwtService infrastructure.JWTTokenService, mailServ mail.EmailService, redisServer domain.CacheService) UserUseCase {
	return UserUseCase{
		userRepo:        userRespository,
		sessionRepo:     sessionRepository,
		passwordService: pwdService,
		jwtService:      jwtService,
		mailService:     mailServ,
		redisService:    redisServer,
	}
}

func (userUC *UserUseCase) RegisterStart(cxt *gin.Context, user *domain.User) domain.Error {
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

	existingUser, errRepo := userUC.userRepo.FindByEmail(context, user.Email)
	if errRepo != nil {
		return errRepo
	}
	if existingUser.Email != "" {
		return &domain.CustomError{Message: "Email already exists", Code: http.StatusBadRequest}
	}
	existingUser, errRepo = userUC.userRepo.FindByUsername(context, user.Username)
	if errRepo != nil {
		return errRepo
	}
	if existingUser.Username != "" {
		return &domain.CustomError{Message: "Username already exists", Code: http.StatusBadRequest}
	}

	verificationToken, errVerification := userUC.jwtService.GenerateVerificationToken(*user)
	if errVerification != nil {
		return &domain.CustomError{Message: errVerification.Error(), Code: http.StatusInternalServerError}
	}

	userUC.sessionRepo.CreateToken(context, &domain.Session{
		Username:          user.Username,
		VerificationToken: verificationToken,
	})

	errEmail := userUC.mailService.SendVerificationEmail(user.Email, user.Username, fmt.Sprintf(os.Getenv("VERIFY_PATH"), verificationToken))

	if errEmail != nil {
		return &domain.CustomError{Message: errEmail.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

func (userUC *UserUseCase) RegisterEnd(cxt *gin.Context, token string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	parsedToken, errParse := userUC.jwtService.ValidateVerificationToken(token)
	if errParse != nil {
		return &domain.CustomError{Message: errParse.Error(), Code: http.StatusBadRequest}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return &domain.CustomError{Message: "error parsing claims", Code: http.StatusInternalServerError}
	}

	newUser := domain.User{
		Username:  claims["username"].(string),
		Email:     claims["email"].(string),
		Role:      claims["role"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  claims["password"].(string),
	}

	_, errCreateUser := userUC.userRepo.Create(context, &newUser)
	if errCreateUser != nil {
		return errCreateUser
	}

	return nil
}

func (userUC *UserUseCase) Login(cxt *gin.Context, username, password string) (map[string]string, domain.Error) {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return map[string]string{}, &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}

	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

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

func (userUC *UserUseCase) ForgotPassword(cxt *gin.Context, email string) domain.Error {
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

	passwordResetToken, errToken := userUC.jwtService.GenerateResetToken(email)
	if errToken != nil {
		return &domain.CustomError{Message: errToken.Error(), Code: http.StatusInternalServerError}
	}

	existingSession, existingCheck, errSession := userUC.sessionRepo.FindTokenByUserUsername(context, existingUser.Username)
	if errSession != nil {
		return errSession
	}

	if existingCheck {
		errUpdate := userUC.sessionRepo.UpdateToken(context, existingSession.ID.Hex(), &domain.Session{
			PasswordResetToken: passwordResetToken,
			RefreshToken:       "",
		})
		if errUpdate != nil {
			return errUpdate
		}
	} else {
		_, errCreatingToken := userUC.sessionRepo.CreateToken(context, &domain.Session{
			ID:                 existingUser.ID,
			Username:           existingSession.Username,
			PasswordResetToken: passwordResetToken,
		})
		if errCreatingToken != nil {
			return errCreatingToken
		}
	}

	errEmail := userUC.mailService.SendPasswordResetEmail(email, existingUser.Username, fmt.Sprintf("http://localhost:8080/reset/%s", passwordResetToken))
	if errEmail != nil {
		return &domain.CustomError{Message: errEmail.Error(), Code: http.StatusInternalServerError}
	}

	return nil
}

func (userUC *UserUseCase) ResetPassword(newPassword, confirmPassword, token string, cxt *gin.Context) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	parsedToken, errParse := userUC.jwtService.ValidateResetToken(token)
	if errParse != nil {
		return &domain.CustomError{Message: "error parsing token", Code: http.StatusInternalServerError}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return &domain.CustomError{Message: "error parsing claims", Code: http.StatusInternalServerError}
	}

	existingUser, errSearch := userUC.userRepo.FindByEmail(context, claims["email"].(string))
	if errSearch != nil {
		return errSearch
	}
	session, existenceCheck, errSession := userUC.sessionRepo.FindTokenByUserUsername(context, existingUser.Username)
	if errSession != nil {
		return &domain.CustomError{Message: "error while fetching session", Code: http.StatusInternalServerError}
	}

	if !existenceCheck {
		return &domain.CustomError{Message: "session not found", Code: http.StatusNotFound}
	}

	if session.PasswordResetToken != token {
		return &domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized}
	}

	if newPassword != confirmPassword {
		return &domain.CustomError{Message: "passwords do not match", Code: http.StatusBadRequest}
	}

	hashedPassword, errHash := userUC.passwordService.HashPassword(newPassword)
	if errHash != nil {
		return &domain.CustomError{Message: errHash.Error(), Code: http.StatusInternalServerError}
	}

	errUpdate := userUC.userRepo.Update(context, existingUser.ID.Hex(), &domain.User{Password: hashedPassword})
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (userUC *UserUseCase) Logout(cxt *gin.Context, token map[string]string) domain.Error {
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
	errRedis := userUC.redisService.Set(accessToken, "blacklisted", time.Minute*15)
	if errRedis != nil {
		return &domain.CustomError{Message: errRedis.Error(), Code: http.StatusInternalServerError}
	}

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
		errDelete := userUC.sessionRepo.UpdateToken(context, existingRefreshToken.ID.Hex(), &domain.Session{
			RefreshToken: "",
		})
		if errDelete != nil {
			return &domain.CustomError{Message: "error while deleting existingSession", Code: http.StatusUnauthorized}
		}
	}

	return nil
}

func (userUC *UserUseCase) PromoteUser(cxt *gin.Context, userID string) domain.Error {
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

func (userUC *UserUseCase) DemoteUser(cxt *gin.Context, userID string) domain.Error {
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

func (userUC *UserUseCase) UpdateProfile(cxt *gin.Context, userID string, user *domain.User) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	return userUC.userRepo.Update(context, userID, user)
}

func (userUC *UserUseCase) ImageUpload(cxt *gin.Context, file *multipart.File, header *multipart.FileHeader) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	if valid := utils.IsValidFileFormat(header, "image/png", "image/jpeg"); !valid {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid file format"})
	}
	filename := fmt.Sprintf("%s%s", uuid.New().String(), header.Filename)

	delivery_url, errSave := utils.SaveImage(*file, filename, context)
	if errSave != nil {
		return &domain.CustomError{Message: errSave.Error(), Code: http.StatusInternalServerError}
	}

	userID, _ := cxt.Get("user_id")
	primimitiveUserID, _ := primitive.ObjectIDFromHex(userID.(string))
	photo := domain.Photo{
		UserID:     primimitiveUserID,
		Filename:   filename,
		FilePath:   delivery_url,
		UploadedAt: time.Now(),
	}

	user, errUser := userUC.userRepo.FindById(context, userID.(string))
	if errUser != nil {
		return &domain.CustomError{Message: errUser.Error(), Code: http.StatusInternalServerError}
	}
	user.ProfilePictureUrl = photo
	return nil
}

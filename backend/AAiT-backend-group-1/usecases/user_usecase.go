package usecases

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userUseCase struct {
	userRepo        domain.UserRepository
	sessionRepo     domain.SessionRepository
	passwordService domain.PasswordService
	jwtService      domain.JwtService
	mailService     domain.EmailService
	redisService    domain.CacheService
}

type ResetPasswordRequest struct {
	NewPasswor      string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string
}

func NewUserUseCase(userRespository domain.UserRepository, sessionRepository domain.SessionRepository, pwdService domain.PasswordService, jwtService domain.JwtService, mailServ domain.EmailService, redisServer domain.CacheService) domain.UserUseCase {
	return &userUseCase{
		userRepo:        userRespository,
		sessionRepo:     sessionRepository,
		passwordService: pwdService,
		jwtService:      jwtService,
		mailService:     mailServ,
		redisService:    redisServer,
	}
}

func (userUC *userUseCase) RegisterStart(cxt *gin.Context, user *domain.User) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: "first" + errTimeout.Error(), Code: http.StatusInternalServerError}
	}

	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	errValidity := user.Validate()
	if errValidity != nil {
		return &domain.CustomError{Message: "second" + errValidity.Error(), Code: http.StatusBadRequest}
	}

	existingEmail, errEmail := userUC.userRepo.FindByEmail(context, user.Email)
	existingUsername, errUsername := userUC.userRepo.FindByUsername(context, user.Username)
	if (errEmail != nil && errEmail.StatusCode() != http.StatusNotFound) || (errUsername != nil && errUsername.StatusCode() != http.StatusNotFound) {
		return domain.CustomError{Message: "error while checking for existing user", Code: http.StatusInternalServerError}
	}
	if (existingEmail != nil && existingEmail.Email != "") || (existingUsername != nil && existingUsername.Username != "") {
		return domain.CustomError{Message: "Email or Username already exists", Code: http.StatusBadRequest}
	}

	user.Role = "user"
	hashedPassword, errHash := userUC.passwordService.HashPassword(user.Password)
	if errHash != nil {
		return domain.CustomError{Message: "error hashing password", Code: http.StatusInternalServerError}
	}

	user.Password = hashedPassword

	verificationToken, errVerification := userUC.jwtService.GenerateVerificationToken(*user)
	if errVerification != nil {
		return &domain.CustomError{Message: errVerification.Error() + "third", Code: http.StatusInternalServerError}
	}

	existingRegisteringUser, checkExistingSession, errRegisteringUser := userUC.sessionRepo.FindTokenByUserUsername(context, user.Username)
	if errRegisteringUser != nil {
		return domain.CustomError{Message: errRegisteringUser.Error(), Code: errRegisteringUser.StatusCode()}
	}

	if checkExistingSession {
		existingRegisteringUser.VerificationToken = verificationToken
		errUpdate := userUC.sessionRepo.UpdateToken(context, existingRegisteringUser.ID.Hex(), existingRegisteringUser)
		if errUpdate != nil {
			return domain.CustomError{Message: errUpdate.Error(), Code: errUpdate.StatusCode()}
		}
	} else {
		_, errStoreSession := userUC.sessionRepo.CreateToken(context, &domain.Session{
			Username:          user.Username,
			VerificationToken: verificationToken,
		})
		if errStoreSession != nil {
			return domain.CustomError{Message: errStoreSession.Error(), Code: errStoreSession.StatusCode()}
		}
	}

	err := userUC.mailService.SendVerificationEmail(user.Email, user.Username, fmt.Sprintf(os.Getenv("VERIFY_PATH"), verificationToken))

	if err != nil {
		return &domain.CustomError{Message: err.Error() + "forth", Code: http.StatusInternalServerError}
	}
	return nil
}

func (userUC *userUseCase) RegisterEnd(cxt *gin.Context, token string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error() + "five", Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	parsedToken, errParse := userUC.jwtService.ValidateVerificationToken(token)
	if errParse != nil {
		return &domain.CustomError{Message: errParse.Error() + "six", Code: http.StatusBadRequest}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return &domain.CustomError{Message: "error parsing claims", Code: http.StatusInternalServerError}
	}

	storedSession, sessionExist, err := userUC.sessionRepo.FindTokenByUserUsername(context, claims["username"].(string))
	if err != nil {
		return err
	}

	if !sessionExist {
		return domain.CustomError{Message: "session not found", Code: http.StatusNotFound}
	}
	if storedSession.VerificationToken != token {
		return domain.CustomError{Message: "invalid token from the link", Code: http.StatusUnauthorized}
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

func (userUC *userUseCase) Login(cxt *gin.Context, username, password string) (map[string]string, domain.Error) {
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
		return map[string]string{}, domain.CustomError{Message: errHash.Error(), Code: http.StatusInternalServerError}
	}

	if !validPassword {
		return map[string]string{}, &domain.CustomError{Message: "Invalid Password", Code: http.StatusUnauthorized}
	}

	existingSession, exsistingCheck, errSession := userUC.sessionRepo.FindTokenByUserUsername(context, username)
	if errSession != nil {
		return map[string]string{}, &domain.CustomError{Message: "error while retriving existingSession", Code: http.StatusUnauthorized}
	}

	if !exsistingCheck {
		return map[string]string{}, &domain.CustomError{Message: "session not found", Code: http.StatusNotFound}
	}

	accessToken, errAccess := userUC.jwtService.GenerateAccessTokenWithPayload(*existingUser)
	if errAccess != nil {
		return map[string]string{}, &domain.CustomError{Message: errAccess.Error(), Code: http.StatusInternalServerError}
	}

	refreshToken, errRefresh := userUC.jwtService.GenerateRefreshTokenWithPayload(*existingUser)
	if errRefresh != nil {
		return map[string]string{}, &domain.CustomError{Message: errRefresh.Error(), Code: http.StatusInternalServerError}
	}

	errUpdate := userUC.sessionRepo.UpdateToken(context, existingSession.ID.Hex(), &domain.Session{
		RefreshToken:       refreshToken,
		VerificationToken:  existingSession.VerificationToken,
		PasswordResetToken: existingSession.PasswordResetToken,
		Username:           existingSession.Username,
	})

	if errUpdate != nil {
		return map[string]string{}, domain.CustomError{Message: errUpdate.Error(), Code: http.StatusInternalServerError}
	}
	return map[string]string{"access_token": accessToken, "refresh_token": refreshToken}, nil
}

func (userUC *userUseCase) RefreshToken(cxt *gin.Context, refreshToken string) (map[string]string, domain.Error) {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return map[string]string{}, &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}

	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	token, errToken := userUC.jwtService.ValidateRefreshToken(refreshToken)
	if errToken != nil {
		return map[string]string{}, domain.CustomError{Message: errToken.Error(), Code: errToken.StatusCode()}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]string{}, domain.CustomError{Message: "error parsing claims", Code: http.StatusInternalServerError}
	}

	existingUser, err := userUC.userRepo.FindById(context, claims["user_id"].(string))
	if err != nil {
		return map[string]string{}, domain.CustomError{Message: err.Error(), Code: err.StatusCode()}
	}

	existingSession, existenceCheck, err := userUC.sessionRepo.FindTokenByUserUsername(context, existingUser.Username)
	if err != nil {
		return map[string]string{}, domain.CustomError{Message: "error while retriving existingSession", Code: http.StatusInternalServerError}
	}

	if !existenceCheck {
		return map[string]string{}, domain.CustomError{Message: "session not found", Code: http.StatusNotFound}
	}

	if existingSession.RefreshToken != refreshToken || existingSession.Username != existingUser.Username {
		return map[string]string{}, domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized}
	}

	newAccessToken, errAccess := userUC.jwtService.GenerateAccessTokenWithPayload(*existingUser)
	if errAccess != nil {
		return map[string]string{}, domain.CustomError{Message: errAccess.Error(), Code: http.StatusInternalServerError}
	}

	return map[string]string{"refresh_token": newAccessToken}, nil
}

func (userUC *userUseCase) ForgotPassword(cxt *gin.Context, email string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	existingUser, err := userUC.userRepo.FindByEmail(context, email)
	if err != nil {
		return domain.CustomError{Message: err.Error(), Code: http.StatusInternalServerError}
	}

	resetCode, errCode := utils.GenerateTokenWithLength(6)
	if errCode != nil {
		return &domain.CustomError{Message: errCode.Error(), Code: http.StatusInternalServerError}
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
		existingSession.PasswordResetToken = passwordResetToken
		existingSession.RefreshToken = ""
		errUpdate := userUC.sessionRepo.UpdateToken(context, existingSession.ID.Hex(), existingSession)
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

	errorEmail := userUC.mailService.SendPasswordResetEmail(existingUser.Email, existingSession.Username, fmt.Sprintf(os.Getenv("RESET_PATH"), passwordResetToken, resetCode), fmt.Sprintf("%v", resetCode))

	if errorEmail != nil {
		return &domain.CustomError{Message: errorEmail.Error(), Code: http.StatusInternalServerError}
	}

	return nil
}

func (userUC *userUseCase) ResetPassword(cxt *gin.Context, newPassword, confirmPassword, token string, code int) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	parsedToken, errParse := userUC.jwtService.ValidateResetToken(token)
	if errParse != nil {
		return &domain.CustomError{Message: "error parsing token " + errParse.Error(), Code: http.StatusInternalServerError}
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

	if session.PasswordResetToken != token || session.ResetPasswordToken != code {
		return &domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized}
	}

	if newPassword != confirmPassword {
		return &domain.CustomError{Message: "passwords do not match", Code: http.StatusBadRequest}
	}

	hashedPassword, errHash := userUC.passwordService.HashPassword(newPassword)
	if errHash != nil {
		return &domain.CustomError{Message: errHash.Error(), Code: http.StatusInternalServerError}
	}
	errUpdate := userUC.userRepo.UpdatePassword(context, existingUser.ID.Hex(), hashedPassword)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (userUC *userUseCase) Logout(cxt *gin.Context, token map[string]string) domain.Error {
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
	refreshToken, ok := token["refresh_token"]
	if !ok {
		return &domain.CustomError{Message: "refresh token not found", Code: http.StatusBadRequest}
	}

	// invalidate the access token by adding it to the blacklist, redis can be used for this
	errRedis := userUC.redisService.Set(accessToken, accessToken, time.Minute*15)
	if errRedis != nil {
		return &domain.CustomError{Message: errRedis.Error(), Code: http.StatusInternalServerError}
	}

	_, err := userUC.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return &domain.CustomError{Message: err.Error(), Code: http.StatusUnauthorized}
	}

	username, ok := cxt.Get("username")
	if !ok {
		return &domain.CustomError{Message: "username not found", Code: http.StatusBadRequest}
	}
	existingSession, existenceCheck, err := userUC.sessionRepo.FindTokenByUserUsername(context, username.(string))
	if err != nil {
		return &domain.CustomError{Message: "error while retriving existingSession " + err.Error(), Code: http.StatusUnauthorized}
	}

	if !existenceCheck {
		return domain.CustomError{Message: "session not found", Code: http.StatusNotFound}
	}

	existingSession.RefreshToken = ""
	existingSession.VerificationToken = ""
	existingSession.PasswordResetToken = ""
	errDelete := userUC.sessionRepo.UpdateToken(context, existingSession.ID.Hex(), existingSession)
	if errDelete != nil {
		return &domain.CustomError{Message: "error while updating existing session" + errDelete.Error(), Code: http.StatusUnauthorized}
	}
	return nil
}

func (userUC *userUseCase) PromoteUser(cxt *gin.Context, userID string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	role := "admin"
	return userUC.userRepo.UpdateRole(context, userID, role)
}

func (userUC *userUseCase) DemoteUser(cxt *gin.Context, userID string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	role := "user"
	return userUC.userRepo.UpdateRole(context, userID, role)
}

func (userUC *userUseCase) UpdateProfile(cxt *gin.Context, userID string, user map[string]interface{}) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	allowedFeilds := map[string]interface{}{}
	existenceCount, errCheckExistence := userUC.userRepo.CheckExistence(context, userID)
	if errCheckExistence != nil {
		return &domain.CustomError{Message: errCheckExistence.Error(), Code: http.StatusInternalServerError}
	}

	if existenceCount == 0 {
		return &domain.CustomError{Message: "user by the given user id doesn't exist", Code: http.StatusBadRequest}
	}

	if user["username"] != nil {
		countUsername, errCountUsername := userUC.userRepo.CountByUsername(context, user["username"].(string))
		if errCountUsername != nil {
			return &domain.CustomError{Message: "error counting by username" + errCountUsername.Error(), Code: http.StatusInternalServerError}
		}
		if countUsername > 0 {
			return &domain.CustomError{Message: "Username already taken", Code: http.StatusBadRequest}
		} else {
			allowedFeilds["username"] = user["username"]
		}
	}

	if user["email"] != nil {
		countEmail, errCountEmail := userUC.userRepo.CountByEmail(context, user["email"].(string))
		if errCountEmail != nil {
			return &domain.CustomError{Message: "error counting by email" + errCountEmail.Error(), Code: http.StatusInternalServerError}
		}

		if countEmail > 0 {
			return &domain.CustomError{Message: "Email already taken", Code: http.StatusBadRequest}
		} else {
			allowedFeilds["email"] = user["email"]
		}
	}

	if user["bio"] != nil {
		allowedFeilds["bio"] = strings.Trim(user["bio"].(string), " ")
	} else {
		allowedFeilds["bio"] = ""
	}
	allowedFeilds["updatedAt"] = time.Now()
	return userUC.userRepo.UpdateProfile(context, userID, user)
}

func (userUC *userUseCase) ImageUpload(cxt *gin.Context, file *multipart.File, header *multipart.FileHeader, id string) domain.Error {
	timeout, errTimeout := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 0)
	if errTimeout != nil {
		return &domain.CustomError{Message: errTimeout.Error(), Code: http.StatusInternalServerError}
	}
	context, cancel := context.WithTimeout(cxt, time.Duration(timeout)*time.Second)
	defer cancel()

	if valid := utils.IsValidFileFormat(header, "image/png", "image/jpeg"); !valid {
		cxt.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid file format"})
	}

	existingUser, err := userUC.userRepo.FindById(context, id)
	if err != nil {
		return &domain.CustomError{Message: err.Error(), Code: http.StatusInternalServerError}
	}

	if existingUser.ProfilePictureUrl.FilePath != "" {
		errDelete := utils.DeleteImage(existingUser.ProfilePictureUrl.Public_id, context)
		if errDelete != nil {
			return &domain.CustomError{Message: errDelete.Error(), Code: http.StatusInternalServerError}
		}
	}

	upload_reslt, errSave := utils.SaveImage(*file, header.Filename, context)
	if errSave != nil {
		return &domain.CustomError{Message: errSave.Error(), Code: http.StatusInternalServerError}
	}

	photo := domain.Photo{
		Filename:   header.Filename,
		FilePath:   upload_reslt.SecureURL,
		Public_id:  upload_reslt.PublicID,
		UploadedAt: time.Now(),
	}

	errUpdate := userUC.userRepo.UploadProfilePicture(context, photo, id)
	if errUpdate != nil {
		return &domain.CustomError{Message: errUpdate.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

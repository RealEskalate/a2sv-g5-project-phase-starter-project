package usecase

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

/* Defines a struct with all the necessary data to implement domain.UserUsecaseInterface */
type UserUsecase struct {
	userRepository  domain.UserRepositoryInterface
	cacheRepository domain.CacheRepositoryInterface
	MailService     domain.MailServiceInterface
	JWTService      domain.JWTServiceInterface
	HashingService  domain.HashingServiceInterface
	GenerateToken   func(int) (string, error)
	VerifyIdToken   func(string, string, string) error
	DeleteFile      func(string) error
	ENV             domain.EnvironmentVariables
}

/* Regex for validation phone numbers*/
var PhoneRegex = regexp.MustCompile(`^\+?[1-9][0-9]{7,14}$`)

/* Creates a new instance of UserUsecase */
func NewUserUsecase(
	userRepository domain.UserRepositoryInterface,
	cacheRepository domain.CacheRepositoryInterface,
	GenerateToken func(int) (string, error),
	MailService domain.MailServiceInterface,
	JWTService domain.JWTServiceInterface,
	HashingService domain.HashingServiceInterface,
	VerifyIdToken func(string, string, string) error,
	DeleteFile func(string) error,
	ENV domain.EnvironmentVariables) *UserUsecase {
	return &UserUsecase{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
		GenerateToken:   GenerateToken,
		MailService:     MailService,
		JWTService:      JWTService,
		HashingService:  HashingService,
		VerifyIdToken:   VerifyIdToken,
		DeleteFile:      DeleteFile,
		ENV:             ENV,
	}
}

/* Validates password length constraints*/
func (u *UserUsecase) ValidatePassword(password string) domain.CodedError {
	// validate length
	if len(password) < 8 {
		return domain.NewError("Password too short", domain.ERR_BAD_REQUEST)
	}

	if len(password) > 71 {
		return domain.NewError("Password too long", domain.ERR_BAD_REQUEST)
	}

	// validate lower case letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return domain.NewError("Password must contain a lower case letter", domain.ERR_BAD_REQUEST)
	}

	// validate upper case letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return domain.NewError("Password must contain an upper case letter", domain.ERR_BAD_REQUEST)
	}

	// validate number
	if !strings.ContainsAny(password, "0123456789") {
		return domain.NewError("Password must contain a number", domain.ERR_BAD_REQUEST)
	}

	// validate special character
	if !strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?/\\") {
		return domain.NewError("Password must contain a special character", domain.ERR_BAD_REQUEST)
	}

	return nil
}

/* Validates username content and length constraints*/
func (u *UserUsecase) ValidateUsername(username string) domain.CodedError {
	if len(username) < 3 {
		return domain.NewError("Username too short", domain.ERR_BAD_REQUEST)
	}

	if len(username) > 20 {
		return domain.NewError("Username too long", domain.ERR_BAD_REQUEST)
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !re.MatchString(username) {
		return domain.NewError("Invalid Username: must contain only letters, numbers and underscores", domain.ERR_BAD_REQUEST)
	}

	return nil
}

/* Validates email format */
func (u *UserUsecase) ValidateEmail(email string) domain.CodedError {
	if _, err := mail.ParseAddress(email); err != nil {
		return domain.NewError("Invalid Email", domain.ERR_BAD_REQUEST)
	}

	return nil
}

/* Sanitizes user email, username, bio and phonenumber fields */
func (u *UserUsecase) SantizeUserFields(user *domain.User) {
	user.Email = strings.ReplaceAll(strings.TrimSpace(strings.ToLower(user.Email)), " ", "")
	user.Username = strings.ReplaceAll(strings.TrimSpace(strings.ToLower(user.Username)), " ", "")
	user.Bio = strings.TrimSpace(user.Bio)
	user.PhoneNumber = strings.ReplaceAll(strings.TrimSpace(user.PhoneNumber), " ", "")
}

/* Calls sanitization and validation functions and validates bio and phonenumber format */
func (u *UserUsecase) SanitizeAndValidateNewUser(user *domain.User) domain.CodedError {
	u.SantizeUserFields(user)
	user.CreatedAt = time.Now().Round(0)

	err := u.ValidateUsername(user.Username)
	if err != nil {
		return err
	}

	err = u.ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	err = u.ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	if len(user.Bio) == 0 {
		return domain.NewError("Bio can not be empty", domain.ERR_BAD_REQUEST)
	}

	if !PhoneRegex.MatchString(user.PhoneNumber) {
		return domain.NewError("Invalid PhoneNumber: must be informat +XXXXXXXXXX", domain.ERR_BAD_REQUEST)
	}

	return nil
}

/* Generates a verification struct with the provided fields */
func (u *UserUsecase) GetVerificationData(c context.Context, verificationType string, expiresAt time.Time, tokenLength int) (domain.VerificationData, domain.CodedError) {
	var verificationData domain.VerificationData
	generatedToken, gErr := u.GenerateToken(tokenLength)
	if gErr != nil {
		return verificationData, domain.NewError("Internal server error", domain.ERR_INTERNAL_SERVER)
	}

	verificationData = domain.VerificationData{
		Token:     generatedToken,
		ExpiresAt: expiresAt,
		Type:      verificationType,
	}

	return verificationData, nil
}

/*
Creates a new user in the system after sanitizing and validating the user fields. It then sends
an email to the user with a verification link to verify their email address
*/
func (u *UserUsecase) Signup(c context.Context, user *domain.User, hostUrl string) domain.CodedError {
	err := u.SanitizeAndValidateNewUser(user)
	if err != nil {
		return err
	}

	hashedPwd, err := u.HashingService.HashString(user.Password)
	if err != nil {
		return domain.NewError("Internal server error", domain.ERR_INTERNAL_SERVER)
	}

	verificationData, err := u.GetVerificationData(c, domain.VerifyEmailType, time.Now().Round(0).Add(time.Hour*2), 32)
	if err != nil {
		return err
	}

	user.VerificationData = verificationData
	user.Password = hashedPwd
	user.Role = domain.RoleUser
	user.IsVerified = false
	user.CreatedAt = time.Now().Round(0)

	err = u.userRepository.CreateUser(c, user)
	if err != nil {
		return err
	}

	// send email verification link with the template and the generated token
	mail := u.MailService.EmailVerificationTemplate(hostUrl, user.Username, verificationData.Token)
	mailErr := u.MailService.SendMail("Blog API", user.Email, mail)
	if mailErr != nil {
		u.userRepository.DeleteUser(c, user.Username)
		return domain.NewError("Internal server error: "+mailErr.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

/*
Creates a new user in the system from a google oauth request, a valid user name and a password
if the user does not exist in the system.
*/
func (u *UserUsecase) OAuthSignup(c context.Context, data *dtos.GoogleResponse, userCreds *dtos.OAuthSignup) domain.CodedError {
	newUser := &domain.User{
		Username: userCreds.Username,
		Email:    data.RawData.Email,
		Password: userCreds.Password,
		ProfilePicture: dtos.ProfilePicture{
			FileName: data.RawData.Picture,
			IsLocal:  false,
		},
	}

	// verify ID token with the google API
	vErr := u.VerifyIdToken(data.IDToken, data.Email, env.ENV.GOOGLE_CLIENT_ID)
	if vErr != nil {
		return domain.NewError(vErr.Error(), domain.ERR_UNAUTHORIZED)
	}

	u.SantizeUserFields(newUser)
	err := u.ValidateUsername(newUser.Username)
	if err != nil {
		return err
	}

	err = u.ValidateEmail(newUser.Email)
	if err != nil {
		return err
	}

	err = u.ValidatePassword(newUser.Password)
	if err != nil {
		return err
	}

	hashedPwd, err := u.HashingService.HashString(newUser.Password)
	if err != nil {
		return domain.NewError("Internal server error", domain.ERR_INTERNAL_SERVER)
	}

	newUser.Password = hashedPwd
	newUser.Role = domain.RoleUser
	newUser.IsVerified = true
	newUser.CreatedAt = time.Now().Round(0)

	err = u.userRepository.CreateUser(c, newUser)
	if err != nil {
		return err
	}

	u.userRepository.VerifyUser(c, newUser.Username)
	return nil
}

/*
Checks if the provided user has the correct credentials and is verified. If the user is verified,
it signs a new access and refresh token and sets the hashed refresh token in the database.
*/
func (u *UserUsecase) Login(c context.Context, user *domain.User) (string, string, domain.CodedError) {
	u.SantizeUserFields(user)
	err := u.ValidateUsername(user.Username)
	if user.Username != "" && err != nil {
		return "", "", err
	}

	err = u.ValidateEmail(user.Email)
	if user.Email != "" && err != nil {
		return "", "", err
	}

	if user.Email == "" && user.Username == "" {
		return "", "", domain.NewError("Username or email required", domain.ERR_BAD_REQUEST)
	}

	err = u.ValidatePassword(user.Password)
	if err != nil {
		return "", "", err
	}

	foundUser, err := u.userRepository.FindUser(c, user)
	if err != nil {
		return "", "", err
	}

	if !foundUser.IsVerified {
		return "", "", domain.NewError("User email not verified", domain.ERR_UNAUTHORIZED)
	}

	err = u.HashingService.ValidateHashedString(foundUser.Password, user.Password)
	if err != nil {
		return "", "", domain.NewError("Incorrect password", domain.ERR_UNAUTHORIZED)
	}

	// sign the new access and refresh tokens
	accessToken, err := u.JWTService.SignJWTWithPayload(foundUser.Username, foundUser.Role, "accessToken", time.Minute*time.Duration(u.ENV.ACCESS_TOKEN_LIFESPAN))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.JWTService.SignJWTWithPayload(foundUser.Username, foundUser.Role, "refreshToken", time.Hour*time.Duration(u.ENV.REFRESH_TOKEN_LIFESPAN))
	if err != nil {
		return "", "", err
	}

	// set the new refresh token in the database after hashing it
	hashedRefreshToken, err := u.HashingService.HashString(strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	err = u.userRepository.SetRefreshToken(c, user, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

/*
Allows the user to obtain access tokens using their google account. If the user does not exist in
the database, a new user IS NOT created. The user must have an account in the system to be able
to obtain tokens using this route.
*/
func (u *UserUsecase) OAuthLogin(c context.Context, data *dtos.GoogleResponse) (string, string, domain.CodedError) {
	foundUser, err := u.userRepository.FindUser(c, &domain.User{Email: data.Email})
	if err != nil && err.GetCode() == domain.ERR_NOT_FOUND {
		return "", "", err
	}

	// verify ID token with the google API
	vErr := u.VerifyIdToken(data.IDToken, data.Email, env.ENV.GOOGLE_CLIENT_ID)
	if vErr != nil {
		return "", "", domain.NewError(vErr.Error(), domain.ERR_UNAUTHORIZED)
	}

	if err != nil {
		return "", "", err
	}

	if !foundUser.IsVerified {
		return "", "", domain.NewError("User email not verified", domain.ERR_UNAUTHORIZED)
	}

	// signs the new access and refresh tokens
	accessToken, err := u.JWTService.SignJWTWithPayload(foundUser.Username, foundUser.Role, "accessToken", time.Minute*time.Duration(u.ENV.ACCESS_TOKEN_LIFESPAN))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.JWTService.SignJWTWithPayload(foundUser.Username, foundUser.Role, "refreshToken", time.Hour*time.Duration(u.ENV.REFRESH_TOKEN_LIFESPAN))
	if err != nil {
		return "", "", err
	}

	// set the new refresh token in the database after hashing it
	hashedRefreshToken, err := u.HashingService.HashString(strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	err = u.userRepository.SetRefreshToken(c, &foundUser, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

/*
Checks if the provided refresh token is valid and not expired. If the token is valid, it signs a new access token.
If the token is invalid, it returns an error. If the token is expired, it deletes the refresh token from the database
*/
func (u *UserUsecase) RenewAccessToken(c context.Context, refreshToken string) (string, domain.CodedError) {
	token, err := u.JWTService.ValidateAndParseToken(refreshToken)
	if err != nil {
		return "", domain.NewError("Invalid token", domain.ERR_UNAUTHORIZED)
	}

	// check whether the token is a refreshToken
	tokenType, err := u.JWTService.GetTokenType(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	if tokenType != "refreshToken" {
		return "", domain.NewError("Invalid token type: make sure to use the refreshToken", domain.ERR_FORBIDDEN)
	}

	// get the username from the token
	username, err := u.JWTService.GetUsername(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	// check expiry date of the refresh token
	expiresAtTime, err := u.JWTService.GetExpiryDate(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	if expiresAtTime.Compare(time.Now()) == -1 {
		u.userRepository.SetRefreshToken(c, &domain.User{Username: username}, "")
		return "", domain.NewError("Token expired", domain.ERR_UNAUTHORIZED)
	}

	// get the hashed refresh token from the database
	foundUser, qErr := u.userRepository.FindUser(c, &domain.User{Username: username})
	if qErr != nil {
		return "", qErr
	}

	// check if the user is verified
	if !foundUser.IsVerified {
		return "", domain.NewError("User email not verified", domain.ERR_UNAUTHORIZED)
	}

	if foundUser.RefreshToken == "" {
		return "", domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	err = u.HashingService.ValidateHashedString(foundUser.RefreshToken, strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	accessToken, err := u.JWTService.SignJWTWithPayload(foundUser.Username, foundUser.Role, "accessToken", time.Minute*time.Duration(u.ENV.ACCESS_TOKEN_LIFESPAN))
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return accessToken, nil
}

/*
Updates the user details if the user is the owner of the account.
The user can only update their bio and phonenumber.
*/
func (u *UserUsecase) UpdateUser(c context.Context, requestUsername string, tokenUsername string, user *dtos.UpdateUser) (map[string]string, domain.CodedError) {
	if requestUsername != tokenUsername {
		return nil, domain.NewError("Only the owner of the account can update its details", domain.ERR_FORBIDDEN)
	}

	if user.PhoneNumber != "" && !PhoneRegex.MatchString(user.PhoneNumber) {
		return nil, domain.NewError("Invalid phone number: must be informat +XXXXXXXXXX", domain.ERR_BAD_REQUEST)
	}

	if user.Bio == "" && user.PhoneNumber == "" && user.ProfilePicture.FileName == "" {
		return nil, domain.NewError("No fields to update", domain.ERR_BAD_REQUEST)
	}

	res, oldPicture, err := u.userRepository.UpdateUser(c, requestUsername, user)
	if oldPicture != "" {
		u.DeleteFile("./local/" + oldPicture)
	}

	return res, err
}

/* Promotes the user with the provided username to the `admin` role` */
func (u *UserUsecase) PromoteUser(c context.Context, username string) domain.CodedError {
	return u.userRepository.ChangeRole(c, username, domain.RoleAdmin)
}

/* Demotes the user with the provided username to the `user` role */
func (u *UserUsecase) DemoteUser(c context.Context, username string) domain.CodedError {
	return u.userRepository.ChangeRole(c, username, domain.RoleUser)
}

/*
Verifies the user email address by checking the provided token against the token in the database.
  - If the token is valid, it sets the user as verified.
  - If the token is invalid, it returns an error.
  - If the token is expired, it generates a new token and sends a new email to the user.
*/
func (u *UserUsecase) VerifyEmail(c context.Context, username string, token string, hostUrl string) domain.CodedError {
	username = strings.TrimSpace(username)
	user, err := u.userRepository.FindUser(c, &domain.User{Username: username})
	if err != nil {
		return err
	}

	if user.VerificationData.Token != token {
		return domain.NewError("Invalid token", domain.ERR_BAD_REQUEST)
	}

	if user.IsVerified {
		return domain.NewError("User already verified", domain.ERR_BAD_REQUEST)
	}

	if user.VerificationData.ExpiresAt.Before(time.Now().Round(0)) {
		verificationData, err := u.GetVerificationData(c, domain.VerifyEmailType, time.Now().Round(0).Add(time.Hour*2), 32)
		if err != nil {
			return err
		}

		err = u.userRepository.UpdateVerificationDetails(c, username, verificationData)
		if err != nil {
			return err
		}

		mail := u.MailService.EmailVerificationTemplate(hostUrl, username, verificationData.Token)
		mailErr := u.MailService.SendMail("Blog API", user.Email, mail)
		if mailErr != nil {
			return domain.NewError("Internal server error: "+mailErr.Error(), domain.ERR_INTERNAL_SERVER)
		}

		return domain.NewError("Token expired. Another link has been sent to your email. Please try again", domain.ERR_BAD_REQUEST)
	}

	return u.userRepository.VerifyUser(c, username)
}

/*
Starts the process of resetting the user password by sending an email with a reset password link
and a token to the user. The token is stored in the database and is used to verify the user when
they want to reset their password.
*/
func (u *UserUsecase) InitResetPassword(c context.Context, username string, email string, hostUrl string) domain.CodedError {
	foundUser, err := u.userRepository.FindUser(c, &domain.User{Username: username, Email: email})
	if err != nil {
		return err
	}

	if !foundUser.IsVerified {
		return domain.NewError("User email not verified", domain.ERR_UNAUTHORIZED)
	}

	verificationData, err := u.GetVerificationData(c, domain.ResetPasswordType, time.Now().Round(0).Add(time.Minute*10), 12)
	if err != nil {
		return err
	}

	err = u.userRepository.UpdateVerificationDetails(c, username, verificationData)
	if err != nil {
		return err
	}

	mail := u.MailService.PasswordResetTemplate(verificationData.Token)
	mailErr := u.MailService.SendMail("Blog API", foundUser.Email, mail)
	if mailErr != nil {
		return domain.NewError("Internal server error: "+mailErr.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

/*
Resets the user password by checking the provided token against the token in the database.
  - If the token is valid, it sets the new password for the user.
  - If the token is invalid, it returns an error.
*/
func (u *UserUsecase) ResetPassword(c context.Context, resetDto dtos.ResetPassword, token string) domain.CodedError {
	user, err := u.userRepository.FindUser(c, &domain.User{Username: resetDto.Username})
	if err != nil {
		return err
	}

	if user.VerificationData.Token != token {
		return domain.NewError("Invalid token", domain.ERR_UNAUTHORIZED)
	}

	if user.VerificationData.Type != domain.ResetPasswordType {
		return domain.NewError("Invalid verification type", domain.ERR_UNAUTHORIZED)
	}

	if user.VerificationData.ExpiresAt.Before(time.Now().Round(0)) {
		return domain.NewError("Token expired", domain.ERR_UNAUTHORIZED)
	}

	err = u.ValidatePassword(resetDto.NewPassword)
	if err != nil {
		return err
	}

	hashedPwd, err := u.HashingService.HashString(resetDto.NewPassword)
	if err != nil {
		return domain.NewError("Internal server error", domain.ERR_INTERNAL_SERVER)
	}

	err = u.userRepository.UpdatePassword(c, resetDto.Username, hashedPwd)
	if err != nil {
		return err
	}

	return nil
}

/*
Logs out the user by deleting the refresh token from the database and setting the access token in the
blacklist cache for a duration equal to the access token lifespan set in the environment variables.
*/
func (u *UserUsecase) Logout(c context.Context, username string, accessToken string) domain.CodedError {
	err := u.cacheRepository.CacheData(accessToken, "", time.Minute*time.Duration(u.ENV.ACCESS_TOKEN_LIFESPAN))
	if err != nil {
		return err
	}

	return u.userRepository.SetRefreshToken(c, &domain.User{Username: username}, "")
}

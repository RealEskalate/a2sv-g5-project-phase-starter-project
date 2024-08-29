package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Repository"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	Register(input Domain.RegisterInput) (*Domain.User, error)
	UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error
	DeleteUser(username string) error
	Login(c *gin.Context, LoginUser *Domain.LoginInput) (string, string, error)
	Logout(tokenString string) error
	ForgotPassword(c *gin.Context, username string) (string, error)
	Reset(c *gin.Context, token string) (string, error)
	UpdatePassword(username string, newPassword string) error
	PromoteTOAdmin(username string) (Domain.User, error)
	Verify(token string) error
	OAuthLogin(c *gin.Context, code string) (*Domain.User, string, error)
	FindUser() ([]Domain.User, error)
	GiveId(username string) (string, error)
	InsertToken(username string, accessToken string, refreshToken string) error
	CleanUpExpiredTokens(ctx context.Context) error
}

type userUsecase struct {
	userRepo        Repository.UserRepository
	emailService    *infrastructure.EmailService
	passwordService *infrastructure.PasswordService
}

func NewUserUsecase(userRepo Repository.UserRepository, emailService *infrastructure.EmailService) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		emailService:    emailService,
		passwordService: infrastructure.NewPasswordService(),
	}
}

const (
	passwordMinLength = 8
	passwordMaxLength = 20
)

func (u *userUsecase) InsertToken(username string, accessToken string, refreshToken string) error {
	err := u.userRepo.InsertToken(username, accessToken, refreshToken)
	if err != nil {
		return err

	}
	return nil
}

func (u *userUsecase) GiveId(username string) (string, error) {
	fmt.Println("===============uffffffffffffffff")

	id, err := u.userRepo.GetIDBYUsername(username)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Register creates a new user, hashes the password, generates a verification token, and sends a welcome email.
func (u *userUsecase) Register(input Domain.RegisterInput) (*Domain.User, error) {
	// Validate username
	if strings.Contains(input.Username, "@") {
		return nil, errors.New("username must not contain '@'")
	}

	// Check if username already exists
	if _, err := u.userRepo.FindByUsername(input.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	// Validate email format
	if !isValidEmail(input.Email) {
		return nil, errors.New("invalid email format")
	}

	// Check if email already registered
	if _, err := u.userRepo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	// Validate password strength
	if err := validatePasswordStrength(input.Password); err != nil {
		return nil, err
	}

	// Hash the password
	hashedPassword, err := u.passwordService.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create new user
	user := &Domain.User{
		Id:             primitive.NewObjectID(),
		Name:           input.Name,
		Username:       input.Username,
		Email:          input.Email,
		Password:       string(hashedPassword),
		ProfilePicture: input.ProfilePicture,
		Bio:            input.Bio,
		Gender:         input.Gender,
		Address:        input.Address,
		IsActive:       false, // Initially inactive
		PostsIDs:       []string{},
	}

	// Set user role based on database state
	if ok, err := u.userRepo.IsDbEmpty(); ok && err == nil {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// Save user to repository
	err = u.userRepo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	// Generate a verification token
	newToken, err := infrastructure.GenerateResetToken(user.Username, user.Role, []byte("BlogManagerSecretKey"))
	if err != nil {
		return nil, fmt.Errorf("failed to generate verification token: %v", err)
	}

	// Construct the email body
	subject := "Welcome to Our Service!"
	body := fmt.Sprintf("Hi %s,\n\nWelcome to our platform! Please verify your account by clicking the link below:\n\nhttp://localhost:8080/verify/%s\n\nThank you!", input.Name, newToken)

	// Send verification email
	err = u.emailService.SendEmail(input.Email, subject, body)
	if err != nil {
		return nil, fmt.Errorf("failed to send welcome email: %v", err)
	}

	return user, nil
}

func (u *userUsecase) UpdateUser(username string, updatedUser *Domain.UpdateUserInput) error {
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return errors.New("user not found")
	}

	updateFields := bson.M{}

	if updatedUser.Username != "" {
		if strings.Contains(updatedUser.Username, "@") {
			return errors.New("username must not contain '@'")
		}
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Password != "" {
		hashedPassword, err := u.passwordService.HashPassword(updatedUser.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %v", err)
		}
		updateFields["password"] = hashedPassword
	}
	if updatedUser.ProfilePicture != "" {
		updateFields["profile_picture"] = updatedUser.ProfilePicture
	}
	if updatedUser.Bio != "" {
		updateFields["bio"] = updatedUser.Bio
	}
	if updatedUser.Address != "" {
		updateFields["address"] = updatedUser.Address
	}

	err = u.userRepo.Update(username, updateFields)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (u *userUsecase) DeleteUser(username string) error {
	_, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	err = u.userRepo.Delete(username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (u *userUsecase) Login(c *gin.Context, LoginUser *Domain.LoginInput) (string, string, error) {
	user, err := u.userRepo.FindByUsername(LoginUser.Username)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	err = u.passwordService.ComparePasswords(user.Password, LoginUser.Password)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	accessToken, err := infrastructure.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(user.Username)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %v", err)
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	err = u.userRepo.InsertToken(user.Username, accessToken, refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to store tokens: %v", err)
	}

	if !user.IsActive {
		return "", "", fmt.Errorf("user not verified")
	}

	return accessToken, refreshToken, nil
}

func (u *userUsecase) Logout(tokenString string) error {
	err := u.userRepo.ExpireToken(tokenString)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) ForgotPassword(c *gin.Context, username string) (string, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	accessToken, err := infrastructure.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %v", err)
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	err = u.userRepo.InsertToken(user.Username, accessToken, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to store tokens: %v", err)
	}

	subject := "Password Reset Request"
	body := fmt.Sprintf(`
	Hi %s,

	It seems like you requested a password reset. No worries, it happens to the best of us! You can reset your password by clicking the link below:

	<a href="http://localhost:8080/reset/%s">Reset Your Password</a>

	If you did not request a password reset, please ignore this email.

Best regards,
	Your Support Team
	`, user.Name, accessToken)

	err = u.emailService.SendEmail(user.Email, subject, body)
	if err != nil {
		return "", fmt.Errorf("failed to send reset email: %v", err)
	}

	return accessToken, nil
}

func (u *userUsecase) Reset(c *gin.Context, token string) (string, error) {

	claims, err := infrastructure.ParseResetToken(token, []byte("BlogManagerSecretKey"))
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return "", err
	}

	user, err := u.userRepo.FindByUsername(claims.Username)

	if err != nil {
		return "", errors.New("user not found")
	}

	accessToken, err := infrastructure.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %v", err)
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	err = u.userRepo.InsertToken(user.Username, accessToken, refreshToken)
	if err != nil {
		return "", errors.New("user not found")
	}
	access_token, err := infrastructure.GenerateJWT(user.Username, user.Role)

	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	return access_token, nil
}

func (u *userUsecase) UpdatePassword(username string, newPassword string) error {

	hashedPassword, err := u.passwordService.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	err = u.userRepo.Update(username, bson.M{"password": hashedPassword})
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

func (u *userUsecase) PromoteTOAdmin(ID string) (Domain.User, error) {
	user, err := u.userRepo.FindByID(ID)
	if err != nil {
		return Domain.User{}, errors.New("user not found")
	}
	username := user.Username
	if user.Role == "user" {
		err = u.userRepo.Update(username, bson.M{"role": "admin"})
		if err != nil {
			return Domain.User{}, fmt.Errorf("failed to promote user to admin: %v", err)
		}
	} else {
		err = u.userRepo.Update(username, bson.M{"role": "user"})
		if err != nil {
			return Domain.User{}, fmt.Errorf("failed to demote user to user: %v", err)
		}
	}

	return user, nil
}

func (u *userUsecase) Verify(token string) error {
	claims, err := infrastructure.ParseResetToken(token, []byte("BlogManagerSecretKey"))
	if err != nil {
		fmt.Println("Error parsing token:", err)
	}

	user, err := u.userRepo.FindByUsername(claims.Username)
	if err != nil {
		return errors.New("user not found")
	}
	err = u.userRepo.Update(user.Username, bson.M{"is_active": true})
	if err != nil {
		return fmt.Errorf("failed to verify user: %v", err)
	}
	return nil
}

// isValidEmail checks if the email format is valid
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// validatePasswordStrength checks if the password meets strength criteria
func validatePasswordStrength(password string) error {
	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return fmt.Errorf("password must be between %d and %d characters", passwordMinLength, passwordMaxLength)
	}

	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, c := range password {
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= '0' && c <= '9':
			hasDigit = true
		case c == '@' || c == '#' || c == '$' || c == '%' || c == '^' || c == '&' || c == '*':
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (u *userUsecase) OAuthLogin(c *gin.Context, code string) (*Domain.User, string, error) {
	// Exchange the authorization code for an access token
	token, err := infrastructure.OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", errors.New("failed to exchange token: " + err.Error())
	}

	// Retrieve user info using the access token
	client := infrastructure.OAuthConfig.Client(context.Background(), token)
	userInfo, err := fetchUserInfo(client)
	if err != nil {
		return nil, "", errors.New("failed to fetch user info")
	}

	// Prepare the RegisterInput from the OAuth user info
	registerInput := Domain.RegisterInput{
		Name:           userInfo["name"].(string),
		Username:       strings.Split(userInfo["email"].(string), "@")[0],
		Email:          userInfo["email"].(string),
		ProfilePicture: userInfo["picture"].(string),
		IsOauth:        true,
	}

	// Check if the user already exists by email
	existingUser, _ := u.userRepo.FindByEmail(registerInput.Email)
	if existingUser != nil {

		accessToken, err := infrastructure.GenerateJWT(existingUser.Username, existingUser.Role)
		if err != nil {
			return nil, "", err
		}

		refreshToken, err := infrastructure.GenerateRefreshToken(existingUser.Username)
		if err != nil {
			return nil, "", err
		}

		c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

		err = u.userRepo.InsertToken(existingUser.Username, accessToken, refreshToken)

		if err != nil {
			return nil, "", err
		}
		// If the user already exists, generate a token for the existing user
		tokenString, err := infrastructure.GenerateJWT(existingUser.Username, existingUser.Role)
		if err != nil {
			return nil, "", err
		}
		return existingUser, tokenString, nil
	}

	// If the user does not exist, register a new user
	newUser := &Domain.User{
		Id:             primitive.NewObjectID(),
		Name:           registerInput.Name,
		Username:       registerInput.Username,
		Email:          registerInput.Email,
		ProfilePicture: registerInput.ProfilePicture,
		IsOauth:        registerInput.IsOauth,
		IsActive:       true,
	}

	err = u.userRepo.Save(newUser)
	if err != nil {
		return nil, "", err
	}

	// Generate a token for the new user
	accessToken, err := infrastructure.GenerateJWT(newUser.Username, newUser.Role)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := infrastructure.GenerateRefreshToken(newUser.Username)
	if err != nil {
		return nil, "", err
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	err = u.userRepo.InsertToken(newUser.Username, accessToken, refreshToken)

	if err != nil {
		return nil, "", err
	}
	// If the user already exists, generate a token for the existing user
	tokenString, err := infrastructure.GenerateJWT(newUser.Username, newUser.Role)
	if err != nil {
		return nil, "", err
	}
	return newUser, tokenString, nil
}

// fetchUserInfo retrieves the user's information from the OAuth provider
func fetchUserInfo(client *http.Client) (map[string]interface{}, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (u *userUsecase) FindUser() ([]Domain.User, error) {
	users, err := u.userRepo.ShowUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) CleanUpExpiredTokens(ctx context.Context) error {
	return u.userRepo.RemoveExpiredTokens(ctx)
}

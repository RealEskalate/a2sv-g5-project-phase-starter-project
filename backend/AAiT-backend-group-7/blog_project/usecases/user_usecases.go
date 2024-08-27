package usecases

import (
	"blog_project/domain"
	"blog_project/infrastructure"
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync/atomic"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserUsecase struct {
	emailService *infrastructure.EmailService
	UserRepo     domain.IUserRepository
	BlogRepo     domain.IBlogRepository
	TokenRepo    domain.ITokenRepository
}

func NewUserUsecase(
	userRepo domain.IUserRepository,
	blogRepo domain.IBlogRepository,
	emailService *infrastructure.EmailService,
	tokenRepo domain.ITokenRepository,
) domain.IUserUsecase {
	return &UserUsecase{
		emailService: emailService,
		UserRepo:     userRepo,
		BlogRepo:     blogRepo,
		TokenRepo:    tokenRepo,
	}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.UserRepo.GetAllUsers(ctx)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if existingUser, _ := u.UserRepo.SearchByEmail(ctx, user.Email); existingUser.ID != 0 {
		return domain.User{}, errors.New("email already in use")
	}

	if existingUser, _ := u.UserRepo.SearchByUsername(ctx, user.Username); existingUser.ID != 0 {
		return domain.User{}, errors.New("username already in use")
	}

	if !isValidEmail(user.Email) {
		return domain.User{}, errors.New("invalid email")
	}

	if !isValidPassword(user.Password) {
		return domain.User{}, errors.New("invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and a minimum length of 8")
	}

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = hashedPassword
	user.ID = generateUniqueID()

	if users, err := u.GetAllUsers(ctx); err != nil {
		return domain.User{}, err
	} else if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	return u.UserRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	claims, ok := ctx.Value("user").(jwt.MapClaims)
	if !ok {
		return domain.User{}, errors.New("failed to get user claims from context")
	}

	userID := int(claims["id"].(float64))
	userRole := claims["role"].(string)

	if userRole != "admin" && userID != id {
		return domain.User{}, errors.New("you are not authorized to update this user")
	}

	existingUser, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	if user.Email != "" {
		if !isValidEmail(user.Email) {
			return domain.User{}, errors.New("invalid email")
		}
		existingUser.Email = user.Email
	}

	if user.Password != "" {
		if !isValidPassword(user.Password) {
			return domain.User{}, errors.New("invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and a minimum length of 8")
		}

		hashedPassword, err := infrastructure.HashPassword(user.Password)
		if err != nil {
			return domain.User{}, err
		}
		existingUser.Password = hashedPassword
	}

	if user.Username != "" {
		u.BlogRepo.UpdateAuthorName(ctx, existingUser.Username, user.Username)
		existingUser.Username = user.Username
	}

	if user.Bio != "" {
		existingUser.Bio = user.Bio
	}

	if user.Phone != "" {
		existingUser.Phone = user.Phone
	}

	if user.ProfilePic != "" {
		existingUser.ProfilePic = user.ProfilePic
	}

	if user.Role != "" {
		existingUser.Role = user.Role
	}

	return u.UserRepo.UpdateUser(ctx, id, existingUser)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	claims, ok := ctx.Value("user").(jwt.MapClaims)
	if !ok {
		return errors.New("failed to get user claims from context")
	}

	userID := int(claims["id"].(float64))
	userRole := claims["role"].(string)

	if userRole != "admin" && userID != id {
		return errors.New("you are not authorized to delete this user")
	}

	return u.UserRepo.DeleteUser(ctx, id)
}

func (u *UserUsecase) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	return u.UserRepo.AddBlog(ctx, userID, blog)
}

func (u *UserUsecase) DeleteBlog(ctx context.Context, userID, blogID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	for i, b := range user.Blogs {
		if b == blogID {
			user.Blogs = append(user.Blogs[:i], user.Blogs[i+1:]...)
			break
		}
	}

	return u.UserRepo.UpdateUser(ctx, userID, user)
}

func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil || user.ID == 0 {
		return "", "", errors.New("invalid credentials")
	}

	if err := infrastructure.ComparePassword(user.Password, password); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	token, err := infrastructure.GenerateJWTAccessToken(&user, os.Getenv("JWT_SECRET"), 1)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := infrastructure.GenerateJWTRefreshToken(&user, os.Getenv("JWT_SECRET"), 5)
	if err != nil {
		return "", "", err
	}

	if err := u.UserRepo.StoreRefreshToken(ctx, user.ID, refreshToken); err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *UserUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	validatedToken, err := infrastructure.IsAuthorized(refreshToken, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", errors.New("invalid token")
	}

	userID, ok := validatedToken["id"].(float64)
	if !ok {
		return "", errors.New("invalid token ID type")
	}

	user, err := u.UserRepo.GetUserByID(ctx, int(userID))
	if err != nil {
		return "", errors.New("user not found")
	}

	newToken, err := infrastructure.GenerateJWTAccessToken(&user, os.Getenv("JWT_SECRET"), 1)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (u *UserUsecase) ForgetPassword(ctx context.Context, email string) error {
	user, err := u.UserRepo.SearchByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}

	resetToken, err := infrastructure.GenerateJWTRefreshToken(&user, os.Getenv("JWT_SECRET"), 1)
	if err != nil {
		return err
	}

	if err := u.emailService.SendPasswordResetEmail(email, resetToken); err != nil {
		return errors.New("failed to send password reset email")
	}

	return nil
}

func (u *UserUsecase) ResetPassword(ctx context.Context, token, newPassword string) error {
	claims, err := infrastructure.IsAuthorized(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return errors.New("invalid token")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return errors.New("invalid token ID type")
	}

	user, err := u.UserRepo.GetUserByID(ctx, int(userID))
	if err != nil {
		return errors.New("user not found")
	}

	hashedPassword, err := infrastructure.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = hashedPassword

	if _, err := u.UserRepo.UpdateUser(ctx, user.ID, user); err != nil {
		return errors.New("failed to update password")
	}

	return nil
}

func (u *UserUsecase) PromoteUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	user.Role = "admin"

	return u.UpdateUser(ctx, user.ID, user)
}

func (u *UserUsecase) DemoteUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	user.Role = "user"

	return u.UpdateUser(ctx, user.ID, user)
}

func (u *UserUsecase) Logout(ctx context.Context, token string) error {
	decodedToken, err := infrastructure.IsAuthorized(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	if err := u.TokenRepo.BlacklistToken(ctx, token); err != nil {
		return err
	}

	userID := int(decodedToken["id"].(float64))

	refreshToken, err := u.UserRepo.GetRefreshToken(ctx, userID)
	if err != nil {
		return err
	}

	return u.TokenRepo.BlacklistToken(ctx, refreshToken)
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}

var counter int32

func generateUniqueID() int {
	timestamp := int(time.Now().UnixNano() / 1e6 % 1e6)
	uniqueID := timestamp*1000 + int(atomic.AddInt32(&counter, 1)%1000)

	if uniqueID > 2147483647 {
		uniqueID %= 1000000
	}

	return uniqueID
}

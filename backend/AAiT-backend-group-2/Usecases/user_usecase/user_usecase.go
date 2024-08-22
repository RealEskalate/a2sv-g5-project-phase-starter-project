package user_usecase

import (
	domain "AAiT-backend-group-2/Domain"
	"AAiT-backend-group-2/Infrastructure/dtos"
	"AAiT-backend-group-2/Infrastructure/services"
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"
)


type userUsecase struct {
	userRepository domain.UserRepository
	jwtService services.JWTService
	emailService *services.EmailService
	imageService *services.ImageService
	contextTimeout time.Duration
	validator *services.ValidatorService
}

func NewUserUsecase(userRepository domain.UserRepository, jwtService services.JWTService, emailService *services.EmailService, imageService *services.ImageService, timeout time.Duration, validator *services.ValidatorService) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		jwtService: jwtService,
		emailService: emailService,
		imageService: imageService,
		contextTimeout: timeout,
		validator: validator,
	}
}

func (uu *userUsecase) GetAllUsers(c context.Context,) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.FindAll(ctx)
}

func (uu *userUsecase) GetUserByID(c context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.FindByID(ctx, id)
}

func (uu *userUsecase) CreateUser(c context.Context, user domain.User) error {
	if err := services.ValidateStruct(uu.validator, user); err != nil {
		return fmt.Errorf("validation error: %v", err.Error())
	}

	_, err := uu.userRepository.FindByEmailOrUsername(c, user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	_, err = uu.userRepository.FindByEmailOrUsername(c, user.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := services.GeneratePasswordHash(user.Password)
	if err != nil {
		return errors.New("internal server error: failed to hash password")
	}

	user.Password = hashedPassword

	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	count, err := uu.userRepository.CountDocuments(ctx)

	if err != nil {
		return errors.New("internal server error: failed to count users")
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()

	return uu.userRepository.Save(ctx, user)
}

func (uu *userUsecase) UpdateUser(c context.Context, id string, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"username": user.Username,
		"update_at": time.Now(),
		"profile": user.Profile,
	}

	return uu.userRepository.Update(ctx, id, updateData)
}

func (uu *userUsecase) DeleteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Delete(ctx, id)
}

func (uu *userUsecase) Login(c context.Context, loginDto *dtos.LoginDTO) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.FindByEmailOrUsername(ctx, loginDto.EmailOrUserName)

	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	if !services.ComparePasswordHash(loginDto.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	tokens, err := uu.jwtService.GenerateToken(user.ID.Hex(), user.Email, user.Role, 15*time.Minute, 24*time.Hour)
	if err != nil {
		return nil, errors.New("internal server error: failed to generate token")
	}

	updateFields := map[string]interface{}{
		"refresh_token": tokens["refresh_token"],
	}
	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateFields); err != nil {
		return nil, errors.New("internal server error")
	}

	return tokens, nil
}

func (uu *userUsecase) RefreshToken(c context.Context, refreshToken string) (string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	claims, err := uu.jwtService.ValidateToken(refreshToken, "refresh_token")
	if err != nil {
		return "", err
	}

	user, err := uu.userRepository.FindByID(ctx, claims["userID"])
	if err != nil {
		return "", errors.New("internal server error: user not found")
	}

	if user.RefreshToken != refreshToken {
		return "", errors.New("invalid token")
	}

	newAcessToken, err := uu.jwtService.RenewToken(claims)

	if err != nil {
		return "", errors.New("internal server error: failed to renew token")
	}

	return newAcessToken, nil
}

func (uu *userUsecase) PromoteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"role": "admin",
	}

	return uu.userRepository.PromoteUser(ctx, id, updateData)
}

func (uu *userUsecase) DemoteAdmin(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"role": "admin",
	}

	return uu.userRepository.DemoteAdmin(ctx, id, updateData)
}

func (uu *userUsecase) Logout(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"refresh_token": "",
	}

	return uu.userRepository.Update(ctx, id, updateData)
}


func (uu *userUsecase) ForgotPassword(c context.Context, userId, email string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.FindByEmailOrUsername(c, email)
	if err != nil {
		return errors.New("internal server error: user not found")
	}

	if user != nil && user.ID.Hex() != userId {
		return errors.New("unatuhorized")
	}

	token, err := services.GenerateSecureToken(32)
	hashedToken := services.HashToken(token)
	if err != nil {
		return errors.New("internal server error: failed to generate token")
	}

	_, err = uu.userRepository.ForgotPassword(ctx, user.Email, hashedToken)

	if err != nil {
		return errors.New("internal server error")
	}

	resetModel, err := uu.emailService.GeneratePasswordResetTemplate(email, user.Username, token)
	if err != nil {
		return errors.New("unable to reset password")
	}

	if err := uu.emailService.SendEmail(email, "Password Reset", "reset_email.html", resetModel); err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) ResetPassword(c context.Context,  userId string, passwordResetDto *dtos.PasswordResetDto) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	if passwordResetDto.Password != passwordResetDto.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	user, err := uu.userRepository.FindByEmailOrUsername(ctx, passwordResetDto.Email)
	if err != nil {
		return errors.New("internal server error: user not found")
	}

	if user == nil {
		return errors.New("user not found")
	}

	if user.ID.Hex() != userId {
		return errors.New("unatuhorized")
	}

	if err := services.ValidateStruct(uu.validator, passwordResetDto); err != nil {
		return fmt.Errorf("validation error: %v", err.Error())
	}

	hashedToken := services.HashToken(passwordResetDto.Token)

	if err := uu.userRepository.ValidateResetToken(c, user.ID.Hex(), hashedToken); err != nil {
		return err
	}

	hashedPassword, err := services.GeneratePasswordHash(passwordResetDto.Password)
	if err != nil {
		return errors.New("internal server error: failed to hash password")
	}

	updateData := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return errors.New("internal server error")
	}

	if err := uu.userRepository.InvalidateResetToken(c, user.ID.Hex()); err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (uu *userUsecase) ChangePassword(c context.Context, userId string, changePasswordDto *dtos.ChangePasswordDto) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()


	if changePasswordDto.NewPassword != changePasswordDto.ConfirmPassword {
		return errors.New("password don't match")
	}

	user, err := uu.userRepository.FindByID(c, userId)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	if user.ID.Hex() != userId {
		return errors.New("unatuhorized")
	}

	hashedPassword, err := services.GeneratePasswordHash(changePasswordDto.NewPassword)
	if err != nil {
		return errors.New("internal server error")
	}

	if user.Password == hashedPassword {
		return errors.New("new password can't be old password")
	}

	updateData := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return errors.New("internal server error")
	}

	return nil
}


func (uu *userUsecase) UpdateProfile(c context.Context, userId string, updateProfileDto *dtos.UpdateProfileDto) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	if err := services.ValidateStruct(uu.validator, updateProfileDto); err != nil {
		return fmt.Errorf("validation error: %v", err.Error())
	}

	existingUser, err := uu.userRepository.FindByEmailOrUsername(c, updateProfileDto.UserProfile.Username)
	if err == nil && existingUser != nil && existingUser.ID.Hex() != userId {
		return errors.New("username already exists")
	}

	user, err := uu.userRepository.FindByID(c, userId)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	if user.ID.Hex() != userId {
		return errors.New("unatuhorized")
	}

	var imageUrl string

	if updateProfileDto.Avatar != nil {
		_, err := uu.imageService.SaveProfileImage(updateProfileDto)
		if err != nil {
			return errors.New("internal server error")
		}

		url, err := uu.imageService.UploadImage(c, updateProfileDto.Avatar)
		imageUrl = url
		if err != nil {
			return errors.New("internal server error")
		}
	}

	updateData := map[string]interface{}{
		"username": user.Username,
		"profile": domain.UserProfile{
			Bio: updateProfileDto.UserProfile.Bio,
			ProfilePic: imageUrl,
			ContactInfo: updateProfileDto.UserProfile.ContactInfo,
		},
	}

	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return err
	}

	return nil
}
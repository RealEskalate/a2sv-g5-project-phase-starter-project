package user_usecase

import (
	domain "AAiT-backend-group-2/Domain"
	"AAiT-backend-group-2/Infrastructure/dtos"
	"AAiT-backend-group-2/Infrastructure/services"
	"AAiT-backend-group-2/Infrastructure/utils"
	"time"

	"golang.org/x/net/context"
)


type userUsecase struct {
	userRepository domain.UserRepository
	jwtService services.JWTService
	emailService services.EmailService
	imageService services.ImageService
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, jwtService services.JWTService, emailService services.EmailService, imageService services.ImageService, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		jwtService: jwtService,
		emailService: emailService,
		imageService: imageService,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) GetAllUsers(c context.Context,) ([]domain.User, domain.CodedError) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.FindAll(ctx)
}

func (uu *userUsecase) GetUserByID(c context.Context, id string) (*domain.User, domain.CodedError) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.FindByID(ctx, id)
}

func (uu *userUsecase) CreateUser(c context.Context, user domain.User) domain.CodedError {
	_, err := uu.userRepository.FindByEmailOrUsername(c, user.Email)
	if err == nil {
		return domain.NewError("email already exists", domain.ERR_CONFLICT)
	}

	_, err = uu.userRepository.FindByEmailOrUsername(c, user.Username)
	if err == nil {
		return domain.NewError("username already exists", domain.ERR_CONFLICT)
	}

	err = utils.ValidateUser(user)
	if err != nil {
		return err
	}

	hashedPassword, passwordErr := services.GeneratePasswordHash(user.Password)
	if passwordErr != nil {
		return domain.NewError("internal server error: failed to hash password", domain.ERR_INTERNAL_SERVER)
	}

	user.Password = hashedPassword

	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	count, err := uu.userRepository.CountDocuments(ctx)

	if err != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
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

func (uu *userUsecase) UpdateUser(c context.Context, id string, user *domain.User) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := utils.ValidateUsername(user.Username)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"username": user.Username,
		"update_at": time.Now(),
		"profile": user.Profile,
	}

	return uu.userRepository.Update(ctx, id, updateData)
}

func (uu *userUsecase) DeleteUser(c context.Context, id string) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Delete(ctx, id)
}

func (uu *userUsecase) Login(c context.Context, loginDto *dtos.LoginDTO) (map[string]string, domain.CodedError) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.FindByEmailOrUsername(ctx, loginDto.EmailOrUserName)

	if err != nil || user == nil {
		return nil, domain.NewError("invalid email or password", domain.ERR_INVALID_CREDENTIALS)
	}

	if !services.ComparePasswordHash(loginDto.Password, user.Password) {
		return nil, domain.NewError("invalid email or password", domain.ERR_INVALID_CREDENTIALS)
	}

	tokens, tokenErr := uu.jwtService.GenerateToken(user.ID.Hex(), user.Email, user.Role, 15*time.Minute, 24*time.Hour)
	if tokenErr != nil {
		return nil, domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	updateFields := map[string]interface{}{
		"refresh_token": tokens["refresh_token"],
	}
	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateFields); err != nil {
		return nil, domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return tokens, nil
}

func (uu *userUsecase) RefreshToken(c context.Context, refreshToken string) (string, domain.CodedError) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	claims, err := uu.jwtService.ValidateToken(refreshToken, "refresh_token")
	if err != nil {
		return "", domain.NewError("invalid token", domain.ERR_INVALID_CREDENTIALS)
	}

	user, err := uu.userRepository.FindByID(ctx, claims["userID"])
	if err != nil {
		return "", domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user.RefreshToken != refreshToken {
		return "", domain.NewError("invalid token", domain.ERR_INVALID_CREDENTIALS)
	}

	newAcessToken, err := uu.jwtService.RenewToken(claims)

	if err != nil {
		return "", domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return newAcessToken, nil
}

func (uu *userUsecase) PromoteUser(c context.Context, id string) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"role": "admin",
	}

	return uu.userRepository.PromoteUser(ctx, id, updateData)
}

func (uu *userUsecase) DemoteAdmin(c context.Context, id string) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"role": "admin",
	}

	return uu.userRepository.DemoteAdmin(ctx, id, updateData)
}

func (uu *userUsecase) Logout(c context.Context, id string) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	updateData := map[string]interface{}{
		"refresh_token": "",
	}

	return uu.userRepository.Update(ctx, id, updateData)
}


func (uu *userUsecase) ForgotPassword(c context.Context, userId, email string) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.FindByEmailOrUsername(c, email)
	if err != nil {
		return domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user != nil && user.ID.Hex() != userId {
		return domain.NewError("unatuhorized", domain.ERR_UNAUTHORIZED)
	}

	token, tokenErr := services.GenerateSecureToken(32)
	hashedToken := services.HashToken(token)
	if tokenErr != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	_, err = uu.userRepository.ForgotPassword(ctx, user.Email, hashedToken)

	if err != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	resetModel, emailErr := uu.emailService.GeneratePasswordResetTemplate(email, user.Username, token)
	if emailErr != nil {
		return domain.NewError("internal server error: unable to send reset email", domain.ERR_INTERNAL_SERVER)
	}

	if err := uu.emailService.SendEmail(email, "Password Reset", "reset_email.html", resetModel); err != nil {
		return domain.NewError("internal server error: unable to send reset email", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (uu *userUsecase) ResetPassword(c context.Context,  userId string, passwordResetDto *dtos.PasswordResetDto) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	if passwordResetDto.Password != passwordResetDto.ConfirmPassword {
		return domain.NewError("password don't match", domain.ERR_BAD_REQUEST)
	}

	user, err := uu.userRepository.FindByEmailOrUsername(ctx, passwordResetDto.Email)
	if err != nil {
		return domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user == nil {
		return domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user.ID.Hex() != userId {
		return domain.NewError("unatuhorized", domain.ERR_UNAUTHORIZED)
	}

	hashedToken := services.HashToken(passwordResetDto.Token)

	if err := uu.userRepository.ValidateResetToken(c, user.ID.Hex(), hashedToken); err != nil {
		return err
	}

	hashedPassword, passwordErr := services.GeneratePasswordHash(passwordResetDto.Password)
	if passwordErr != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	updateData := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	if err := uu.userRepository.InvalidateResetToken(c, user.ID.Hex()); err != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (uu *userUsecase) ChangePassword(c context.Context, userId string, changePasswordDto *dtos.ChangePasswordDto) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()


	if changePasswordDto.NewPassword != changePasswordDto.ConfirmPassword {
		return domain.NewError("password don't match", domain.ERR_BAD_REQUEST)
	}

	user, err := uu.userRepository.FindByID(c, userId)
	if err != nil || user == nil {
		return domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user.ID.Hex() != userId {
		return domain.NewError("unatuhorized", domain.ERR_UNAUTHORIZED)
	}

	hashedPassword, passwordErr := services.GeneratePasswordHash(changePasswordDto.NewPassword)
	if passwordErr != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	if user.Password == hashedPassword {
		return domain.NewError("password is the same as the current password", domain.ERR_BAD_REQUEST)
	}

	updateData := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := uu.userRepository.Update(ctx, user.ID.Hex(), updateData); err != nil {
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}


func (uu *userUsecase) UpdateProfile(c context.Context, userId string, updateProfileDto *dtos.UpdateProfileDto) domain.CodedError {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	existingUser, err := uu.userRepository.FindByEmailOrUsername(c, updateProfileDto.UserProfile.Username)
	if err == nil && existingUser != nil && existingUser.ID.Hex() != userId {
		return domain.NewError("username already exists", domain.ERR_CONFLICT)
	}

	user, err := uu.userRepository.FindByID(c, userId)
	if err != nil || user == nil {
		return domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if user.ID.Hex() != userId {
		return domain.NewError("unatuhorized", domain.ERR_UNAUTHORIZED)
	}

	var imageUrl string

	if updateProfileDto.Avatar != nil {
		_, err := uu.imageService.SaveProfileImage(updateProfileDto)
		if err != nil {
			return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
		}

		url, err := uu.imageService.UploadImage(c, updateProfileDto.Avatar)
		imageUrl = url
		if err != nil {
			return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
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
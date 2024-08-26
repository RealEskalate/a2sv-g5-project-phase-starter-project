package usecase

import (
	"blog/domain"
	"blog/internal/tokenutil"
	"blog/internal/userutil"
	"context"
	"errors"
	"net/smtp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signupUsecase struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	otpRepository   domain.OTPRepository
	contextTimeout  time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, tokenRepository domain.TokenRepository, otpRepository domain.OTPRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		otpRepository:   otpRepository,
		contextTimeout:  timeout,
	}
}

func (su *signupUsecase) RegisterUser(c context.Context, user *domain.AuthSignup) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	hashedPassword, err := userutil.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	adduser := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	err = su.userRepository.CreateUser(ctx, adduser)
	return &adduser.ID, err
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (su *signupUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (su *signupUsecase) SaveOTP(c context.Context, otp *domain.OTP) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.SaveOTP(ctx, otp)
}

func (su *signupUsecase) GetOTPByEmail(c context.Context, email string) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.GetOTPByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.AuthSignup, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.AuthSignup, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (su *signupUsecase) SaveRefreshToken(c context.Context, token string, userid primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	tkn := domain.Token{
		ID:           primitive.NewObjectID(),
		UserID:       userid,
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7),
	}
	defer cancel()
	return su.tokenRepository.SaveToken(ctx, &tkn)
}

func (su *signupUsecase) VerifyOTP(c context.Context, otp *domain.OTPRequest) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	storedOTP, err := su.GetOTPByEmail(ctx, otp.Email)
	if err != nil {
		return nil, errors.New("OTP not found please signup again")
	}
	if storedOTP.Value != otp.Value {
		return nil, errors.New("invalid OTP")
	}
	if time.Now().After(storedOTP.ExpiresAt) {
		return nil, errors.New("OTP expired")
	}
	err = su.otpRepository.DeleteOTP(ctx, storedOTP.Email)
	if err != nil {
		return nil, err
	}
	return storedOTP, nil
}

func (su *signupUsecase) SendOTP(c context.Context, user *domain.AuthSignup, smtpusername, smtppassword string, devicePrint string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	if !userutil.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !userutil.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long")
	}

	storedOTP, _ := su.GetOTPByEmail(ctx, user.Email)

	if storedOTP != nil {
		if time.Now().Before(storedOTP.ExpiresAt) {
			return errors.New("OTP already sent")
		}

		errChan := make(chan error, 1)
		go func() {
			errChan <- su.otpRepository.DeleteOTP(ctx, storedOTP.Email)
		}()

		otp := domain.OTP{
			Value:     userutil.GenerateOTP(),
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Minute * 5),
			DeviceFingerprint: devicePrint,
		}

		if err := su.SaveOTP(ctx, &otp); err != nil {
			return err
		}

		if err := <-errChan; err != nil {
			return err
		}

		if err := su.SendEmail(user.Email, otp.Value, smtpusername, smtppassword); err != nil {
			return err
		}
	}
	return nil
}

func (su *signupUsecase) SendEmail(email string, otpValue, smtpusername, smtppassword string) error {
	errChan := make(chan error, 1)
	go func() {
		from := smtpusername
		password := smtppassword

		to := []string{email}

		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		message := []byte("Your OTP is " + otpValue)

		auth := smtp.PlainAuth("", from, password, smtpHost)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {

			// use a channel to signal the error
			errChan <- err
			return
		}
	}()
	return nil
}

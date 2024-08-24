package infrastructure

import (
	"context"
	"fmt"
	"math/big"

	"crypto/rand"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type OTPService struct {
	repo interfaces.URLServiceRepository
	ctx  context.Context
}

func NewOTPService(repo interfaces.URLServiceRepository) interfaces.OTPService {
	return &OTPService{
		repo: repo,
		ctx:  context.Background(),
	}
}

func (uc *OTPService) GenerateOTP(token string) (string, *models.ErrorResponse) {
	max := big.NewInt(1000000)

	otp, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", models.InternalServerError("Error while generating OTP")
	}

	otp_code := fmt.Sprintf("%06d", otp.Int64())

	newURL := models.URL{
		ShortURLCode: otp_code,
		Token:        token,
	}

	if err := uc.repo.SaveURL(newURL, uc.ctx); err != nil {
		return "", models.InternalServerError("Error while saving the URL")
	}

	return otp_code, nil
}
func (uc *OTPService) RemoveOTP(otp_code string) *models.ErrorResponse {

	if _, err := uc.repo.GetURL(otp_code, uc.ctx); err != nil {
		return err
	}

	if err := uc.repo.DeleteURL(otp_code, uc.ctx); err != nil {
		return models.InternalServerError("Error while deleting the URL")
	}

	return nil
}

func (uc *OTPService) GetOTP(otp_code string) (*models.URL, *models.ErrorResponse) {
	url, err := uc.repo.GetURL(otp_code, uc.ctx)

	if err != nil {
		return nil, err
	}

	return url, nil
}

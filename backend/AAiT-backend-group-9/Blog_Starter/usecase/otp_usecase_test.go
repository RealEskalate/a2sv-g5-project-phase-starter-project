package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks"
	"context"
	"testing"
	"errors"
	"time"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OtpUseCaseTestSuite struct {
	suite.Suite
	otpUseCase    domain.OtpUsecase
	otpRepository *mocks.OtpRepository
}

func (suite *OtpUseCaseTestSuite) SetupTest() {
	suite.otpRepository = new(mocks.OtpRepository)
	suite.otpUseCase = NewOtpUsecase(suite.otpRepository, 2*time.Second)
}

func (suite *OtpUseCaseTestSuite) TestSaveOtp_Success() {
	otp := &domain.Otp{
		ID:        primitive.NewObjectID(),
		Email:     "test@example.com",
		Otp:       "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.otpRepository.On("SaveOtp", mock.Anything, otp).Return(nil)

	err := suite.otpUseCase.SaveOtp(context.Background(), otp)
	suite.NoError(err)
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestSaveOtp_Error() {
	otp := &domain.Otp{
		ID:        primitive.NewObjectID(),
		Email:     "test@example.com",
		Otp:       "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.otpRepository.On("SaveOtp", mock.Anything, otp).Return(errors.New("database error"))

	err := suite.otpUseCase.SaveOtp(context.Background(), otp)
	suite.Error(err)
	suite.EqualError(err, "database error")
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestGetOtpByEmail_Success() {
	expectedOtp := domain.Otp{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Otp:   "123456",
	}

	suite.otpRepository.On("GetOtpByEmail", mock.Anything, expectedOtp.Email).Return(expectedOtp, nil)

	otp, err := suite.otpUseCase.GetOtpByEmail(context.Background(), expectedOtp.Email)
	suite.NoError(err)
	suite.Equal(expectedOtp, otp)
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestGetOtpByEmail_Error() {
	suite.otpRepository.On("GetOtpByEmail", mock.Anything, "test@example.com").Return(domain.Otp{}, errors.New("database error"))

	otp, err := suite.otpUseCase.GetOtpByEmail(context.Background(), "test@example.com")
	suite.Error(err)
	suite.EqualError(err, "database error")
	suite.Empty(otp)
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestInvalidateOtp_Success() {
	otp := &domain.Otp{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Otp:   "123456",
	}

	suite.otpRepository.On("InvalidateOtp", mock.Anything, otp).Return(nil)

	err := suite.otpUseCase.InvalidateOtp(context.Background(), otp)
	suite.NoError(err)
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestInvalidateOtp_Error() {
	otp := &domain.Otp{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Otp:   "123456",
	}

	suite.otpRepository.On("InvalidateOtp", mock.Anything, otp).Return(errors.New("database error"))

	err := suite.otpUseCase.InvalidateOtp(context.Background(), otp)
	suite.Error(err)
	suite.EqualError(err, "database error")
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestGetByID_Success() {
	expectedOtp := domain.Otp{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Otp:   "123456",
	}

	suite.otpRepository.On("GetByID", mock.Anything, expectedOtp.ID.Hex()).Return(expectedOtp, nil)

	otp, err := suite.otpUseCase.GetByID(context.Background(), expectedOtp.ID.Hex())
	suite.NoError(err)
	suite.Equal(expectedOtp, otp)
	suite.otpRepository.AssertExpectations(suite.T())
}

func (suite *OtpUseCaseTestSuite) TestGetByID_Error() {
	suite.otpRepository.On("GetByID", mock.Anything, "invalidID").Return(domain.Otp{}, errors.New("database error"))

	otp, err := suite.otpUseCase.GetByID(context.Background(), "invalidID")
	suite.Error(err)
	suite.EqualError(err, "database error")
	suite.Empty(otp)
	suite.otpRepository.AssertExpectations(suite.T())
}

func TestOtpUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(OtpUseCaseTestSuite))
}

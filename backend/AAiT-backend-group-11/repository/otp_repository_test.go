package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetByID(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	repo := repository.NewOtpRepository(mockCollection)

	userID := "60d5f3c4a6e6b5b0d4a3d66f"

	objectID, _ := primitive.ObjectIDFromHex(userID)
	
	expectedResponse := entities.OTP{
		ID: objectID,
		Email: "test@gmail.com",
		Code: "12345",
	}

	mockSingleresult := new(mocks.SingleResult)
	mockCollection.On("FindOne", mock.Anything, bson.M{"_id": userID}).Return(mockSingleresult, nil)
    mockSingleresult.On("Decode", mock.AnythingOfType("*entities.OTP")).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.OTP)
        *arg = expectedResponse
    }).Return(nil)

	result, err := repo.GetByID(userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockCollection.AssertExpectations(t)
}

func TestGetOtpByEmail(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	repo := repository.NewOtpRepository(mockCollection)

	expectedOtp := entities.OTP{
		Email: "test@example.com",
		Code:  "123456",
	}

	mockSingleResult := new(mocks.SingleResult)
	mockSingleResult.On("Decode", mock.AnythingOfType("*entities.OTP")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.OTP)
		*arg = expectedOtp
	}).Return(nil)

	mockCollection.On("FindOne", mock.Anything, bson.M{"email": expectedOtp.Email}, mock.Anything).Return(mockSingleResult)

	// Call the GetOtpByEmail method
	actualOtp, err := repo.GetOtpByEmail(expectedOtp.Email)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned OTP matches the expected one
	assert.Equal(t, expectedOtp, actualOtp)

	// Ensure the mock methods were called as expected
	mockSingleResult.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}


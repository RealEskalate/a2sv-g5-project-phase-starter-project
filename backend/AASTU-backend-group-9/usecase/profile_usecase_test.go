package usecase

import (
    "context"
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "blog/domain"
    "blog/domain/mocks"
)

type ProfileUsecaseSuite struct {
    suite.Suite
    userRepoMock    *mocks.UserRepository
    profileUsecase  *ProfileUsecase
}

func (suite *ProfileUsecaseSuite) SetupTest() {
    suite.userRepoMock = new(mocks.UserRepository)
    suite.profileUsecase = &ProfileUsecase{
        userRepository: suite.userRepoMock,
        contextTimeout: time.Second * 2,
    }
}

func (suite *ProfileUsecaseSuite) TestUpdateProfile_Success() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.profileUsecase.contextTimeout)
    defer cancel()
    userID := primitive.NewObjectID()
    profile := &domain.Profile{
        First_Name:      "John",
        Last_Name:       "Doe",
        Bio:             "A bio",
        Profile_Picture: "profile.jpg",
        Contact_Info: []domain.ContactInfo{
            {
                Address:      "123 Main St",
                Phone_number: "123-456-7890",
            },
        },
    }
    user := &domain.User{
        ID:              userID,
        First_Name:      profile.First_Name,
        Last_Name:       profile.Last_Name,
        Bio:             profile.Bio,
        Profile_Picture: profile.Profile_Picture,
        Contact_Info:    profile.Contact_Info,
    }

    suite.userRepoMock.On("UpdateUser", ctx, user).Return(nil)

    result, err := suite.profileUsecase.UpdateProfile(ctx, profile, userID)

    assert.NoError(suite.T(), err)
    assert.NotNil(suite.T(), result)
    assert.Equal(suite.T(), profile.First_Name, result.First_Name)
    assert.Equal(suite.T(), profile.Last_Name, result.Last_Name)
    assert.Equal(suite.T(), profile.Bio, result.Bio)
    assert.Equal(suite.T(), profile.Profile_Picture, result.Profile_Picture)
    assert.Equal(suite.T(), profile.Contact_Info, result.Contact_Info)
}

func (suite *ProfileUsecaseSuite) TestUpdateProfile_Failure() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.profileUsecase.contextTimeout)
    defer cancel()
    userID := primitive.NewObjectID()
    profile := &domain.Profile{
        First_Name:      "John",
        Last_Name:       "Doe",
        Bio:             "A bio",
        Profile_Picture: "profile.jpg",
        Contact_Info: []domain.ContactInfo{
            {
                Address:      "123 Main St",
                Phone_number: "123-456-7890",
            },
        },
    }
    user := &domain.User{
        ID:              userID,
        First_Name:      profile.First_Name,
        Last_Name:       profile.Last_Name,
        Bio:             profile.Bio,
        Profile_Picture: profile.Profile_Picture,
        Contact_Info:    profile.Contact_Info,
    }

    suite.userRepoMock.On("UpdateUser", ctx, user).Return(errors.New("failed to update profile"))

    result, err := suite.profileUsecase.UpdateProfile(ctx, profile, userID)

    assert.Error(suite.T(), err)
    assert.Nil(suite.T(), result)
    assert.Equal(suite.T(), "failed to update profile", err.Error())
}

func TestProfileUsecaseSuite(t *testing.T) {
    suite.Run(t, new(ProfileUsecaseSuite))
}
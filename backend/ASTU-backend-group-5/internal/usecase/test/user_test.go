package usecase_test


// import (
// 	"github.com/stretchr/testify/suite"
// 	"testing"
// 	"blogApp/mocks/repository"
// 	"blogApp/pkg/hash"
// 	"github.com/stretchr/testify/mock"
// 	"blogApp/internal/usecase/user"
// 	"blogApp/internal/domain"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"time"
// 	"context"
// )
	

// type UserUsecaseSuite struct {
// 	suite.Suite
// 	repository *mocks.UserRepository
// 	usecase *user.UserUsecase
// }


// func (suite *UserUsecaseSuite) SetupTest() {
// 	repository := new(mocks.UserRepository)
// 	usecase := user.NewUserUsecase(repository)
// 	suite.repository = repository
// 	suite.usecase  = usecase
// }


// func (suite *UserUsecaseSuite) TestRegisterUser_Positive() {
// 	mockUserProfile := domain.UserProfile{
// 		ProfileUrl: "https://example.com/profiles/user123",
// 		FirstName:  "John",
// 		LastName:   "Doe",
// 		Gender:     "Male",
// 		Bio:        "A passionate software developer.",
// 		Profession: "Software Engineer",
//     }


// 	mockUser := domain.User{
// 		ID:       primitive.NewObjectID(),
// 		UserName: "johndoe",
// 		Email:    "johndoe@example.com",
// 		Password: "hashedpassword123", 
// 		Profile:  mockUserProfile,
// 		Role:     "user",
// 		Created:  primitive.NewDateTimeFromTime(time.Now()),
// 		Updated:  primitive.NewDateTimeFromTime(time.Now()),
// 		Verified: true,
// 	}

// 	suite.repository.On("FindUserByEmail", context.Background(), "johndoe@example.com").Return(nil, nil)
// 	suite.repository.On("IsEmptyCollection", context.Background()).Return(true, nil)
// 	suite.repository.On("CreateUser", context.Background(), &mockUser).Return(nil)

// 	_, err := suite.usecase.RegisterUser(&mockUser)

// 	suite.Nil(err, "err is a nil pointer so no error in this process")
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *UserUsecaseSuite) TestUpdateUser_Positive() {
// 	mockUserProfile := domain.UserProfile{
// 		ProfileUrl: "https://example.com/profiles/user123",
// 		FirstName:  "John",
// 		LastName:   "Doe",
// 		Gender:     "Male",
// 		Bio:        "A passionate software developer.",
// 		Profession: "Software Engineer",
// 	}

// 	mockUser := domain.User{
// 		ID:       primitive.NewObjectID(),
// 		UserName: "johndoe",
// 		Email:    "johndoe@example.com",
// 		Password: "hashedpassword123",
// 		Profile:  mockUserProfile,
// 		Role:     "user",
// 		Created:  primitive.NewDateTimeFromTime(time.Now()),
// 		Updated:  primitive.NewDateTimeFromTime(time.Now()),
// 		Verified: true,
// 	}

	
// 	suite.repository.On("FindUserByEmail", mock.Anything, mockUser.ID.Hex()).Return(&mockUser, nil)
// 	suite.repository.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
// 	err := suite.usecase.UpdateUser(&mockUser)
// 	suite.Nil(err, "error should be nil")
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *UserUsecaseSuite) TestLogin_Positive() {
// 	password, _ := hash.HashPassword("password123")
// 	mockUser := &domain.User{
// 		ID:       primitive.NewObjectID(),
// 		UserName: "johndoe",
// 		Email:    "johndoe@example.com",
// 		Password: password, 
// 		Role:     "user",
// 	}

// 	suite.repository.On("FindUserByEmail", mock.Anything, mockUser.Email).Return(mockUser, nil)
// 	user, token, err := suite.usecase.Login(mockUser.Email, "password123")
// 	suite.Nil(err)
// 	suite.Equal(mockUser.Email, user.Email)
// 	suite.Equal(mockUser.Password, user.Password)
// 	suite.NotNil(token)
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *UserUsecaseSuite) TestLogin_InvalidPassword_Negative() {
// 	password, _ := hash.HashPassword("password123")
// 	mockUser := &domain.User{
// 		ID:       primitive.NewObjectID(),
// 		UserName: "johndoe",
// 		Email:    "johndoe@example.com",
// 		Password: password,
// 		Role:     "user",
// 	}

// 	suite.repository.On("FindUserByEmail", mock.Anything, "johndoe@example.com").Return(mockUser, nil)

// 	user, token, err := suite.usecase.Login("johndoe@example.com", "wrongpassword")

// 	suite.NotNil(err)
// 	suite.Nil(user)
// 	suite.Nil(token)
// 	suite.EqualError(err, "invalid credentials")
// 	suite.repository.AssertExpectations(suite.T())
// }



// func TestUserUsecaseSuite(t *testing.T) {
// 	suite.Run(t, new(UserUsecaseSuite))
// }
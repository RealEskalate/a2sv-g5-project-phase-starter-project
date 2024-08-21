package test

import (
	"errors"
	"testing"
	"time"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCaseSuite struct {
	suite.Suite
	usecase *usecase.UserUseCase
	repo    *mocks.User_Repository_interface
}

func (suite *UserUseCaseSuite) SetupTest() {
	repo := new(mocks.User_Repository_interface)
	suite.usecase = usecase.NewUserUseCase(repo)
	suite.repo = repo
}

func (suite *UserUseCaseSuite) TestGetOneUser() {
	id := primitive.NewObjectID()
	pp := domain.Media{
		Uplaoded_date: time.Now(),
		Path: "path/img.png",
		ID: primitive.NewObjectID(),
	}
	user := domain.User{
		ID:    id,
		UserName: "username",
		Email: "test@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: pp,
	}

	suite.repo.On("GetUserDocumentByID", id.Hex()).Return(user, nil)

	responseUser, err := suite.usecase.GetOneUser(id.Hex())

	suite.NoError(err)
	suite.Equal(domain.CreateResponseUser(user), responseUser)
}

func (suite *UserUseCaseSuite) TestGetOneUser_Error() {
	id := primitive.NewObjectID().Hex()
	suite.repo.On("GetUserDocumentByID", id).Return(domain.User{}, errors.New("user not found"))

	responseUser, err := suite.usecase.GetOneUser(id)

	suite.Error(err)
	suite.Equal(domain.ResponseUser{}, responseUser)
}


func (suite *UserUseCaseSuite) TestGetUsers() {
	pp := domain.Media{
		Uplaoded_date: time.Now(),
		Path: "path/img.png",
		ID: primitive.NewObjectID(),
	}
	users := []domain.User{
		{
			ID:    primitive.NewObjectID(),
			UserName: "username",
			Email: "test@gmail.com",
			Is_Admin: false,
			Bio: "test bio",
			ProfilePicture: pp,
		},
	}

	suite.repo.On("GetUserDocuments").Return(users, nil)

	responseUsers, err := suite.usecase.GetUsers()

	suite.NoError(err)
	suite.Len(responseUsers, len(users))
}


func (suite *UserUseCaseSuite) TestGetUsers_Error() {
	suite.repo.On("GetUserDocuments").Return(nil, errors.New("failed to retrieve users"))

	responseUsers, err := suite.usecase.GetUsers()

	suite.Error(err)
	suite.Empty(responseUsers)
}


func (suite *UserUseCaseSuite) TestUpdateUser() {
	id := primitive.NewObjectID()
	updateUser := domain.UpdateUser{
		UserName: "updated username",
	}

	pp := domain.Media{
		Uplaoded_date: time.Now(),
		Path: "path/img.png",
		ID: primitive.NewObjectID(),
	}

	updatedUser := domain.User{
		ID:    id,
		UserName: "updated username",
		Email: "test@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: pp,
	}

	suite.repo.On("UpdateUserDocument", id.Hex(), updateUser).Return(updatedUser, nil)

	responseUser, err := suite.usecase.UpdateUser(id.Hex(), updateUser)

	suite.NoError(err)
	suite.Equal(domain.CreateResponseUser(updatedUser), responseUser)
}


func (suite *UserUseCaseSuite) TestUpdateUser_Error() {
	id := primitive.NewObjectID().Hex()
	updateUser := domain.UpdateUser{
		UserName: "updated username",
	}

	suite.repo.On("UpdateUserDocument", id, updateUser).Return(domain.User{}, errors.New("update failed"))

	responseUser, err := suite.usecase.UpdateUser(id, updateUser)

	suite.Error(err)
	suite.Equal(domain.ResponseUser{}, responseUser)
}


func (suite *UserUseCaseSuite) TestDeleteUser() {
	id := primitive.NewObjectID().Hex()

	suite.repo.On("DeleteUserDocument", id).Return(nil)

	err := suite.usecase.DeleteUser(id)

	suite.NoError(err)
}


func (suite *UserUseCaseSuite) TestDeleteUser_Error() {
	id := primitive.NewObjectID().Hex()

	suite.repo.On("DeleteUserDocument", id).Return(errors.New("delete failed"))

	err := suite.usecase.DeleteUser(id)

	suite.Error(err)
}


// func (suite *UserUseCaseSuite) TestLogIn() {
// 	loginUser := domain.LogINUser{
// 		UserName: "username",
// 		Email: "test@gmail.com",
// 		Password: "password",
// 	}

// 	pp := domain.Media{
// 		Uplaoded_date: time.Now(),
// 		Path: "path/img.png",
// 		ID: primitive.NewObjectID(),
// 	}

// 	loggedUser := domain.User{
// 		ID:    primitive.NewObjectID(),
// 		UserName: "username",
// 		Email: "test@gmail.com",
// 		Is_Admin: false,
// 		Bio: "test bio",
// 		ProfilePicture: pp,
// 	}

// 	suite.repo.On("LogIn", loginUser).Return(loggedUser, nil)

// 	responseUser, err := suite.usecase.LogIn(loginUser)

// 	suite.NoError(err)
// 	suite.Equal(domain.CreateResponseUser(loggedUser), responseUser)
// }


// func (suite *UserUseCaseSuite) TestLogIn_Error() {
// 	loginUser := domain.LogINUser{
// 		UserName: "username",
// 		Email: "test@gmail.com",
// 		Password: "password",
// 	}

// 	suite.repo.On("LogIn", loginUser).Return(domain.User{}, errors.New("login failed"))

// 	responseUser, err := suite.usecase.LogIn(loginUser)

// 	suite.Error(err)
// 	suite.Equal(domain.ResponseUser{}, responseUser)
// }


// func (suite *UserUseCaseSuite) TestRegister() {
	
	
// 	registerUser := domain.RegisterUser{
// 		UserName: "username",
// 		Bio: "test bio",
// 		Email: "test@gmail.com",
// 		Password: "password",
// 	}
// 	pp := domain.Media{
// 		Uplaoded_date: time.Now(),
// 		Path: "path/img.png",
// 		ID: primitive.NewObjectID(),
// 	}
// 	registeredUser := domain.User{
// 		ID:    primitive.NewObjectID(),
// 		UserName: "username",
// 		Email: "test@gmail.com",
// 		Is_Admin: false,
// 		Bio: "test bio",
// 		ProfilePicture: pp,
// 	}

// 	suite.repo.On("Register", registerUser).Return(registeredUser, nil)

// 	responseUser,err := suite.usecase.Register(registerUser)
// 	suite.Equal(domain.CreateResponseUser(registeredUser), responseUser)
// 	suite.NoError(err)
// }

// func (suite *UserUseCaseSuite) TestRegister_Error() {
// 	registerUser := domain.RegisterUser{
// 		UserName: "username",
// 		Bio: "test bio",
// 		Email: "test@gmail.com",
// 		Password: "password",
// 	}

// 	suite.repo.On("Register", registerUser).Return(domain.User{}, errors.New("registration failed"))

// 	responseUser, err := suite.usecase.Register(registerUser)

// 	suite.Error(err)
// 	suite.Equal(domain.ResponseUser{}, responseUser)
// }


func (suite *UserUseCaseSuite) TestFilterUser() {
	filter := map[string]string{"email": "test@gmail.com"}
	
	pp := domain.Media{
		Uplaoded_date: time.Now(),
		Path: "path/img.png",
		ID: primitive.NewObjectID(),
	}
	
	users := []domain.User{
		{
			ID:    primitive.NewObjectID(),
			UserName: "username",
			Email: "test@gmail.com",
			Is_Admin: false,
			Bio: "test bio",
			ProfilePicture: pp,
		},
	}

	suite.repo.On("FilterUserDocument", filter).Return(users, nil)

	responseUsers, err := suite.usecase.FilterUser(filter)

	suite.NoError(err)
	suite.Len(responseUsers, len(users))
}


func (suite *UserUseCaseSuite) TestFilterUser_Error() {
	filter := map[string]string{"email": "test@gmail.com"}

	suite.repo.On("FilterUserDocument", filter).Return(nil, errors.New("filter failed"))

	responseUsers, err := suite.usecase.FilterUser(filter)

	suite.Error(err)
	suite.Empty(responseUsers)
}

func TestUserUseCase(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}

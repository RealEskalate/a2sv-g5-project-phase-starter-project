package userusecase_test

import (
	"blogs/domain"
	"blogs/mocks"
	"blogs/usecase/userusecase"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	usecase *userusecase.UserUsecase
	repo    *mocks.UserRepository
	oauth  *mocks.OAuthStateRepository
}

func (s *UserUsecaseSuite) SetupTest() {
	s.repo = &mocks.UserRepository{}
	s.oauth = &mocks.OAuthStateRepository{}
	s.usecase = userusecase.NewUserUsecase(s.repo, s.oauth)
}

func (s *UserUsecaseSuite) TestRegisterUser() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
		Email: "testemail@gmail.com",
	}

	s.repo.On("RegisterUser", user).Return(user, nil).Once()

	err := s.usecase.RegisterUser(user)

	s.NoError(err)
	s.Equal(user, err)
}

func (s *UserUsecaseSuite) TestLoginUser() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}

	s.repo.On("LoginUser", user).Return(user, nil).Once()

	_,_ ,err := s.usecase.LoginUser(user.Username, user.Password)

	s.NoError(err)
	s.Equal(user, err)
}

func (s *UserUsecaseSuite) TestLogoutUser() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}

	s.repo.On("LogoutUser", user).Return(user, nil).Once()

	err := s.usecase.LogoutUser(user.Username)

	s.NoError(err)
	s.Equal(user, err)
}

func (s *UserUsecaseSuite) TestUpdateUser() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}

	claims := &domain.LoginClaims{
		Username: "testuser",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("UpdateUser", user).Return(user, nil).Once()

	err := s.usecase.UpdateProfile(user, claims)

	s.NoError(err)
	s.Equal(user, err)
}


func (s *UserUsecaseSuite) TestUserPromote() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}

	claims := &domain.LoginClaims{
		Username: "testuser",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("PromoteUser", user).Return(user, nil).Once()

	err := s.usecase.PromoteUser(user.Username,true, claims)

	s.NoError(err)
	s.Equal(user, err)
}

func (s *UserUsecaseSuite) TestUserDemote() {
	user := &domain.User{
		Username: "testuser",
		Password : "testpassword",
	}

	claims := &domain.LoginClaims{
		Username: "testuser",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("PromoteUser", user).Return(user, nil).Once()

	err := s.usecase.PromoteUser(user.Username,false, claims)

	s.NoError(err)
	s.Equal(user, err)
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}




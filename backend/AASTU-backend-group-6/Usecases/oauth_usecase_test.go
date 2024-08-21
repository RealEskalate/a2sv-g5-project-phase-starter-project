package usecases

// // import (
// // 	domain "blogs/Domain"
// // 	"blogs/mocks"
// // 	"time"

// // 	"github.com/stretchr/testify/suite"
// // )

// // type OAuthUsecaseTestSuite struct {
// // 	suite.Suite
// // 	oauthusecase   domain.OauthUsecase
// // 	mockSignUpRepo *mocks.SignupRepository
// // 	OauthConfig    *mocks.OauthConfig
// // 	contextTimeout time.Duration
// // }

// // func (suite *OAuthUsecaseTestSuite) SetupTest() {
// // 	suite.oauthusecase = new(mocks.OauthUsecase)
// // 	suite.mockSignUpRepo = new(mocks.SignupRepository)
// // 	suite.OauthConfig = new(mocks.OauthConfig)
// // 	suite.contextTimeout = time.Second * 5
// // 	suite.oauthusecase = NewOauthUsecase(suite.mockSignUpRepo, suite.contextTimeout , suite.OauthConfig)
// // }

// // func (suite *OAuthUsecaseTestSuite) TestOauthService() {

// // }
// package usecases

// import (
// 	domain "blogs/Domain"
// 	"blogs/mocks"
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type OAuthUsecaseTestSuite struct {
// 	suite.Suite
// 	oauthusecase   *OauthUseCase
// 	mockSignUpRepo *mocks.SignupRepository
// 	oauthConfig    *mocks.OauthConfig
// 	contextTimeout time.Duration
// }

// func (suite *OAuthUsecaseTestSuite) SetupTest() {
// 	suite.oauthusecase = &OauthUseCase{}
// 	suite.mockSignUpRepo = &mocks.SignupRepository{}
// 	suite.oauthConfig = &mocks.OauthConfig{}
// 	suite.contextTimeout = time.Second * 5
// 	suite.oauthusecase = NewOauthUsecase(suite.mockSignUpRepo, suite.contextTimeout, suite.oauthConfig)
// }

// func (suite *OAuthUsecaseTestSuite) TestOauthCallback_Success() {
// 	ctx := context.TODO()
// 	query := "example_query"

// 	// Mock the dependencies
// 	suite.oauthConfig.On("InitialConfig").Return(nil, nil)
// 	suite.mockSignUpRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(nil, nil)
// 	suite.mockSignUpRepo.On("Create", mock.Anything, mock.Anything).Return(nil, nil)

// 	// Call the function
// 	result := suite.oauthusecase.OauthCallback(ctx, query)

// 	// Assert the result
// 	assert.NotNil(suite.T(), result)
// 	assert.IsType(suite.T(), &domain.UserResponse{}, result)
// }

// func (suite *OAuthUsecaseTestSuite) TestOauthCallback_Error() {
// 	ctx := context.TODO()
// 	query := "example_query"

// 	// Mock the dependencies
// 	suite.oauthConfig.On("InitialConfig").Return(nil, nil)
// 	suite.mockSignUpRepo.On("FindUserByEmail", mock.Anything, mock.Anything).Return(nil, nil)
// 	suite.mockSignUpRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error creating user"))

// 	// Call the function
// 	result := suite.oauthusecase.OauthCallback(ctx, query)

// 	// Assert the result
// 	assert.NotNil(suite.T(), result)
// 	assert.IsType(suite.T(), &domain.ErrorResponse{}, result)
// }

// func TestOAuthUsecaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(OAuthUsecaseTestSuite))
// }
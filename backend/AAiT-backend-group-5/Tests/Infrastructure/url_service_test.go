package infrastructure_test

import (
	"context"
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type URLServiceTestSuite struct {
	suite.Suite
	urlService interfaces.URLService
	repoMock   *mocks.MockURLServiceRepository
	env        *config.Env
	ctr        *gomock.Controller
}

func (suite *URLServiceTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.repoMock = mocks.NewMockURLServiceRepository(suite.ctr)
	suite.env = &config.Env{
		BASE_URL: "http://example.com",
	}

	suite.urlService = infrastructure.NewURLService(suite.env, suite.repoMock)
}

func (suite *URLServiceTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *URLServiceTestSuite) TestGenerateURL_Success() {
	token := "some-token"
	purpose := "confirmRegistration"
	expectedURLPart := suite.env.BASE_URL + "/" + purpose + "/"

	suite.repoMock.
		EXPECT().
		SaveURL(gomock.Any(), gomock.Any()).
		Return(nil)

	url, err := suite.urlService.GenerateURL(token, purpose)
	suite.Nil(err)
	suite.Contains(url, expectedURLPart)
}

func (suite *URLServiceTestSuite) TestGenerateURL_SaveURLError() {
	token := "some-token"
	purpose := "confirmRegistration"
	expectedError := models.InternalServerError("Error while saving the URL")

	suite.repoMock.
		EXPECT().
		SaveURL(gomock.Any(), gomock.Any()).
		Return(expectedError)

	url, err := suite.urlService.GenerateURL(token, purpose)
	suite.Error(err)
	suite.Equal("", url)
	suite.Equal(expectedError, err)
}

func (suite *URLServiceTestSuite) TestRemoveURL_Success() {
	shortURLCode := "short-url-code"

	suite.repoMock.
		EXPECT().
		DeleteURL(shortURLCode, context.Background()).
		Return(nil)

	err := suite.urlService.RemoveURL(shortURLCode)
	suite.Nil(err)
}

func (suite *URLServiceTestSuite) TestRemoveURL_DeleteURLError() {
	shortURLCode := "short-url-code"
	expectedError := models.InternalServerError("Error while deleting the URL")

	suite.repoMock.
		EXPECT().
		DeleteURL(shortURLCode, context.Background()).
		Return(expectedError)

	err := suite.urlService.RemoveURL(shortURLCode)
	suite.Error(err)
	suite.Equal(expectedError, err)
}

func (suite *URLServiceTestSuite) TestGetURL_Success() {
	shortURLCode := "short-url-code"
	expectedURL := &models.URL{
		ShortURLCode: shortURLCode,
		Token:        "some-token",
	}

	suite.repoMock.
		EXPECT().
		GetURL(shortURLCode, context.Background()).
		Return(expectedURL, nil)

	url, err := suite.urlService.GetURL(shortURLCode)
	suite.Nil(err)
	suite.Equal(expectedURL, url)
}

func (suite *URLServiceTestSuite) TestGetURL_GetURLError() {
	shortURLCode := "short-url-code"
	expectedError := models.InternalServerError("Error while getting the URL")

	suite.repoMock.
		EXPECT().
		GetURL(shortURLCode, context.Background()).
		Return(nil, expectedError)

	url, err := suite.urlService.GetURL(shortURLCode)
	suite.Error(err)
	suite.Nil(url)
	suite.Equal(expectedError, err)
}

func TestURLServiceTestSuite(t *testing.T) {
	suite.Run(t, new(URLServiceTestSuite))
}

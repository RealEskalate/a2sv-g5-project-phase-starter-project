package controllers_test

import (
	"blogapp/Delivery/controllers"
	"blogapp/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

// generate tag controller test suite

type UserCtrlTestSuite struct {
	suite.Suite
	ctrl        *controllers.TagController
	mockUsecase *mocks.TagUseCase
}

// setup test suite
func (suite *UserCtrlTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.TagUseCase)
	suite.ctrl = controllers.NewTagsController(suite.mockUsecase)
}

// tear down
func (suite *UserCtrlTestSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

// test create tag
func (suite *UserCtrlTestSuite) TestCreateTag() {
	// test case success 
	suite.Run("CreateTag success", func() {
		w := httptest.NewRecorder()
		c, err := gin.CreateTestContext(w)
		suite.Nil(err)
		tag := mocks.GetSampleTag()

		suite.mockUsecase.On("CreateTag", c, tag).Return(nil, http.StatusCreated)

		body, err := json.Marshal(tag)
		suite.Nil(err)

		c.Request, err = http.NewRequest(http.MethodPost, "/tags/create", bytes.NewReader(body))
		suite.Nil(err)
		}
	
}	
// run test suite
func TestUserCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(UserCtrlTestSuite))
}

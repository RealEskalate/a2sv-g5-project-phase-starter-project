package controllers_test

import (
	controllers "blogapp/Delivery/controllers"
	"blogapp/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
		c, _ := gin.CreateTestContext(w)
		tag := mocks.GetSampleTag()

		suite.mockUsecase.On("CreateTag", c, tag).Return(nil, http.StatusCreated).Once()

		body, err := json.Marshal(tag)
		suite.Nil(err)

		c.Request = httptest.NewRequest(http.MethodPost, "/tags/create", bytes.NewBuffer(body))
		suite.Nil(err)
		// fmt.Println(body)

		suite.ctrl.CreateTag(c)
		// suite.Equal(http.StatusCreated, w.Code)
	})
}

// run test suite
func TestUserCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(UserCtrlTestSuite))
}

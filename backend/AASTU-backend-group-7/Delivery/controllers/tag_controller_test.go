package controllers_test

import (
	controllers "blogapp/Delivery/controllers"
	"blogapp/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// generate tag controller test suite

type TagCtrlTestSuit struct {
	suite.Suite
	ctrl        *controllers.TagController
	mockUsecase *mocks.TagUseCase
}

// setup test suite
func (suite *TagCtrlTestSuit) SetupTest() {
	suite.mockUsecase = new(mocks.TagUseCase)
	suite.ctrl = controllers.NewTagsController(suite.mockUsecase)
}

// tear down
func (suite *TagCtrlTestSuit) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

// test create tag
func (suite *TagCtrlTestSuit) TestCreateTag() {
	// test case success
	suite.Run("CreateTag success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		tag := mocks.GetSampleTag()

		suite.mockUsecase.On("CreateTag", c, mock.Anything).Return(nil, http.StatusCreated).Once()

		body, err := json.Marshal(tag)
		suite.Nil(err)

		c.Request = httptest.NewRequest(http.MethodPost, "/tags/create", bytes.NewBuffer(body))
		suite.Nil(err)
		// fmt.Println(body)

		suite.ctrl.CreateTag(c)
		// suite.Equal(http.StatusCreated, w.Code)
	})
	// test case fail
	suite.Run("CreateTag fail", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		tag := mocks.GetSampleTag()

		suite.mockUsecase.On("CreateTag", c, mock.Anything).Return(errors.New("bad request"), http.StatusBadRequest).Once()

		body, err := json.Marshal(tag)
		suite.Nil(err)

		c.Request = httptest.NewRequest(http.MethodPost, "/tags/create", bytes.NewBuffer(body))
		suite.Nil(err)

		suite.ctrl.CreateTag(c)
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

// test get tag by slug
func (suite *TagCtrlTestSuit) TestGetTagBySlug() {
	// test case success
	suite.Run("GetTagBySlug success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		slug := "sample-tag"
		c.Params = gin.Params{
			gin.Param{Key: "slug", Value: slug},
		}

		tag := mocks.GetSampleTag()

		suite.mockUsecase.On("GetTagBySlug", c, slug).Return(tag, nil, http.StatusOK).Once()

		c.Request = httptest.NewRequest(http.MethodGet, "/tags/get/"+slug, nil)

		suite.ctrl.GetTagBySlug(c)
		suite.Equal(http.StatusOK, w.Code)
	})
	// test case fail
	suite.Run("GetTagBySlug fail", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		slug := "sample-tag"
		c.Params = gin.Params{
			gin.Param{Key: "slug", Value: slug},
		}

		suite.mockUsecase.On("GetTagBySlug", c, slug).Return(nil, errors.New("bad request"), http.StatusBadRequest).Once()

		c.Request = httptest.NewRequest(http.MethodGet, "/tags/get/"+slug, nil)

		suite.ctrl.GetTagBySlug(c)
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}		

// test delete tag
func (suite *TagCtrlTestSuit) TestDeleteTag() {
	// test case success
	suite.Run("DeleteTag success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		slug := "sample-tag"
		c.Params = gin.Params{
			gin.Param{Key: "slug", Value: slug},
		}

		suite.mockUsecase.On("DeleteTag", c, slug).Return(nil, http.StatusOK).Once()

		c.Request = httptest.NewRequest(http.MethodDelete, "/tags/delete/"+slug, nil)

		suite.ctrl.DeleteTag(c)
		suite.Equal(http.StatusOK, w.Code)
	})
	// test case fail
	suite.Run("DeleteTag fail", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		slug := "sample-tag"
		c.Params = gin.Params{
			gin.Param{Key: "slug", Value: slug},
		}

		suite.mockUsecase.On("DeleteTag", c, slug).Return(errors.New("bad request"), http.StatusBadRequest).Once()

		c.Request = httptest.NewRequest(http.MethodDelete, "/tags/delete/"+slug, nil)

		suite.ctrl.DeleteTag(c)
		suite.Equal(http.StatusBadRequest, w.Code)
	})

}

// test get all tags
func (suite *TagCtrlTestSuit) TestGetAllTags() {
	// test case success
	suite.Run("GetAllTags success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		tags := mocks.GetSampleTags()

		suite.mockUsecase.On("GetAllTags", c).Return(tags, nil, http.StatusOK).Once()

		c.Request = httptest.NewRequest(http.MethodGet, "/tags/all", nil)

		suite.ctrl.GetAllTags(c)
		suite.Equal(http.StatusOK, w.Code)
	})

	// test case fail
	suite.Run("GetAllTags fail", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		// tags := mocks.GetSampleTags()

		suite.mockUsecase.On("GetAllTags", c).Return(nil, errors.New("bad request"), http.StatusBadRequest).Once()

		c.Request = httptest.NewRequest(http.MethodGet, "/tags/all", nil)

		suite.ctrl.GetAllTags(c)
		suite.Equal(http.StatusBadRequest, w.Code)
	})

		
}

// run test suite
func TestTagCtrlTestSuit(t *testing.T) {
	suite.Run(t, new(TagCtrlTestSuit))
}

package usecases_test

import (
	"blogapp/Domain"
	usecases "blogapp/UseCases"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TagUsecaseSuite struct {
	suite.Suite
	context    context.Context
	tagUsecase *usecases.TagsUseCase
	repo       *mocks.TagRepository
}

// TagsUseCase

func (suite *TagUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.TagRepository)
	suite.tagUsecase = usecases.NewTagsUseCase(suite.repo)
	suite.context = context.Background()
}

// create Tag

func (suite *TagUsecaseSuite) TestCreateTag() {
	c, _ := gin.CreateTestContext(nil)
	tag := Domain.Tag{}
	suite.repo.On("CreateTag", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.tagUsecase.CreateTag(c, &tag)
	suite.Nil(err)
	suite.Equal(200, status)
}

// delete Tag

func (suite *TagUsecaseSuite) TestDeleteTag() {
	c, _ := gin.CreateTestContext(nil)
	slug := "fasdf"
	suite.repo.On("DeleteTag", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.tagUsecase.DeleteTag(c, slug)
	suite.Nil(err)
	suite.Equal(200, status)
}

// get all Tags

func (suite *TagUsecaseSuite) TestGetAllTags() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("GetAllTags", mock.Anything).Return([]*Domain.Tag{}, nil, 200)
	_, err, status := suite.tagUsecase.GetAllTags(c)
	suite.Nil(err)
	suite.Equal(200, status)
}

// get Tag by slug
func (suite *TagUsecaseSuite) TestGetTagBySlug() {
	c, _ := gin.CreateTestContext(nil)
	slug := "slug"
	suite.repo.On("GetTagBySlug", mock.Anything, mock.Anything).Return(&Domain.Tag{}, nil, 200)
	_, err, status := suite.tagUsecase.GetTagBySlug(c, slug)
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestTagUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TagUsecaseSuite))
}

package usecases

import (
	domain "blogs/Domain"
	"blogs/mocks"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecaseSuite struct {
	suite.Suite
	repo        mocks.BlogRepository
	blogUsecase domain.BlogUsecase
	idConverter mocks.IDConverterInterface
}

func (suite *BlogUsecaseSuite) SetupTest() {
	suite.repo = *new(mocks.BlogRepository)
	suite.idConverter = *new(mocks.IDConverterInterface)
	suite.blogUsecase = NewBlogUsecase(&suite.repo, &suite.idConverter)
}

func (suite *BlogUsecaseSuite) TestReactOnBlog() {
	suite.repo.On("ReactOnBlog", "1", true, "1").Return(domain.ErrorResponse{})

	err := suite.blogUsecase.ReactOnBlog("1", "true", "1")
	// log.Default(err)
	suite.IsTypef(err, domain.ErrorResponse{}, "should be similar in Type")
	suite.Equal(err, domain.ErrorResponse{}, "should be equal")
	// suite.NotEqual(err, domain.ErrorResponse{})
	suite.repo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseSuite) TestCommentOnBlog() {
	newID, _ := primitive.ObjectIDFromHex("1")
	comment := domain.Comment{
		Commentor_ID:       newID,
		Commentor_username: "name",
	}
	suite.repo.On("CommentOnBlog", "1", "name", comment).Return(nil)
	suite.idConverter.On("ToObjectID", "1").Return(newID)

	err := suite.blogUsecase.CommentOnBlog("user_id", comment)

	suite.IsTypef(err, nil, "should be similar in Type")
	suite.Equal(err, nil, "should be equal")
	// suite.NotEqual(err, errors.New(""))
	suite.repo.AssertExpectations(suite.T())

}
func (suite *BlogUsecaseSuite) TestCreateBlog() {
	id, _ := primitive.ObjectIDFromHex("1")
	blog := domain.Blog{
		Author:     "Eyerusalem",
		Title:      "my blogs",
		Content:    "This is my first blog",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Tags:       make([]string, 1),
		Commenters_ID:   make([]primitive.ObjectID, 1),
		Creator_id: id,
		Blog_image: "https://media.istockphoto.com/id/922745190/photo/blogging-blog-concepts-ideas-with-worktable.jpg?s=2048x2048&w=is&k=20&c=QNKuhWRD7f0P5hybe28_AHo_Wh6W93McWY157Vmmh4Q=",
	}

	suite.repo.On("CreateBlog", "1", blog, "user").Return(blog, nil)

	newBlog, err := suite.blogUsecase.CreateBlog("1", blog)
	suite.Assert().IsType(newBlog, domain.Blog{}, "must be of the same type")
	suite.Nil(err, "should no be nil")
	suite.repo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseSuite) TestDeleteBlogByID() {
	id, _ := primitive.ObjectIDFromHex("1")
	blog := domain.Blog{
		Author:     "Eyerusalem",
		Title:      "my blogs",
		Content:    "This is my first blog",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Tags:       make([]string, 1),
		Commenters_ID:   make([]primitive.ObjectID, 1),
		Creator_id: id,
		Blog_image: "https://media.istockphoto.com/id/922745190/photo/blogging-blog-concepts-ideas-with-worktable.jpg?s=2048x2048&w=is&k=20&c=QNKuhWRD7f0P5hybe28_AHo_Wh6W93McWY157Vmmh4Q=",
	}

	suite.repo.On("GetBlogByID", "1", true).Return(blog, nil)
	suite.idConverter.On("ToString", blog.Creator_id).Return("1")
	suite.repo.On("DeleteBlogByID", "1", "1").Return(domain.ErrorResponse{})

	errResponse := suite.blogUsecase.DeleteBlogByID("1", "user")
	suite.IsType(errResponse, domain.ErrorResponse{}, "should of the same type")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestFilterBlogsByTag() {
	startDate := strings.ReplaceAll("2024-08-21T08:04:55.859+00:00", " ", "+")
	endDate := strings.ReplaceAll("2024-08-21T08:04:55.859+00:00", " ", "+")
	StartDate, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		log.Fatalln(err)
	}
	EndDate, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		log.Fatalln(err)
	}
	suite.repo.On("FilterBlogsByTag", []string{}, int64(1), int64(1), StartDate, EndDate, "").Return([]domain.Blog{}, domain.Pagination{}, nil)

	blogs, pagination, err := suite.blogUsecase.FilterBlogsByTag([]string{}, "1", "1", "2024-08-21T08:04:55.859 00:00", "2024-08-21T08:04:55.859 00:00", "")

	suite.IsType(blogs, []domain.Blog{}, "should have the same value")
	suite.IsType(pagination, domain.Pagination{}, "should have the same value")
	suite.Nil(err, "should have nil value")

	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestGetBlogByID() {
	suite.repo.On("GetBlogByID", "1", true).Return(domain.Blog{}, nil)

	blog, err := suite.blogUsecase.GetBlogByID("1", true)
	suite.IsType(blog, domain.Blog{}, "should be equal")
	suite.Nil(err, "should be nil")

	suite.repo.AssertExpectations(suite.T())
}
func (suite *BlogUsecaseSuite) TestGetBlogs() {
	suite.repo.On("GetBlogs", int64(1), int64(1), "").Return([]domain.Blog{}, domain.Pagination{}, nil)
	blogs, pagination, err := suite.blogUsecase.GetBlogs("1", "1", "")

	suite.IsType(blogs, []domain.Blog{}, "should be equal")
	suite.IsType(blogs, []domain.Blog{}, "should be equal")
	suite.IsType(pagination, domain.Pagination{}, "should be equal")
	suite.Nil(err, "should be nil")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestGetMyBlogByID() {
	id, _ := primitive.ObjectIDFromHex("66c59fa762c7e4ec02998609")
	blog := domain.Blog{
		Creator_id: id,
	}
	suite.repo.On("GetMyBlogByID", "66c59fa762c7e4ec02998609", "1").Return(blog, nil)
	blog, err := suite.blogUsecase.GetMyBlogByID(blog.Creator_id.Hex(),  "user")
	suite.IsType(blog, domain.Blog{}, "should be equal")
	suite.Nil(err, "should be nil")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestGetMyBlogs() {
	suite.repo.On("GetMyBlogs", "1", int64(1), int64(1), "").Return([]domain.Blog{}, domain.Pagination{}, nil)
	blogs, pagination, err := suite.blogUsecase.GetMyBlogs("1", "1", "1", "")

	suite.IsType(blogs, []domain.Blog{}, "should be equal")
	suite.IsType(pagination, domain.Pagination{}, "should be equal")
	suite.Nil(err, "should be nil")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestSearchBlogByTitleAndAuthor() {
	suite.repo.On("SearchBlogByTitleAndAuthor", "x", "y", int64(1), int64(1), "").Return([]domain.Blog{}, domain.Pagination{}, nil)
	blogs, pagination, err := suite.blogUsecase.SearchBlogByTitleAndAuthor("x", "y", "1", "1", "")

	suite.IsType(blogs, []domain.Blog{}, "should be equal")
	suite.IsType(pagination, domain.Pagination{}, "should be equal")
	suite.IsType(err, domain.ErrorResponse{}, "should be equal")
	suite.Equal(err.Status, 0, "should be equal")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseSuite) TestUpdateBlogByID() {
	id, _ := primitive.ObjectIDFromHex("66c59fa762c7e4ec02998609")
	blog := domain.Blog{
		Creator_id: id,
	}
	suite.repo.On("UpdateBlogByID", "66c59fa762c7e4ec02998609", "1", blog).Return(blog, nil)
	suite.repo.On("GetBlogByID", "1", true).Return(blog, nil)
	blog, err := suite.blogUsecase.UpdateBlogByID("66c59fa762c7e4ec02998609", "1", blog)

	suite.IsType(blog, domain.Blog{}, "should be equal")
	suite.Nil(err, "should be nil")

	suite.repo.AssertExpectations(suite.T())
}

func TestBloBlogUsecaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseSuite))
}


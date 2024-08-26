package blogusecase_test

import (
	"blogs/domain"
	"blogs/mocks"
	"blogs/usecase/blogusecase"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecaseSuite struct {
	suite.Suite
	usecase *blogusecase.BlogUsecase
	repo    *mocks.BlogRepository
}

func (s *BlogUsecaseSuite) SetupTest() {
	s.repo = &mocks.BlogRepository{}
	s.usecase = blogusecase.NewBlogUsecase(s.repo, nil)
}

func (s *BlogUsecaseSuite) TestInsertBlog () {
	blog := &domain.Blog{
		Title: "test title",
		Content: "testn content",
		Author: "test author",
		Tags: []string{"test"},
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount: 0,
		LikesCount: 0,
		CommentsCount: 0,
	}

	s.repo.On("InsertBlog", blog).Return(blog, nil)

	newblog, err := s.usecase.InsertBlog(blog)

	s.NoError(err)
	s.Equal(blog, newblog)
}

func (s *BlogUsecaseSuite) TestGetBlogByID () {
	blog := &domain.Blog{
		Title: "test title",
		Content: "testn content",
		Author: "test author",
		Tags: []string{"test"},
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount: 0,
		LikesCount: 0,
		CommentsCount: 0,
	}

	s.repo.On("GetBlogByID", primitive.NewObjectID()).Return(blog, nil)
	
	newblog, err := s.usecase.GetBlogByID(primitive.NewObjectID().Hex())

	s.NoError(err)
	s.Equal(blog, newblog)
}

func (s *BlogUsecaseSuite) TestUpdateBlogByID () {
	blog := &domain.Blog{
		Title: "test title",
		Content: "testn content",
		Author: "test author",
		Tags: []string{"test"},
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount: 0,
		LikesCount: 0,
		CommentsCount: 0,
	}

	claim := &domain.LoginClaims{
		Username: "test author",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("UpdateBlogByID", primitive.NewObjectID().Hex(), blog).Return(nil)

	result , err := s.usecase.UpdateBlogByID(primitive.NewObjectID().Hex(), blog, claim)

	s.NoError(err)
	s.Equal(blog, result)


}

func (s *BlogUsecaseSuite) TestDeleteBlogByID () {
	claim := &domain.LoginClaims{
		Username: "test author",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("DeleteBlogByID", primitive.NewObjectID().Hex()).Return(nil)

	err := s.usecase.DeleteBlogByID(primitive.NewObjectID().Hex(), claim)

	s.NoError(err)
}

func (s *BlogUsecaseSuite) TestSearchBlog () {
	blog := &domain.Blog{
		Title: "test title",
		Content: "testn content",
		Author: "test author",
		Tags: []string{"test"},
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount: 0,
		LikesCount: 0,
		CommentsCount: 0,
	}

	s.repo.On("SearchBlog", "test title", "test author", []string{"test"}).Return([]*domain.Blog{blog}, nil)

	result, err := s.usecase.SearchBlog("test title", "test author", []string{"test"})

	s.NoError(err)
	s.Equal([]*domain.Blog{blog}, result)
}

func (s *BlogUsecaseSuite) TestFilterBlog () {
	blog := &domain.Blog{
		Title: "test title",
		Content: "testn content",
		Author: "test author",
		Tags: []string{"test"},
		CreatedAt: time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount: 0,
		LikesCount: 0,
		CommentsCount: 0,
	}

	s.repo.On("FilterBlog", []string{"test"}, time.Now(), time.Now()).Return([]*domain.Blog{blog}, nil)
	
	result, err := s.usecase.FilterBlog([]string{"test"}, time.Now(), time.Now())

	s.NoError(err)
	s.Equal([]*domain.Blog{blog}, result)
}

func (s *BlogUsecaseSuite) TestAddView () {
	objId := primitive.NewObjectID() 
	view := &domain.View{
		BlogID: objId,
		User: "test user",
	}

	claim := &domain.LoginClaims{
		Username: "test author",
		Role: "admin",
		Type: "access",
	}

	s.repo.On("AddView", view).Return(nil)

	err := s.usecase.AddView([]primitive.ObjectID{objId}, *claim)

	s.NoError(err)
}

func (s *BlogUsecaseSuite) TestAddLike () {
	objId := primitive.NewObjectID() 
	like := &domain.Like{
		BlogID: objId,
		User: "test user",
		Like: true,
	}

	s.repo.On("AddLike", like).Return(nil)

	err := s.usecase.AddLike(like)

	s.NoError(err)
}

func (s *BlogUsecaseSuite) TestGenerateContent(){
	prompt := "test prompt"

	s.repo.On("GenerateAIContent", prompt).Return("test content", nil)

	result, err := s.usecase.GenerateAiContent(prompt)

	s.NoError(err)
	s.Equal("test content", result)


}

func TestBlogUsecaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseSuite))
}





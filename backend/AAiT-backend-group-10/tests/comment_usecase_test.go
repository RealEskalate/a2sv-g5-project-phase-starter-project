package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommentUseCaseSuite struct {
	suite.Suite
	commentRepo    *mocks.CommentRepositoryInterface
	userRepo       *mocks.IUserRepository
	cacheRepo      *mocks.CacheRepoInterface
	commentUseCase usecases.CommentUsecaseInterface
}

func (suite *CommentUseCaseSuite) SetupTest() {
	commentRepo := new(mocks.CommentRepositoryInterface)
	cacheRepo := new(mocks.CacheRepoInterface)
	userRepo := new(mocks.IUserRepository)
	commentUC := usecases.NewCommentUsecase(commentRepo, userRepo, cacheRepo)
	suite.commentRepo = commentRepo
	suite.cacheRepo = cacheRepo
	suite.commentUseCase = commentUC
	suite.userRepo = userRepo
}

func (suite *CommentUseCaseSuite) TearDownTest() {
	suite.commentRepo.AssertExpectations(suite.T())
	suite.cacheRepo.AssertExpectations(suite.T())
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *CommentUseCaseSuite) TestGetCommentByID_Positive() {
	dummyComment := domain.Comment{
		ID:          uuid.New(),
		CommenterID: uuid.New(),
	}
	commenter := &domain.User{
		FullName: "SomeOne",
	}
	suite.commentRepo.On("GetCommentByID", dummyComment.ID).Return(dummyComment, nil)
	suite.userRepo.On("GetUserByID", dummyComment.CommenterID).Return(commenter, nil)
	result, err := suite.commentUseCase.GetCommentByID(dummyComment.ID)
	suite.Nil(err)
	suite.NotNil(result)
}

func (suite *CommentUseCaseSuite) TestGetCommentByID_Negative() {
	commentID := uuid.New()
	suite.commentRepo.On("GetCommentByID", commentID).Return(domain.Comment{}, domain.ErrCommentNotFound)
	result, err := suite.commentUseCase.GetCommentByID(commentID)
	suite.Nil(result)
	suite.Equal(domain.ErrCommentNotFound, err)
}

func (suite *CommentUseCaseSuite) TestAddComment_CaheExists_Positive() {
	comment := domain.Comment{
		BlogID:      uuid.New(),
		CommenterID: uuid.New(),
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("AddComment", mock.AnythingOfType("domain.Comment")).Return(nil)
	suite.cacheRepo.On("Delete", mock.AnythingOfType("string")).Return(nil)
	err := suite.commentUseCase.AddComment(comment)
	suite.Nil(err)
}
func (suite *CommentUseCaseSuite) TestAddComment_CacheSet_Failed_Positive() {
	comment := domain.Comment{
		BlogID:      uuid.New(),
		CommenterID: uuid.New(),
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("AddComment", mock.AnythingOfType("domain.Comment")).Return(nil)
	suite.cacheRepo.On("Delete", mock.AnythingOfType("string")).Return(domain.ErrCacheSetFailed)
	err := suite.commentUseCase.AddComment(comment)
	suite.Nil(err)
}
func (suite *CommentUseCaseSuite) TestAddComment_Negative() {
	comment := domain.Comment{}
	suite.commentRepo.On("AddComment", mock.AnythingOfType("domain.Comment")).Return(domain.ErrCommentCreationFailed)
	err := suite.commentUseCase.AddComment(comment)
	suite.Equal(err, domain.ErrCommentCreationFailed)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_Positive() {
	requesterID := uuid.New()
	dummyComment := domain.Comment{
		ID:          uuid.New(),
		BlogID:      uuid.New(),
		CommenterID: requesterID,
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("GetCommentByID", dummyComment.ID).Return(dummyComment, nil)
	suite.commentRepo.On("DeleteComment", dummyComment.ID).Return(nil)
	suite.cacheRepo.On("Delete", fmt.Sprintf("Comments:%s", dummyComment.BlogID)).Return(nil)
	err := suite.commentUseCase.DeleteComment(dummyComment.ID, requesterID, false)
	suite.Nil(err)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_NotFound_Negative() {
	commentID := uuid.New()
	requesterID := uuid.New()
	suite.commentRepo.On("GetCommentByID", commentID).Return(domain.Comment{}, domain.ErrCommentNotFound)
	err := suite.commentUseCase.DeleteComment(commentID, requesterID, false)
	suite.Equal(domain.ErrCommentNotFound, err)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_UnAuthorized_Negative() {
	commentID := uuid.New()
	requesterID := uuid.New()
	dummyComment := domain.Comment{
		ID:          commentID,
		BlogID:      uuid.New(),
		CommenterID: uuid.New(),
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("GetCommentByID", commentID).Return(dummyComment, nil)
	err := suite.commentUseCase.DeleteComment(commentID, requesterID, false)
	suite.Equal(domain.ErrUnAuthorized, err)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_ByAdmin_Positve() {
	commentID := uuid.New()
	requesterID := uuid.New()
	blogID := uuid.New()
	dummyComment := domain.Comment{
		ID:          commentID,
		BlogID:      blogID,
		CommenterID: uuid.New(),
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("GetCommentByID", commentID).Return(dummyComment, nil)
	suite.commentRepo.On("DeleteComment", commentID).Return(nil)
	suite.cacheRepo.On("Delete", fmt.Sprintf("Comments:%s", blogID)).Return(nil)
	err := suite.commentUseCase.DeleteComment(commentID, requesterID, true)
	suite.Nil(err)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_ServerError_Negative() {
	commentID := uuid.New()
	requesterID := uuid.New()
	dummyComment := domain.Comment{
		ID:          commentID,
		BlogID:      uuid.New(),
		CommenterID: uuid.New(),
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("GetCommentByID", commentID).Return(dummyComment, nil)
	suite.commentRepo.On("DeleteComment", commentID).Return(domain.ErrCommentDeletionFailed)
	err := suite.commentUseCase.DeleteComment(commentID, requesterID, true)
	suite.Equal(domain.ErrCommentDeletionFailed, err)
}

func (suite *CommentUseCaseSuite) TestDeleteComment_CacheFailed_Positive() {
	commentID := uuid.New()
	requesterID := uuid.New()
	blogID := uuid.New()
	dummyComment := domain.Comment{
		ID:          commentID,
		BlogID:      blogID,
		CommenterID: requesterID,
		Comment:     "This is a comment",
	}
	suite.commentRepo.On("GetCommentByID", commentID).Return(dummyComment, nil)
	suite.commentRepo.On("DeleteComment", commentID).Return(nil)
	suite.cacheRepo.On("Delete", fmt.Sprintf("Comments:%s", blogID)).Return(domain.ErrCacheSetFailed)
	err := suite.commentUseCase.DeleteComment(commentID, requesterID, false)
	suite.Nil(err)
}

func (suite *CommentUseCaseSuite) TestGetComments_CachedComments_Positive() {
    // Arrange
    blogID := uuid.New()
    expectedComments := []*dto.CommentDto{
        {
            ID: uuid.New(),
            Comment:   "Great post!",
            Commenter: "John Doe",
        },
    }
    cachedData, _ := json.Marshal(expectedComments)
    suite.cacheRepo.On("Get", fmt.Sprintf("Comments:%s", blogID)).Return(string(cachedData), nil)

    // Act
    result, err := suite.commentUseCase.GetComments(blogID)

    // Assert
    suite.Nil(err)
    suite.Equal(expectedComments, result)
}

func (suite *CommentUseCaseSuite) TestGetComments_NoCache_FetchFromRepository_Positive() {
    // Arrange
    blogID := uuid.New()
    comments := []domain.Comment{
        {
            ID:         uuid.New(),
            BlogID:     blogID,
            CommenterID: uuid.New(),
            Comment:    "Great post!",
        },
    }
    commenter := &domain.User{
        ID:       comments[0].CommenterID,
        FullName: "John Doe",
    }
    expectedComments := []*dto.CommentDto{
        dto.NewCommentDto(comments[0], commenter.FullName),
    }
    suite.cacheRepo.On("Get", fmt.Sprintf("Comments:%s", blogID)).Return("", domain.ErrCacheNotFound)
    suite.commentRepo.On("GetComments", blogID).Return(comments, nil)
    suite.userRepo.On("GetUserByID", comments[0].CommenterID).Return(commenter, nil)
    suite.cacheRepo.On("Set", fmt.Sprintf("Comments:%s", blogID), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil)

    // Act
    result, err := suite.commentUseCase.GetComments(blogID)

    // Assert
    suite.Nil(err)
    suite.Equal(expectedComments, result)
    suite.commentRepo.AssertCalled(suite.T(), "GetComments", blogID)
    suite.userRepo.AssertCalled(suite.T(), "GetUserByID", comments[0].CommenterID)
    suite.cacheRepo.AssertCalled(suite.T(), "Set", fmt.Sprintf("Comments:%s", blogID), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration"))
}

func (suite *CommentUseCaseSuite) TestGetComments_RepositoryError_Negative() {
    // Arrange
    blogID := uuid.New()
    suite.cacheRepo.On("Get", fmt.Sprintf("Comments:%s", blogID)).Return("", domain.ErrCacheNotFound)
    suite.commentRepo.On("GetComments", blogID).Return(nil, domain.ErrCommentFetchFailed)

    // Act
    result, err := suite.commentUseCase.GetComments(blogID)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrCommentFetchFailed)
}

func (suite *CommentUseCaseSuite) TestGetComments_UserRepoError_Negative() {
    // Arrange
    blogID := uuid.New()
    comments := []domain.Comment{
        {
            ID:         uuid.New(),
            BlogID:     blogID,
            CommenterID: uuid.New(),
            Comment:    "Great post!",
        },
    }
    suite.cacheRepo.On("Get", fmt.Sprintf("Comments:%s", blogID)).Return("", domain.ErrCacheNotFound)
    suite.commentRepo.On("GetComments", blogID).Return(comments, nil)
    suite.userRepo.On("GetUserByID", comments[0].CommenterID).Return(&domain.User{}, domain.ErrUserNotFound)

    // Act
    result, err := suite.commentUseCase.GetComments(blogID)

    // Assert
    suite.Nil(result)
    suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *CommentUseCaseSuite) TestGetComments_CacheSetError_Ignored() {
    // Arrange
    blogID := uuid.New()
    comments := []domain.Comment{
        {
            ID:         uuid.New(),
            BlogID:     blogID,
            CommenterID: uuid.New(),
            Comment:    "Great post!",
        },
    }
    commenter := domain.User{
        ID:       comments[0].CommenterID,
        FullName: "John Doe",
    }
    expectedComments := []*dto.CommentDto{
        dto.NewCommentDto(comments[0], commenter.FullName),
    }
    suite.cacheRepo.On("Get", fmt.Sprintf("Comments:%s", blogID)).Return("", domain.ErrCacheNotFound)
    suite.commentRepo.On("GetComments", blogID).Return(comments, nil)
    suite.userRepo.On("GetUserByID", comments[0].CommenterID).Return(&commenter, nil)
    suite.cacheRepo.On("Set", fmt.Sprintf("Comments:%s", blogID), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(domain.ErrCacheSetFailed)

    // Act
    result, err := suite.commentUseCase.GetComments(blogID)

    // Assert
    suite.Nil(err)
    suite.Equal(expectedComments, result)
}

func (suite *CommentUseCaseSuite) TestUpdateComment_Success_Positive() {
    // Arrange
    requesterID := uuid.New()
    updatedComment := domain.Comment{
        ID:        uuid.New(),
        BlogID:    uuid.New(),
        CommenterID: requesterID,
        Comment:   "Updated content",
    }
    originalComment := domain.Comment{
        ID:        updatedComment.ID,
        BlogID:    updatedComment.BlogID,
        CommenterID: requesterID,
        Comment:   "Original content",
    }
    suite.commentRepo.On("GetCommentByID", updatedComment.ID).Return(originalComment, nil)
    suite.commentRepo.On("UpdateComment", updatedComment).Return(nil)
    suite.cacheRepo.On("Delete", fmt.Sprintf("Comments:%s", updatedComment.BlogID)).Return(nil)

    // Act
    err := suite.commentUseCase.UpdateComment(requesterID, updatedComment)

    // Assert
    suite.Nil(err)
    suite.commentRepo.AssertCalled(suite.T(), "UpdateComment", updatedComment)
    suite.cacheRepo.AssertCalled(suite.T(), "Delete", fmt.Sprintf("Comments:%s", updatedComment.BlogID))
}

func (suite *CommentUseCaseSuite) TestUpdateComment_Unauthorized_Negative() {
    // Arrange
    requesterID := uuid.New()
    updatedComment := domain.Comment{
        ID:        uuid.New(),
        BlogID:    uuid.New(),
        CommenterID: uuid.New(), // Different commenter
        Comment:   "Updated content",
    }
    originalComment := domain.Comment{
        ID:        updatedComment.ID,
        BlogID:    updatedComment.BlogID,
        CommenterID: uuid.New(), // Different commenter
        Comment:   "Original content",
    }
    suite.commentRepo.On("GetCommentByID", updatedComment.ID).Return(originalComment, nil)

    // Act
    err := suite.commentUseCase.UpdateComment(requesterID, updatedComment)

    // Assert
    suite.Equal(err, domain.ErrUnAuthorized)
    suite.commentRepo.AssertNotCalled(suite.T(), "UpdateComment", updatedComment)
}

func (suite *CommentUseCaseSuite) TestUpdateComment_RepositoryError_Negative() {
    // Arrange
    requesterID := uuid.New()
    updatedComment := domain.Comment{
        ID:        uuid.New(),
        BlogID:    uuid.New(),
        CommenterID: requesterID,
        Comment:   "Updated content",
    }
    originalComment := domain.Comment{
        ID:        updatedComment.ID,
        BlogID:    updatedComment.BlogID,
        CommenterID: requesterID,
        Comment:   "Original content",
    }
    suite.commentRepo.On("GetCommentByID", updatedComment.ID).Return(originalComment, nil)
    suite.commentRepo.On("UpdateComment", updatedComment).Return(domain.ErrCommentUpdateFailed)

    // Act
    err := suite.commentUseCase.UpdateComment(requesterID, updatedComment)

    // Assert
    suite.Equal(err, domain.ErrCommentUpdateFailed)
    suite.commentRepo.AssertCalled(suite.T(), "UpdateComment", updatedComment)
}

func (suite *CommentUseCaseSuite) TestUpdateComment_CacheDeleteError_Ignored() {
    // Arrange
    requesterID := uuid.New()
    updatedComment := domain.Comment{
        ID:        uuid.New(),
        BlogID:    uuid.New(),
        CommenterID: requesterID,
        Comment:   "Updated content",
    }
    originalComment := domain.Comment{
        ID:        updatedComment.ID,
        BlogID:    updatedComment.BlogID,
        CommenterID: requesterID,
        Comment:   "Original content",
    }
    suite.commentRepo.On("GetCommentByID", updatedComment.ID).Return(originalComment, nil)
    suite.commentRepo.On("UpdateComment", updatedComment).Return(nil)
    suite.cacheRepo.On("Delete", fmt.Sprintf("Comments:%s", updatedComment.BlogID)).Return(domain.ErrCacheDeleteFailed)

    // Act
    err := suite.commentUseCase.UpdateComment(requesterID, updatedComment)

    // Assert
    suite.Nil(err)
    suite.commentRepo.AssertCalled(suite.T(), "UpdateComment", updatedComment)
    suite.cacheRepo.AssertCalled(suite.T(), "Delete", fmt.Sprintf("Comments:%s", updatedComment.BlogID))
}

func TestCommentUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CommentUseCaseSuite))
}

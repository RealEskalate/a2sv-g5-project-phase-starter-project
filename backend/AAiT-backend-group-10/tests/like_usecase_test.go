package tests

import (
	"fmt"
	"testing"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LikeUseCaseSuite struct {
	suite.Suite
	likeRepo    *mocks.LikeRepositoryInterface
	cacheRepo   *mocks.CacheRepoInterface
	likeUseCase usecases.LikeUsecaseInterface
}

func (suite *LikeUseCaseSuite) SetupTest() {
	likeRepo := new(mocks.LikeRepositoryInterface)
	cacheRepo := new(mocks.CacheRepoInterface)
	likeUC := usecases.NewLikeUseCase(likeRepo, cacheRepo)
	suite.likeRepo = likeRepo
	suite.cacheRepo = cacheRepo
	suite.likeUseCase = likeUC
}
func (suite *LikeUseCaseSuite) TearDownTest() {
	suite.likeRepo.AssertExpectations(suite.T())
	suite.cacheRepo.AssertExpectations(suite.T())
}
func (suite *LikeUseCaseSuite) TestGetLike_ValidInputs_Positive() {
	// Arrange
	blogID := uuid.New()
	reacterID := uuid.New()
	islike := true
	expectedLike := &domain.Like{
		ID:        uuid.New(),
		BlogID:    blogID,
		ReacterID: reacterID,
		IsLike:    &islike,
	}
	suite.likeRepo.On("GetLike", blogID, reacterID).Return(expectedLike, nil)

	// Act
	result, err := suite.likeUseCase.GetLike(blogID, reacterID)
	// Assert
	suite.Nil(err)
	suite.Equal(expectedLike, result)
}

func (suite *LikeUseCaseSuite) TestGetLike_LikeNotFound_Negative() {
	// Arrange
	blogID := uuid.New()
	reacterID := uuid.New()
	suite.likeRepo.On("GetLike", blogID, reacterID).Return(nil, domain.ErrLikeNotFound)

	// Act
	result, err := suite.likeUseCase.GetLike(blogID, reacterID)

	// Assert
	suite.Nil(result)
	suite.Equal(err, domain.ErrLikeNotFound)
}

func (suite *LikeUseCaseSuite) TestLikeBlog_NewLike_Positive() {
	// Arrange
	islike := true
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &islike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(nil, domain.ErrLikeNotFound)
	suite.likeRepo.On("AddLike", mock.AnythingOfType("domain.Like")).Return(nil)
	suite.cacheRepo.On("Get", fmt.Sprintf("LikeCount:%s", like.BlogID)).Return("0", nil)
	suite.cacheRepo.On("Set", fmt.Sprintf("LikeCount:%s", like.BlogID), "1", mock.AnythingOfType("time.Duration")).Return(nil)

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Nil(err)
}

func (suite *LikeUseCaseSuite) TestLikeBlog_ExistingLike_UpdateToSameLike_Positive() {
	// Arrange
	islike := true
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &islike,
	}
	existingLike := &domain.Like{
		BlogID:    like.BlogID,
		ReacterID: like.ReacterID,
		IsLike:    &islike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(existingLike, nil)

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Nil(err)
	suite.likeRepo.AssertNotCalled(suite.T(), "UpdateLike", mock.AnythingOfType("domain.Like"))
}

func (suite *LikeUseCaseSuite) TestLikeBlog_ExistingLike_UpdateToOppositeLike_Positive() {
	// Arrange
	islike := true
	existingLike := &domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &islike,
	}
	updatedIsLike := false
	like := domain.Like{
		BlogID:    existingLike.BlogID,
		ReacterID: existingLike.ReacterID,
		IsLike:    &updatedIsLike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(existingLike, nil)
	suite.likeRepo.On("UpdateLike", like).Return(nil)
	suite.cacheRepo.On("Get", fmt.Sprintf("LikeCount:%s", like.BlogID)).Return("5", nil)
	suite.cacheRepo.On("Get", fmt.Sprintf("DisLikeCount:%s", like.BlogID)).Return("2", nil)
	suite.cacheRepo.On("Set", fmt.Sprintf("LikeCount:%s", like.BlogID), "4", time.Duration(0)).Return(nil)
	suite.cacheRepo.On("Set", fmt.Sprintf("DisLikeCount:%s", like.BlogID), "3", time.Duration(0)).Return(nil)

	// Act
	err := suite.likeUseCase.LikeBlog(like)
	suite.likeRepo.AssertCalled(suite.T(), "UpdateLike", like)
    suite.cacheRepo.AssertCalled(suite.T(), "Set", fmt.Sprintf("LikeCount:%s", like.BlogID), "4", time.Duration(0))
    suite.cacheRepo.AssertCalled(suite.T(), "Set", fmt.Sprintf("DisLikeCount:%s", like.BlogID), "3", time.Duration(0))

	// Assert
	suite.Nil(err)
}

func (suite *LikeUseCaseSuite) TestLikeBlog_AddLike_Error_Negative() {
	// Arrange
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    func(b bool) *bool { return &b }(true),
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(nil, domain.ErrLikeNotFound)
	suite.likeRepo.On("AddLike", mock.AnythingOfType("domain.Like")).Return(domain.ErrLikeCreationFailed)

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Equal(err, domain.ErrLikeCreationFailed)
	suite.likeRepo.AssertCalled(suite.T(), "AddLike", mock.AnythingOfType("domain.Like"))
}

func (suite *LikeUseCaseSuite) TestLikeBlog_UpdateLike_Error_Negative() {
	// Arrange
	islike := true
	existingLike := &domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &islike,
	}
	updatedIsLike := false
	like := domain.Like{
		BlogID:    existingLike.BlogID,
		ReacterID: existingLike.ReacterID,
		IsLike:    &updatedIsLike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(existingLike, nil)
	suite.likeRepo.On("UpdateLike", like).Return(domain.ErrLikeUpdateFailed)

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Equal(err, domain.ErrLikeUpdateFailed)
	suite.likeRepo.AssertCalled(suite.T(), "UpdateLike", like)
}

func (suite *LikeUseCaseSuite) TestLikeBlog_CacheErrorIgnored_Positive() {
	// Arrange
	isLike := true
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &isLike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(nil, domain.ErrLikeNotFound)
	suite.likeRepo.On("AddLike", mock.AnythingOfType("domain.Like")).Return(nil)
	suite.cacheRepo.On("Get", fmt.Sprintf("LikeCount:%s", like.BlogID)).Return("", domain.ErrCacheNotFound)
	

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Nil(err)
	suite.cacheRepo.AssertNotCalled(suite.T(), "Set", fmt.Sprintf("LikeCount:%s", like.BlogID), "1", mock.AnythingOfType("time.Duration"))
	suite.likeRepo.AssertCalled(suite.T(), "AddLike", mock.AnythingOfType("domain.Like"))
}

func (suite *LikeUseCaseSuite) TestLikeBlog_CacheSuccess_Positive() {
	// Arrange
	isLike := false
	like := domain.Like{
		BlogID:    uuid.New(),
		ReacterID: uuid.New(),
		IsLike:    &isLike,
	}
	suite.likeRepo.On("GetLike", like.BlogID, like.ReacterID).Return(nil, domain.ErrLikeNotFound)
	suite.likeRepo.On("AddLike", mock.AnythingOfType("domain.Like")).Return(nil)
	suite.cacheRepo.On("Get", fmt.Sprintf("DisLikeCount:%s", like.BlogID)).Return("0", nil)
	suite.cacheRepo.On("Set", fmt.Sprintf("DisLikeCount:%s", like.BlogID), "1", mock.AnythingOfType("time.Duration")).Return(nil)
	

	// Act
	err := suite.likeUseCase.LikeBlog(like)

	// Assert
	suite.Nil(err)
	suite.cacheRepo.AssertCalled(suite.T(), "Set", fmt.Sprintf("DisLikeCount:%s", like.BlogID), "1", mock.AnythingOfType("time.Duration"))
	suite.likeRepo.AssertCalled(suite.T(), "AddLike", mock.AnythingOfType("domain.Like"))
}

func (suite *LikeUseCaseSuite) Test_Delete_Like_Positive() {
	blogID := uuid.New()
	reacterID := uuid.New()
	like := dto.UnlikeDto{
		BlogID:    blogID,
		ReacterID: reacterID,
	}
	suite.likeRepo.On("DeleteLike", like).Return(nil)
	err := suite.likeUseCase.DeleteLike(like)
	suite.Nil(err)

}
func (suite *LikeUseCaseSuite) Test_Delete_Like_Negative() {
	blogID := uuid.New()
	reacterID := uuid.New()
	like := dto.UnlikeDto{
		BlogID:    blogID,
		ReacterID: reacterID,
	}
	suite.likeRepo.On("DeleteLike", like).Return(domain.ErrLikeDeletionFailed)
	err := suite.likeUseCase.DeleteLike(like)
	suite.Equal(err, domain.ErrLikeDeletionFailed)
}
func TestLikeUseCaseSuite(t *testing.T) {
	suite.Run(t, new(LikeUseCaseSuite))
}

package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type LikeUsecase struct {
	LikeRepo interfaces.LikeRepositoryInterface
}

type LikeUsecaseInterface interface {
	GetLike(blogID uuid.UUID, reacterID uuid.UUID) (domain.Like, *domain.CustomError)
	LikeBlog(like domain.Like) *domain.CustomError
	DeleteLike(like dto.UnlikeDto) *domain.CustomError
	BlogLikeCount(blogID uuid.UUID) (int, *domain.CustomError)
}

func NewLikeUseCase(likeRepo interfaces.LikeRepositoryInterface) *LikeUsecase {
	return &LikeUsecase{
		LikeRepo: likeRepo,
	}
}
func (l *LikeUsecase) GetLike(blogID uuid.UUID, reacterID uuid.UUID) (domain.Like, *domain.CustomError) {
	return l.LikeRepo.GetLike(blogID, reacterID)
}

// LikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) LikeBlog(like domain.Like) *domain.CustomError {
	like.ID = uuid.New()
	return l.LikeRepo.LikeBlog(like)
}

// DisLikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) DeleteLike(like dto.UnlikeDto) *domain.CustomError {
	return l.LikeRepo.DeleteLike(like)
}

func (l *LikeUsecase) BlogLikeCount(blogID uuid.UUID) (int, *domain.CustomError) {
	return l.LikeRepo.BlogLikeCount(blogID)
}

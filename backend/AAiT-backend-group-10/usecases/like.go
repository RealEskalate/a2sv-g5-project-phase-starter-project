package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type LikeUsecase struct {
	LikeRepo interfaces.LikeRepositoryInterface
}

type LikeUsecaseInterface interface {
	LikeBlog(like domain.Like) error
	DeleteLike(like domain.Like) error
	BlogLikeCount(blogID uuid.UUID) (int, error)
}

func NewLikeUseCase(likeRepo interfaces.LikeRepositoryInterface) LikeUsecaseInterface {
	return &LikeUsecase{
		LikeRepo: likeRepo,
	}
}

// LikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) LikeBlog(like domain.Like) error {
	return l.LikeRepo.LikeBlog(like)
}

// DisLikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) DeleteLike(like domain.Like) error {
	return l.LikeRepo.DeleteLike(like)
}

func (l *LikeUsecase) BlogLikeCount(blogID uuid.UUID) (int, error) {
	return l.LikeRepo.BlogLikeCount(blogID)
}

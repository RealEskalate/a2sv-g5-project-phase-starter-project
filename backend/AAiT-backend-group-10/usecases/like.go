package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type LikeUsecase struct {
	LikeRepo interfaces.LikeRepositoryInterface
}

type LikeUsecaseInterface interface {
	GetLike(likeID uuid.UUID) (domain.Like, error)
	LikeBlog(like domain.Like) error
	DeleteLike(like domain.Like) error
	BlogLikeCount(blogID uuid.UUID) (int, error)
}

func NewLikeUseCase(likeRepo interfaces.LikeRepositoryInterface) LikeUsecaseInterface {
	return &LikeUsecase{
		LikeRepo: likeRepo,
	}
}
func (l *LikeUsecase) GetLike(likeID uuid.UUID) (domain.Like, error) {
	return l.LikeRepo.GetLike(likeID)
}

// LikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) LikeBlog(like domain.Like) error {
	return l.LikeRepo.LikeBlog(like)
}

// DisLikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) DeleteLike(like domain.Like) error {
	originalLike, err := l.LikeRepo.GetLike(like.ID)
	if err != nil {
		return err
	}
	if originalLike.UserID != like.UserID {
		return errors.New("You are not authorized to delete this like")
	}
	return l.LikeRepo.DeleteLike(like)
}

func (l *LikeUsecase) BlogLikeCount(blogID uuid.UUID) (int, error) {
	return l.LikeRepo.BlogLikeCount(blogID)
}

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
	GetLike(likeID uuid.UUID) (domain.Like, *domain.CustomError)
	LikeBlog(like domain.Like) *domain.CustomError
	DeleteLike(like domain.Like) *domain.CustomError
	BlogLikeCount(blogID uuid.UUID) (int, *domain.CustomError)
}

func NewLikeUseCase(likeRepo interfaces.LikeRepositoryInterface) LikeUsecaseInterface {
	return &LikeUsecase{
		LikeRepo: likeRepo,
	}
}
func (l *LikeUsecase) GetLike(likeID uuid.UUID) (domain.Like, *domain.CustomError) {
	return l.LikeRepo.GetLike(likeID)
}

// LikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) LikeBlog(like domain.Like) *domain.CustomError {
	return l.LikeRepo.LikeBlog(like)
}

// DisLikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) DeleteLike(like domain.Like) *domain.CustomError {
	originalLike, err := l.LikeRepo.GetLike(like.ID)
	if err != nil {
		return err
	}
	if originalLike.UserID != like.UserID {
		return domain.ErrUnAuthorized
	}
	return l.LikeRepo.DeleteLike(like)
}

func (l *LikeUsecase) BlogLikeCount(blogID uuid.UUID) (int, *domain.CustomError) {
	return l.LikeRepo.BlogLikeCount(blogID)
}

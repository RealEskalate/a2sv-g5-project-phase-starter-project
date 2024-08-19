package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type LikeUsecase struct {
	LikeRepo       domain.LikeRepository
	contextTimeout time.Duration
}

func NewLikeUsecase(Likerepo domain.LikeRepository, timeout time.Duration) domain.LikeUsecase {
	return &LikeUsecase{
		LikeRepo:       Likerepo,
		contextTimeout: timeout,
	}

}

func (luse *LikeUsecase) GetLikes(c context.Context, post_id string) ([]domain.Like, error) {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.LikeRepo.GetLikes(post_id)
}

func (luse *LikeUsecase) CreateLike(c context.Context, user_id string, post_id string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.LikeRepo.CreateLike(user_id, post_id)
}

func (luse *LikeUsecase) DeleteLike(c context.Context, like_id string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.LikeRepo.DeleteLike(like_id)
}

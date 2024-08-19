package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type DisLikeUsecase struct {
	DisLikeRepo    domain.DisLikeRepository
	contextTimeout time.Duration
}

func NewDislikeUsecase(DisLikerepo domain.DisLikeRepository, timeout time.Duration) domain.DisLikeUsecase {
	return &DisLikeUsecase{
		DisLikeRepo:    DisLikerepo,
		contextTimeout: timeout,
	}

}

func (luse *DisLikeUsecase) GetDisLikes(c context.Context, post_id string) ([]domain.DisLike, error) {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.DisLikeRepo.GetDisLikes(post_id)
}

func (luse *DisLikeUsecase) CreateDisLike(c context.Context, user_id string, post_id string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.DisLikeRepo.CreateDisLike(user_id, post_id)
}

func (luse *DisLikeUsecase) DeleteDisLike(c context.Context, Dislike_id string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.DisLikeRepo.DeleteDisLike(Dislike_id)
}

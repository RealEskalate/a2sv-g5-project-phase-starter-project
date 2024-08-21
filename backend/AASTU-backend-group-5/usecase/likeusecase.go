package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/repository"
)

type LikeUseCase struct {
	LikeRepo    domain.Like_Repositroy_interface
	DislikeRepo *repository.DislikeRepository
}

func NewLikeUseCase(likeRepo domain.Like_Repositroy_interface, dislikeRepo *repository.DislikeRepository) *LikeUseCase {
	return &LikeUseCase{
		LikeRepo:    likeRepo,
		DislikeRepo: dislikeRepo,
	}
}

func (BC *LikeUseCase) GetLikes(post_id string) ([]domain.Like, error) {
	return BC.LikeRepo.GetLikes(post_id)
}

func (BC *LikeUseCase) CreateLike(user_id string, post_id string) error {
	return BC.LikeRepo.CreateLike(user_id, post_id)
}

func (BC *LikeUseCase) ToggleLike(user_id string, post_id string) error {
	likeRepo := BC.LikeRepo.(*repository.LikeRepository)
	return likeRepo.ToggleLike(user_id, post_id)
}

func (BC *LikeUseCase) RemoveLike(user_id string, post_id string) error {
	likeRepo := BC.LikeRepo.(*repository.LikeRepository)
	return likeRepo.RemoveLike(user_id, post_id)
}

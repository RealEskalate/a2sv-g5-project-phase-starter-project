package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/repository"
)

type DislikeUseCase struct {
	DislikeRepo domain.DisLike_Repository_interface
	LikeRepo    *repository.LikeRepository
}

func NewDislikeUseCase(dislikeRepo domain.DisLike_Repository_interface, likeRepo *repository.LikeRepository) *DislikeUseCase {
	return &DislikeUseCase{
		DislikeRepo: dislikeRepo,
		LikeRepo:    likeRepo,
	}
}

func (BC *DislikeUseCase) GetDislikes(post_id string) ([]domain.DisLike, error) {
	return BC.DislikeRepo.GetDisLikes(post_id)
}

func (BC *DislikeUseCase) CreateDislike(user_id string, post_id string) error {
	return BC.DislikeRepo.CreateDisLike(user_id, post_id)
}

func (BC *DislikeUseCase) ToggleDislike(user_id string, post_id string) error {
	dislikeRepo := BC.DislikeRepo.(*repository.DislikeRepository)
	return dislikeRepo.ToggleDislike(user_id, post_id)
}

func (BC *DislikeUseCase) RemoveDislike(user_id string, post_id string) error {
	dislikeRepo := BC.DislikeRepo.(*repository.DislikeRepository)
	return dislikeRepo.RemoveDislike(user_id, post_id)
}

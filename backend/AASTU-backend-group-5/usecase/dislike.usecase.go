package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type DislikeUseCase struct {
	DislikeRepo domain.DisLike_Repository_interface
}

func (BC *DislikeUseCase) GetDisLikes(post_id string) ([]domain.DisLike, error) {
	return BC.DislikeRepo.GetDisLikes(post_id)
}
func (BC *DislikeUseCase) CreateDisLike(user_id string, post_id string) error {
	return BC.DislikeRepo.CreateDisLike(user_id, post_id)
}
func (BC *DislikeUseCase) DeleteDisLike(like_id string) error {
	return BC.DislikeRepo.DeleteDisLike(like_id)
}

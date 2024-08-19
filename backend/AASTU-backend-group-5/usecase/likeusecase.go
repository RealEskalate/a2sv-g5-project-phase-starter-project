package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type LikeUseCase struct {
	LikeRepo domain.Like_Repositroy_interface
}

func (BC *LikeUseCase) GetLikes(post_id string) ([]domain.Like, error) {
	return BC.LikeRepo.GetLikes(post_id)
}

func (BC *LikeUseCase) CreateLike(user_id string, post_id string) error {
	return BC.LikeRepo.CreateLike(user_id, post_id)
}
func (BC *LikeUseCase) DeleteLike(like_id string) error {
	return BC.LikeRepo.DeleteLike(like_id)
}

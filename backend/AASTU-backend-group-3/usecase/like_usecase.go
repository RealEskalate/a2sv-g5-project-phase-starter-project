package usecase

import "group3-blogApi/domain"

type LikeUsecaseImpl struct {
	likeRepo domain.LikeRepository
}

func NewLikeUsecase(likeRepo domain.LikeRepository) domain.LikeUsecase {
	return &LikeUsecaseImpl{
		likeRepo: likeRepo,
	}
}

func (lu *LikeUsecaseImpl) LikeBlog(userID, blogID, Type string) error {
	return lu.likeRepo.LikeBlog(userID, blogID, Type)
}

func (lu *LikeUsecaseImpl) DisLikeBlog(userID, blogID, Type string) error {
	return lu.likeRepo.DisLikeBlog(userID, blogID, Type)
}
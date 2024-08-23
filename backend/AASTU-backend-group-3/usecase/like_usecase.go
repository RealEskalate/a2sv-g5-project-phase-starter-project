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

func (lu *LikeUsecaseImpl) LikeBlog(userID, blogID, Type string) *domain.CustomError {
	if lu.likeRepo.LikeBlog(userID, blogID, Type) != nil {
		return domain.ErrFailedToLikePost
	}
	return nil
}

func (lu *LikeUsecaseImpl) DisLikeBlog(userID, blogID, Type string) *domain.CustomError {
	if lu.likeRepo.DisLikeBlog(userID, blogID, Type) != nil {
		return domain.ErrFailedToUnlikePost
	}
	return nil
}
package usecases

import (
	"meleket/domain"
	"meleket/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeUsecase struct {
	likeRepo repository.LikeRepositoryInterface
}

func NewLikeUsecase(likeRepo repository.LikeRepositoryInterface) *LikeUsecase {
	return &LikeUsecase{likeRepo: likeRepo}
}

func (u *LikeUsecase) AddLike(blogID, userID primitive.ObjectID) error {
	like := &domain.Like{
		ID:        primitive.NewObjectID(),
		BlogID:    blogID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return u.likeRepo.AddLike(like)
}

func (u *LikeUsecase) RemoveLike(likeID primitive.ObjectID) error {
	return u.likeRepo.RemoveLike(likeID)
}

func (u *LikeUsecase) GetLikesByBlogID(blogID primitive.ObjectID) ([]domain.Like, error) {
	return u.likeRepo.GetLikesByBlogID(blogID)
}

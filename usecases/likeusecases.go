package usecases

import (
	"meleket/domain"
	"meleket/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeDislikeUsecase struct {
	repo repository.LikeDislikeRepositoryInterface
}

func NewLikeDislikeUsecase(repo repository.LikeDislikeRepositoryInterface) *LikeDislikeUsecase {
	return &LikeDislikeUsecase{repo: repo}
}

func (u *LikeDislikeUsecase) ToggleLikeDislike(blogID, userID primitive.ObjectID, likeType string) error {
	// Check if the user has already liked or disliked the blog
	existing, err := u.repo.GetLikeDislikeByBlogAndUser(blogID, userID)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if existing != nil {
		if existing.Type == likeType {
			// If the user is trying to do the same action, remove it
			return u.repo.RemoveLikeDislike(existing.ID)
		} else {
			// If the user is toggling the action, update it
			return u.repo.RemoveLikeDislike(existing.ID)
		}
	}

	// Otherwise, add the new like or dislike
	newLikeDislike := &domain.LikeDislike{
		ID:        primitive.NewObjectID(),
		BlogID:    blogID,
		UserID:    userID,
		Type:      likeType,
		CreatedAt: time.Now(),
	}

	return u.repo.AddLikeDislike(newLikeDislike)
}

// func (u *LikeDislikeUsecase) GetLikesDislikesByBlog(blogID primitive.ObjectID) ([]domain.LikeDislike, error) {
// 	// This function will return all likes and dislikes for a specific blog post
// 	// Implement as needed
// }

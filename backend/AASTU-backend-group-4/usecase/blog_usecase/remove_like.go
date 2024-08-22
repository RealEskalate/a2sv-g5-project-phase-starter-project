package blog_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) RemoveLike(ctx context.Context, userID primitive.ObjectID, likeID primitive.ObjectID, isAdmin bool) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	existingLike, err := bu.likeRepo.GetLikeByID(ctx, likeID)
	if err != nil {
		return err
	}

	if !isAdmin && existingLike.UserID != userID {
		return errors.New("you do not have permission to delete this blog post")
	}

	return bu.blogRepo.DeleteBlog(ctx, likeID)
}

package blog_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) RemoveLike(ctx context.Context, userID primitive.ObjectID, likeID primitive.ObjectID, isAdmin bool) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.likeRepo.RemoveLike(ctx, userID, likeID, isAdmin)
}

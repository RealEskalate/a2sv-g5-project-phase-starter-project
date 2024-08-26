package blog_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) RemoveComment(ctx context.Context, userID primitive.ObjectID, commentID primitive.ObjectID, isAdmin bool) error {
	// ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	// defer cancel()

	return bu.commentRepo.RemoveComment(ctx, userID, commentID, isAdmin)
}

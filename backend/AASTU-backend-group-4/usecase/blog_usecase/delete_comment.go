package blog_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) DeleteComment(ctx context.Context, userID primitive.ObjectID, commentID primitive.ObjectID, isAdmin bool) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	existingComment, err := bu.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if !isAdmin && existingComment.UserID != userID {
		return errors.New("you do not have permission to delete this blog post")
	}

	return bu.blogRepo.DeleteBlog(ctx, commentID)
}

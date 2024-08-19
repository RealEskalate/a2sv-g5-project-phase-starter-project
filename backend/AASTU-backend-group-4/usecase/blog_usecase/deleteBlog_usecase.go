package blog_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) DeleteBlog(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID, isAdmin bool) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	existingBlog, err := bu.blogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}
	if !isAdmin && existingBlog.AuthorID != userID {
		return errors.New("you do not have permission to delete this blog post")
	}

	err = bu.blogRepo.DeleteBlog(ctx, blogID)
	if err != nil {
		return err
	}

	return nil
}

package blog_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) DeleteBlog(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID, isAdmin bool) error {
	// ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	// defer cancel()

	err := bu.blogRepo.DeleteBlog(ctx, userID, blogID, isAdmin)
	if err != nil {
		return errors.New("couldn't delete blog, process is canceled")
	}

	err = bu.likeRepo.RemoveBlogLikes(ctx, blogID)
	if err != nil {
		return errors.New("blog is deleted successfully, but couldn't delete likes")
	}
	err = bu.commentRepo.RemoveBlogComments(ctx, blogID)
	if err != nil {
		return errors.New("blog is deleted successfully, but couldn't delete comments")
	}
	return nil
}
